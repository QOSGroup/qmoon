// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/QOSGroup/qmoon/plugins"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/hashicorp/go-version"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tidwall/gjson"
)

type QOS struct {
	node  *service.Node
	tmcli *lib.TmClient
}

var (
	qos0_0_3 *version.Version
	qos0_0_4 *version.Version
	qos0_0_8 *version.Version
)

func init() {
	qos0_0_3, _ = version.NewVersion("0.0.3")
	qos0_0_4, _ = version.NewVersion("0.0.4")
	qos0_0_8, _ = version.NewVersion("0.0.8")
}

func (s QOS) RpcPeers(ctx context.Context) error {
	netinfo, err := s.tmcli.NetInfo()
	if err != nil {
		return err
	}

	//var peers []string
	for _, peer := range netinfo.Peers {
		remoteIp := peer.RemoteIP
		u, err := url.Parse(peer.NodeInfo.Other.RPCAddress)
		if err != nil {
			log.Printf("[QOS] Parse Peer RPCAddress err:%s", err.Error())
			continue
		}

		u.Host = fmt.Sprintf("%s:%s", remoteIp, u.Port())
		//peers = append(peers, u.String())

		_ = models.CreateNetworkOrUpdate(s.node.ChainID, &models.Network{Remote: u.String()})
	}

	return nil
}

// BlockLoop 同步块
func (s QOS) BlockLoop(ctx context.Context) error {
	key := "lock_" + s.node.ChainID + "-" + LockTypeBlock
	fmt.Println("[Sync] Start Syncing blocks")

	if !Lock(key) {
		log.Printf("[Sync] QOS BlockLoop %v err, has been locked.", s.node.ChainID)
		return nil
	}
	defer Unlock(key)

	var height int64 = 1
	latest, err := s.node.LatestBlock()
	if err == nil && latest != nil {
		height = latest.Height + 1
	}
	//latestheight, err := s.node.LatestBlockHeight()
	//if err == nil && latestheight != 0 {
	//	height = latestheight + 1
	//}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			b, err := s.tmcli.RetrieveBlock(&height)
			if err != nil {
				time.Sleep(time.Millisecond * 1000)
				continue
			}

			if err := s.block(b); err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}
			s.Validator(height, b.Header.Time)
			height += 1
			// 为什么要同步proposal？
			// s.Proposals()
		}
	}

	return nil
}

// block
func (s QOS) block(b *types.Block) error {
	err := s.node.CreateBlock(b)
	if err != nil {
		return err
	}

	err = s.tx(b)
	if err != nil {
		fmt.Println("parse tx err:", err)
	}

	err = s.node.SaveBlockValidator(b.Precommits)
	if err != nil {
		fmt.Println("parse BlockValidator err:", err)
	}

	return nil
}

func (s QOS) tx(b *types.Block) error {
	// if len(b.Txs) == 0 {
	// 	return nil
	// }

	// var txs []string
	// for _, v := range b.Txs {
	// 	txs = append(txs, utils.Base64En(v))
	// }

	for k, v := range b.Txs {
		qbasetx, err := getQbaseTx(v)
		if err != nil {
			return err
		}
		// var txTypes []string
		// for _, tt := range v.Txs {
		// 	txTypes = append(txTypes, tt.Type)
		// }

		hash := tmtypes.Tx(b.Txs[k]).Hash()
		mt := &models.Tx{}
		mt.Hash = strings.ToUpper(hex.EncodeToString(hash))
		mt.Height = b.Header.Height
		mt.Index = int64(k)
		mt.OriginTx = utils.Base64En(v)
		// mt.OriginTx = txs[k]
		// if d, err := json.Marshal(v.Txs); err == nil {
		// 	mt.JsonTx = string(d)
		// }
		mt.Time = b.Header.Time
		// mt.TxType = qbasetx.Type()
		// mt.TxType = strings.Join(txTypes, ";")

		txResult, errtx := s.tmcli.RetrieveTx(hash)
		if errtx == nil {
			mt.TxStatus = int(txResult.TxStatus)
			mt.GasWanted = txResult.GasWanted
			mt.GasUsed = txResult.GasUsed
			mt.Log = txResult.Log
		}

		if err := parseQosTx(b.Header, qbasetx, mt); err != nil {
			return err
		}

		if err := mt.InsertOrUpdate(s.node.ChainID); err != nil {
			return err
		}

	}
	return nil
}

