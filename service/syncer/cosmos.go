// Copyright 2018 The COSMOS Authors

package syncer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/metric"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/sirupsen/logrus"
	tmtypes "github.com/tendermint/tendermint/types"
)

type COSMOS struct {
	node  *service.Node
	tmcli *lib.TmClient
}

func (s COSMOS) RpcPeers(ctx context.Context) error {
	return nil
}

// BlockLoop 同步块
func (s COSMOS) BlockLoop(ctx context.Context) error {
	if !s.Lock(LockTypeBlock) {
		log.Printf("COSMOS [Sync] BlockLoop %v err, has been locked.", s.node.ChanID)
		return nil
	}
	defer s.Unlock(LockTypeBlock)

	var height int64 = 1
	latest, err := s.node.LatestBlock()
	if err == nil && latest != nil {
		height = latest.Height + 1
	}
	logrus.Infof("COSMOS [Sync] block start:%d", height)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("COSMOS [Sync] block height:%d", height)
			b, err := s.tmcli.RetrieveBlock(&height)
			if err != nil {
				log.Printf("COSMOS [Sync] BlockLoop  RetrieveBlock height:%d, err:%v", height, err)
				time.Sleep(time.Millisecond * 1000)
				continue
			}

			if err := s.block(b); err != nil {
				log.Printf("[Sync] BlockLoop block height:%d, err:%v", height, err)
				time.Sleep(time.Millisecond * 100)
				continue
			}
			s.Validator(height, b.Header.Time)
			height += 1
		}
	}

	return nil
}

// block
func (s COSMOS) block(b *types.Block) error {
	err := s.node.CreateBlock(b)
	if err != nil {
		return err
	}

	if err := s.node.CreateEvidence(b); err != nil {
		log.Printf("COSMOS [Sync] CreateEvidence  err:%v", err)
	}

	if err := s.tx(b); err != nil {
		log.Printf("COSMOS [Sync] block  tx err:%v", err)
	}
	if err := s.node.SaveBlockValidator(b.Precommits); err != nil {
		log.Printf("COSMOS [Sync] block  tx err:%v", err)
	}

	return nil
}

func (s COSMOS) tx(b *types.Block) error {
	if len(b.Txs) == 0 {
		return nil
	}

	var txs []string
	for _, v := range b.Txs {
		txs = append(txs, utils.Base64En(v))
	}
	txres, err := lib.NewCosmosCli("").Txs(txs)
	if err != nil {
		return err
	}

	for k, v := range txres {
		var txTypes []string
		for _, tt := range v.Txs {
			txTypes = append(txTypes, tt.Type)
		}
		hash := tmtypes.Tx(b.Txs[k]).Hash()
		mt := &models.Tx{}
		mt.Height = b.Header.Height
		mt.Index = int64(k)
		mt.Hash = strings.ToUpper(hex.EncodeToString(hash))
		mt.TxType = strings.Join(txTypes, ";")
		mt.OriginTx = txs[k]
		if d, err := json.Marshal(v.Txs); err == nil {
			mt.JsonTx = string(d)
		}
		mt.Time = b.Header.Time

		txResult, err := s.tmcli.RetrieveTx(hash)
		if err == nil {
			mt.TxStatus = int(txResult.TxStatus)
			mt.GasWanted = txResult.GasWanted
			mt.GasUsed = txResult.GasUsed
			mt.Log = txResult.Log
		}
		mt.Fee = v.Fee
		if err := mt.Insert(s.node.ChanID); err != nil {
			log.Printf("tx insert data:%+v, err:%v", mt, err.Error())
			return err
		}

		if v.Fee != "" {
			if err := models.UpdateFee(b.Header.ChainID, txTypes[0], v.Fee, mt.GasWanted, mt.GasUsed); err != nil {
				log.Printf("UpdateFee err:%s", err.Error())
			}
		}
	}

	return nil
}

