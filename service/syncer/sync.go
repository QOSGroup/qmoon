// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"time"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
)

const (
	SyncConsensusStateDuration = time.Second * 2
	SyncPeerDuration           = time.Second * 30
	SyncValidator              = time.Second * 2
)

const maxLockDuration = time.Minute * 5

const (
	SyncLocked   = "1"
	SyncUnlocked = "0"
)

const (
	LockTypeBlock          = "block"
	LockTypeConsensusState = "consensus_state"
	LockTypePeer           = "peer"
	LockTypeValidator      = "validator"
)

// Syncer
type Syncer interface {
	Block(ctx context.Context) error
	Validator(ctx context.Context) error
	ConsensusState(ctx context.Context) error
	Peer(ctx context.Context) error

	Lock(key string) bool
	Unlock(key string) bool
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