func getQbaseTx(tx []byte) (*qbasetxs.TxStd, error) {
	cdc := lib.MakeCodec()
	qbasetx, err := qbasetypes.DecoderTx(cdc, tx)
	if err != nil {
		return nil, err
	}

	var std *qbasetxs.TxStd
	switch implTx := qbasetx.(type) {
	case *qbasetxs.TxStd:
		std = implTx
	case *qbasetxs.TxQcp:
		std = implTx.TxStd
	default:
		return nil, errors.New("未知交易")
	}

	return std, nil
}

func parseQosTx(blockHeader types.BlockHeader, t qbasetypes.Tx, mt *models.Tx) error {
	var std *qbasetxs.TxStd
	switch implTx := t.(type) {
	case *qbasetxs.TxStd:
		std = implTx
	case *qbasetxs.TxQcp:
		mt.QcpFrom = implTx.From
		mt.QcpTo = implTx.To
		mt.QcpSequence = implTx.Sequence
		mt.QcpTxindex = implTx.TxIndex
		mt.QcpIsresult = implTx.IsResult
		std = implTx.TxStd
	default:
		mt.TxType = t.Type()
	}

	mt.Maxgas = std.MaxGas.Int64()

	i := int64(0)
	for _, iTx := range std.ITxs {
		var itx = &models.ITx{Hash: mt.Hash, Seq: i}
		if err := parseQosITx(blockHeader, iTx, itx); err != nil {
			return err
		}
		mt.ITxs = append(mt.ITxs, itx)
		i++
	}

	return nil
}

func parseQosITx(blockHeader types.BlockHeader, iTx qbasetxs.ITx, itx *models.ITx) error {
	cdc := lib.MakeCodec()
	d, err := cdc.MarshalJSON(iTx)
	if err != nil {
		return err
	}

	valueInJson := gjson.Get(string(d), "value")
	if valueInJson.Exists() {
		itx.JsonTx = valueInJson.String()
	} else {
		itx.JsonTx = string(d)
	}

	typeInJson := gjson.Get(string(d), "type")
	if typeInJson.Exists() {
		itx.TxType = typeInJson.String()
	} else {
		itx.TxType = "Unknown"
	}

	name, err := plugins.Parse(blockHeader, iTx)
	log.Printf("plugins.Parse name:%s, err:%v", name, err)

	return nil
}

