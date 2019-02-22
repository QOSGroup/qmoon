// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"log"
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
	qosapprove "github.com/QOSGroup/qos/txs/approve"
	"github.com/QOSGroup/qos/txs/qsc"
	"github.com/QOSGroup/qos/txs/staking"
	"github.com/QOSGroup/qos/txs/transfer"
	"github.com/QOSGroup/qstars/x/bank"
	"github.com/QOSGroup/qstars/x/kvstore"
	"github.com/tidwall/gjson"
)

type QOS struct {
	node  *service.Node
	tmcli *lib.TmClient
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
				return err
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
	cdc := lib.MakeCodec()

	for k, v := range b.Txs {
		qbasetx, err := qbasetypes.DecoderTx(cdc, v)
		if err != nil {
			return err
		}

		mt := &models.Tx{}
		mt.Height = b.Header.Height
		mt.Index = int64(k)
		mt.OriginTx = utils.Base64En(v)
		mt.Time = b.Header.Time
		mt.TxStatus = int(s.tmcli.RetrieveTxResult(v))

		if err := parseQosTx(b.Header, qbasetx, mt); err != nil {
			return err
		}

		if err := mt.Insert(s.node.ChanID); err != nil {
			return err
		}
	}
	return nil
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
	mt.JsonTx = string(d)

	switch t.(type) {
	case *qosapprove.TxCancelApprove:
		mt.TxType = "ApproveCancelTx"
	case *qosapprove.TxCreateApprove:
		mt.TxType = "ApproveCreateTx"
	case *qosapprove.TxDecreaseApprove:
		mt.TxType = "ApproveDecreaseTx"
	case *qosapprove.TxIncreaseApprove:
		mt.TxType = "ApproveIncreaseTx"
	case *qosapprove.TxUseApprove:
		mt.TxType = "ApproveUseTx"
	case *staking.TxRevokeValidator:
		mt.TxType = "TxRevokeValidator"
	case *staking.TxCreateValidator:
		mt.TxType = "TxCreateValidator"
	case *staking.TxActiveValidator:
		mt.TxType = "TxActiveValidator"
	case *kvstore.KvstoreTx:
		mt.TxType = "KvstoreTx"
	case *qbasetxs.QcpTxResult:
		mt.TxType = "QcpTxResult"
	case *transfer.TxTransfer:
		mt.TxType = "TransferTx"
	case *qsc.TxCreateQSC:
		mt.TxType = "TxCreateQSC"
	case *qsc.TxIssueQSC:
		mt.TxType = "TxIssueQsc"
	case *bank.WrapperSendTx:
		mt.TxType = "WrapperSendTx"
	default:
		typeInJson := gjson.Get(string(d), "type")
		if typeInJson.Exists() {
			mt.TxType = typeInJson.String()
		} else {
			mt.TxType = "Unknown"
		}
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
			vals, err := s.tmcli.QOSValidator(0)
			if err != nil {
				return err
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
