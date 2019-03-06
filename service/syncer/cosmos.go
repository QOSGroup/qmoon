// Copyright 2018 The COSMOS Authors

package syncer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/metric"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	tmtypes "github.com/tendermint/tendermint/types"
)

type COSMOS struct {
	node  *service.Node
	tmcli *lib.TmClient
}

func (s COSMOS) Block(b types.Block) error {
	return s.block(&b)
}
func (s COSMOS) Validator(val types.Validators) error {
	return nil
}
func (s COSMOS) ConsensusState(cs types.ResultConsensusState) error {
	return nil
}

// BlockLoop 同步块
func (s COSMOS) BlockLoop(ctx context.Context) error {
	if !s.Lock(LockTypeBlock) {
		log.Printf("[Sync] BlockLoop %v err, has been locked.", s.node.ChanID)
		return nil
	}
	defer s.Unlock(LockTypeBlock)

	var height int64 = 1
	latest, err := s.node.LatestBlock()
	if err == nil && latest != nil {
		height = latest.Height + 1
	}
	log.Printf("[Sync] block start:%d", height)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			b, err := s.tmcli.RetrieveBlock(&height)
			if err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}

			if err := s.block(b); err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}
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

	s.tx(b)
	s.node.SaveBlockValidator(b.Precommits)

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

func (s COSMOS) ValidatorLoop(ctx context.Context) error {
	if !s.Lock(LockTypeValidator) {
		log.Printf("[Sync] ValidatorLoop %v err, has been locked.", s.node.ChanID)

		return nil
	}
	defer s.Unlock(LockTypeValidator)
	var height int64 = 1

	for {
		time.Sleep(SyncValidator)
		latest, err := s.node.LatestBlock()
		if err == nil && latest != nil {
			height = latest.Height + 1
		}
		select {
		case <-ctx.Done():
			return nil
		default:
			vals, err := s.tmcli.Validator(height)
			if err != nil {
				time.Sleep(time.Millisecond * 100)
				continue
			}

			for _, val := range vals {
				s.node.CreateValidator(val)
			}
			metric.ValidatorVotingPower(vals)

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