//QOSStakingValidator struct
type QOSStakingValidator struct {
	Commission struct {
		MaxChangeRate string `json:"max_change_rate"`
		MaxRate       string `json:"max_rate"`
		Rate          string `json:"rate"`
		UpdateTime    string `json:"update_time"`
	} `json:"commission"`
	ConsensusPubkey string `json:"consensus_pubkey"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Details string `json:"details"`
		Logo    string `json:"logo"`
		Moniker string `json:"moniker"`
		Website string `json:"website"`
	} `json:"description"`
	Jailed            bool   `json:"jailed"`
	MinSelfDelegation string `json:"min_self_delegation"`
	OperatorAddress   string `json:"operator_address"`
	Status            int    `json:"status"`
	Tokens            string `json:"tokens"`
	UnbondingHeight   string `json:"unbonding_height"`
	UnbondingTime     string `json:"unbonding_time"`
}

func (s QOS) stakingValidators() map[string]QOSStakingValidator {
	k := "qosStakingValidators"
	if v, ok := cache.Get(k); ok {
		return v.(map[string]QOSStakingValidator)
	}

	res := make(map[string]QOSStakingValidator)

	response, err := s.tmcli.ABCIQuery("/store/validator/subspace", nil)
	if err != nil {
		return res
	}

	var sv []QOSStakingValidator
	if response.Response.Code == 0 {
		json.Unmarshal(response.Response.Value, &sv)
	}

	for _, v := range sv {
		res[v.ConsensusPubkey] = v
	}

	if len(res) != 0 {
		cache.Set(k, res, time.Minute*5)
	}

	return res
}

func (s QOS) Validator(height int64, t time.Time) error {
	var vals []types.Validator
	valsMap := make(map[string]types.Validator)
	//var vals_display []stake_types.ValidatorDisplayInfo
	//var err error
	//if !s.node.NodeVersion.GreaterThan(qos0_0_4) {
	//	vals, err = s.tmcli.QOSValidator(height)
	//} else {
	vals_display, err := qos.NewQosCli("").QueryValidators(s.node.BaseURL)
	if err != nil {
		return err
	}
	// fmt.Println(len(vals), " validators")
	// svs := s.stakingValidators()
	for _, dist := range vals_display {
		val, err := s.node.ConvertDisplayValidators(dist)
		if err != nil {
			log.Printf("QOS [Sync] ValidatorLoop  Validator err:%v", err)
			return err
		}
		//if sv, ok := svs[lib.PubkeyToBech32Address(s.node.Bech32PrefixConsPub(), val.PubKeyType, val.PubKeyValue)]; ok {
		//	val.Name = sv.Description.Moniker
		//	val.Website = sv.Description.Website
		//	val.Logo = sv.Description.Logo
		//	val.Details = sv.Description.Details
		//	val.Commission = sv.Commission.Rate
		//}
		// fmt.Println("before Create ", val.Address, val.BondedTokens, val.SelfBond)
		s.node.CreateValidator(val)
		valsMap[val.Address] = val
		//valMap[val.Address] = val
		vals = append(vals, val)
	}

	mvs, err := models.Validators(s.node.ChainID)
	if err == nil {
		for _, v := range mvs {
			if _, ok := valsMap[v.Address]; !ok {
				v.Del(s.node.ChainID)
			}
		}
	}

	//metric.ValidatorVotingPower(s.node.ChainID, t, vals)

	return nil
}

//qos proposals
func (s QOS) Proposals() error {
	prores, err := qos.NewQosCli("").QueryProposals(s.node.BaseURL)
	if err != nil {
		return err
	}
	mt := &models.Proposal{}
	for _, pro := range prores {
		mt.Description = pro.Description
		mt.ProposalID = pro.ProposalID
		mt.Status = pro.Status
		mt.SubmitTime = pro.SubmitTime
		mt.Title = pro.Title
		totalDeposit, err := strconv.ParseInt(pro.TotalDeposit.String(), 10, 64)
		if err == nil {
			mt.TotalDeposit = totalDeposit
		}
		mt.Type = pro.Type
		mt.VotingStartTime = pro.VotingStartTime
		mt.VotingEndTime = pro.VotingEndTime
	}
	if err := mt.InsertOrUpdate(s.node.ChainID); err != nil {
		log.Printf("proposals insert data:%+v, err:%v", mt, err.Error())
		return err
	}
	return nil
}

func (s QOS) ConsensusStateLoop(ctx context.Context) error {
	key := "lock_" + s.node.ChainID + "-" + LockTypeConsensusState

	if !Lock(key) {
		log.Printf("[Sync] ConsensusStateLoop %v err, has been locked.", s.node.ChainID)
		return nil
	}
	defer Unlock(key)

	for {
		time.Sleep(SyncConsensusStateDuration)
		select {
		case <-ctx.Done():
			return nil
		default:
			cs, err := s.tmcli.ConsensusState()
			if err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}
			if err := s.node.UpdateConsensusState(cs); err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}
		}
	}

	return nil
}

func (s QOS) PeerLoop(ctx context.Context) error {
	key := "lock_" + s.node.ChainID + "-" + LockTypePeer
	if !Lock(key) {
		log.Printf("[Sync] PeerLoop %v err, has been locked.", s.node.ChainID)
		return nil
	}
	defer Unlock(key)

	for {
		time.Sleep(SyncPeerDuration)
		select {
		case <-ctx.Done():
			return nil
		default:
			b, err := s.tmcli.NetInfo()
			if err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}

			if b != nil {
				if err := s.node.CreatePeers(b.Peers); err != nil {
					time.Sleep(time.Millisecond * 100)
					continue
				}
			}
		}
	}

	return nil
}
