// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/models/errors"
	"log"
	"os"
	"time"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
)

const (
	SyncConsensusStateDuration = time.Second * 2
	SyncPeerDuration           = time.Second * 30
	SyncValidator              = time.Second * 2

	FastSyscThreshold = 100 // 本地高度和链高度相差值超过FastSyscThreshold时，启动fastsync mode
)

const maxLockDuration = time.Minute * 5

const (
	SyncLocked   = "1"
	SyncUnlocked = "0"
)

var SYSTEM_NAME string

func init() {
	SYSTEM_NAME, _ = os.Hostname()
}

const (
	LockTypeBlock          = "block"
	LockTypeConsensusState = "consensus_state"
	LockTypePeer           = "peer"
	LockTypeValidator      = "validator"
)

// Syncer
type Syncer interface {
	Validator(height int64, t time.Time) error
	BlockLoop(ctx context.Context) error
	ConsensusStateLoop(ctx context.Context) error
	RpcPeers(ctx context.Context) error

	//Lock(key string) bool
	//Unlock(key string) bool
}

func NewSyncer(node *service.Node) Syncer {
	switch types.NodeType(node.NodeType) {
	case types.NodeTypeQOS:
		return QOS{node: node, tmcli: lib.TendermintClient(node.BaseURL)}
	case types.NodeTypeQSC:
		return QSC{node: node, tmcli: lib.TendermintClient(node.BaseURL)}
	case types.NodeTypeCOSMOS:
		return COSMOS{node: node, tmcli: lib.TendermintClient(node.BaseURL)}
	default:
		return QOS{node: node, tmcli: lib.TendermintClient(node.BaseURL)}
	}
}

// SyncLock 同步时锁定，同一个时间只会有一个同步协程
func Lock(key string) bool {
	qs, err := models.RetrieveQmoonStatusByKey(key)
	if err != nil {
		log.Printf("Sync Lock err:%v", err.Error())
		if errors.IsNotExist(err) {
			qs = &models.QmoonStatus{
				Key:   key,
				Value: SYSTEM_NAME,
			}
			err := qs.Insert()
			return err == nil
		} else {
			return false
		}
	}
	return false
}

func Unlock(key string) bool {
	err := models.DeleteKey(key)
	return err == nil
}
