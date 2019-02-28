// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"encoding/hex"
	"log"
	"strings"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/lib"
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
)

func init() {
	qos0_0_3, _ = version.NewVersion("0.0.3")
	qos0_0_4, _ = version.NewVersion("0.0.4")
}

func (s QOS) Block(b types.Block) error {
	return s.block(&b)
}
func (s QOS) Validator(val types.Validators) error {
	return nil
}
func (s QOS) ConsensusState(cs types.ResultConsensusState) error {
	return nil
}

// BlockLoop 同步块
func (s QOS) BlockLoop(ctx context.Context) error {
	if !s.Lock(LockTypeBlock) {
		log.Printf("[Sync] QOS BlockLoop %v err, has been locked.", s.node.ChanID)
		return nil
	}
	defer s.Unlock(LockTypeBlock)

	var height int64 = 1
	latest, err := s.node.LatestBlock()
	if err == nil && latest != nil {
		height = latest.Height + 1
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			b, err := s.tmcli.RetrieveBlock(&height)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			if err := s.block(b); err != nil {
				time.Sleep(time.Second)
				continue
			}
			height += 1
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
	// TODO delete block

	err = s.node.SaveBlockValidator(b.Precommits)
	// TODO delete block and tx

	return nil
}

func (s QOS) tx(b *types.Block) error {
	for k, v := range b.Txs {
		qbasetx, err := getQbaseTx(v)
		if err != nil {
			return err
		}

		hash := tmtypes.Tx(b.Txs[k]).Hash()
		mt := &models.Tx{}
		mt.Hash = strings.ToUpper(hex.EncodeToString(hash))
		mt.Height = b.Header.Height
		mt.Index = int64(k)
		mt.OriginTx = utils.Base64En(v)
		mt.Time = b.Header.Time
		mt.TxType = qbasetx.Type()

		txResult, errtx := s.tmcli.RetrieveTx(hash)
		if errtx == nil {
			mt.TxStatus = int(txResult.TxStatus)
			mt.GasWanted = txResult.GasWanted
			mt.GasUsed = txResult.GasUsed
		}

		if err := parseQosTx(b.Header, qbasetx, mt); err != nil {
			return err
		}

		if err := mt.Insert(s.node.ChanID); err != nil {
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

	if err := parseQosITx(blockHeader, std.ITx, mt); err != nil {
		return err
	}

	if mt.TxType == "Unknown" {
		mt.TxType = t.Type()
	}

	return nil
}

func parseQosITx(blockHeader types.BlockHeader, t qbasetxs.ITx, mt *models.Tx) error {
	cdc := lib.MakeCodec()
	d, err := cdc.MarshalJSON(t)
	if err != nil {
		return err
	}

	valueInJson := gjson.Get(string(d), "value")
	if valueInJson.Exists() {
		mt.JsonTx = valueInJson.String()
	} else {
		mt.JsonTx = string(d)
	}

	typeInJson := gjson.Get(string(d), "type")
	if typeInJson.Exists() {
		mt.TxType = typeInJson.String()
	} else {
		mt.TxType = "Unknown"
	}

	name, err := plugins.Parse(blockHeader, t)
	log.Printf("plugins.Parse name:%s, err:%v", name, err)

	return nil
}

func (s QOS) ValidatorLoop(ctx context.Context) error {
	if !s.Lock(LockTypeValidator) {
		log.Printf("[Sync] ValidatorLoop %v err, has been locked.", s.node.ChanID)

		return nil
	}
	defer s.Unlock(LockTypeValidator)

	for {
		time.Sleep(SyncValidator)
		select {
		case <-ctx.Done():
			return nil
		default:
			var vals []types.Validator
			var err error
			if !s.node.NodeVersion.GreaterThan(qos0_0_4) {
				vals, err = s.tmcli.QOSValidator(0)
			} else {
				vals, err = s.tmcli.QOSValidatorV0_0_4(0)
			}
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			for _, val := range vals {
				s.node.CreateValidator(val)
			}

			valMap := make(map[string]types.Validator)
			for _, v := range vals {
				valMap[v.Address] = v
			}

			allVals, err := s.node.Validators()
			if err == nil {
				for _, v := range allVals {
					if v.Status == types.Active {
						if _, ok := valMap[v.Address]; !ok {
							s.node.InactiveValidator(v.Address, 0, 0, time.Time{})
						}
					}
				}
			}
		}
	}

	return nil
}

func (s QOS) ConsensusStateLoop(ctx context.Context) error {
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
				return err
			}
			if err := s.node.UpdateConsensusState(cs); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s QOS) PeerLoop(ctx context.Context) error {
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
				return err
			}

			if b != nil {
				if err := s.node.CreatePeers(b.Peers); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// SyncLock 同步时锁定，同一个时间只会有一个同步协程
func (s QOS) Lock(key string) bool {
	key = "lock_" + s.node.ChanID + "-" + key

	qs, err := models.RetrieveQmoonStatusByKey(key)
	if err != nil {
		log.Printf("Sync Lock err:%v", err.Error())
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
		return false
	}

	return true
}

func (s QOS) Unlock(key string) bool {
	key = "lock_" + s.node.ChanID + "-" + key

	qs, err := models.RetrieveQmoonStatusByKey(key)
	if err != nil {
		return true
	}

	qs.Value = SyncUnlocked

	if err := qs.Update(); err != nil {
		log.Printf("Sync Unlock err:%v", err.Error())

		return false
	}

	return true
}