type StakingValidator struct {
	Commission struct {
		MaxChangeRate string `json:"max_change_rate"`
		MaxRate       string `json:"max_rate"`
		Rate          string `json:"rate"`
		UpdateTime    string `json:"update_time"`
	} `json:"commission"`
	ConsensusPubkey string `json:"consensus_pubkey"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Details  string `json:"details"`
		Identity string `json:"identity"`
		Moniker  string `json:"moniker"`
		Website  string `json:"website"`
	} `json:"description"`
	Jailed            bool   `json:"jailed"`
	MinSelfDelegation string `json:"min_self_delegation"`
	OperatorAddress   string `json:"operator_address"`
	Status            int    `json:"status"`
	Tokens            string `json:"tokens"`
	UnbondingHeight   string `json:"unbonding_height"`
	UnbondingTime     string `json:"unbonding_time"`
}

func (s COSMOS) stakingValidators() map[string]StakingValidator {
	k := "cosmosStakingValidators"
	if v, ok := cache.Get(k); ok {
		return v.(map[string]StakingValidator)
	}

	res := make(map[string]StakingValidator)

	response, err := s.tmcli.ABCIQuery("custom/staking/validators", nil)
	if err != nil {
		return res
	}

	var sv []StakingValidator
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

func (s COSMOS) Validator(height int64, t time.Time) error {
	vals, err := s.tmcli.Validator(height)
	if err != nil {
		log.Printf("COSMOS [Sync] ValidatorLoop  Validator err:%v", err)
		return err
	}

	svs := s.stakingValidators()
	for _, val := range vals {
		if sv, ok := svs[lib.PubkeyToBech32Address(s.node.Bech32PrefixConsPub(), val.PubKeyType, val.PubKeyValue)]; ok {
			val.Name = sv.Description.Moniker
			val.Website = sv.Description.Website
			val.Identity = sv.Description.Identity
			val.Details = sv.Description.Details
		}

		s.node.CreateValidator(val)
	}

	valMap := make(map[string]types.Validator)
	for _, v := range vals {
		valMap[v.Address] = v
	}

	oldVals, err := s.node.Validators()
	if err == nil {
		for _, v := range oldVals {
			if v.Status == types.Active {
				if _, ok := valMap[v.Address]; !ok {
					_ = s.node.InactiveValidator(v.Address, int(types.Inactive), height, time.Time{})
				}
			} else {
				if _, ok := valMap[v.Address]; ok {
					_ = s.node.InactiveValidator(v.Address, int(types.Active), height, time.Time{})
				}
			}
		}
	}
	metric.ValidatorVotingPower(s.node.ChanID, t, oldVals)

	return nil
}

func (s COSMOS) ConsensusStateLoop(ctx context.Context) error {
	if !s.Lock(LockTypeConsensusState) {
		log.Printf("[Sync] ConsensusStateLoop %v err, has been locked.", s.node.ChanID)

		return nil
	}
	defer s.Unlock(LockTypeConsensusState)

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

func (s COSMOS) PeerLoop(ctx context.Context) error {
	if !s.Lock(LockTypePeer) {
		log.Printf("[Sync] PeerLoop %v err, has been locked.", s.node.ChanID)
		return nil
	}
	defer s.Unlock(LockTypePeer)

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

// SyncLock 同步时锁定，同一个时间只会有一个同步协程
func (s COSMOS) Lock(key string) bool {
	key = "lock_" + s.node.ChanID + "-" + key

	qs, err := models.RetrieveQmoonStatusByKey(key)
	if err != nil {
		if errors.IsNotExist(err) {
			qs = &models.QmoonStatus{
				Key:   key,
				Value: SyncLocked,
			}
			err := qs.Insert()
			return err == nil
		} else {
			return false
		}
	}

	// 被锁住未超过最大值
	if qs.Value == SyncLocked && qs.UpdatedAt.Add(maxLockDuration).After(time.Now()) {
		return false
	}

	qs.Value = SyncLocked
	if err := qs.Update(); err != nil {
		log.Printf("Sync Lock %s err:%v", key, err.Error())
		return false
	}

	return true
}

func (s COSMOS) Unlock(key string) bool {
	key = "lock_" + s.node.ChanID + "-" + key

	qs, err := models.RetrieveQmoonStatusByKey(key)
	if err != nil {
		return true
	}

	qs.Value = SyncUnlocked

	if err := qs.Update(); err != nil {
		log.Printf("Sync Unlock %s err:%v", key, err.Error())

		return false
	}

	return true
}
