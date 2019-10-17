// Copyright 2018 The QOS Authors

package service

import (
	"log"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

func convertToPeer(mp *models.Peer) *types.ResultPeer {
	return &types.ResultPeer{
		Moniker:    mp.Moniker,
		ID:         mp.ID,
		PeerID:     mp.PeerID,
		ListenAddr: mp.ListenAddr,
		Network:    mp.Network,
		Version:    mp.Version,
		Channels:   mp.Channels,
		SendStart:  types.ResultTime(mp.SendStart),
		RecvStart:  types.ResultTime(mp.RecvStart),
	}
}

// Peers 查询所有peer
func (n Node) Peers() (*types.ResultPeers, error) {
	var result types.ResultPeers

	mps, err := models.Peers(n.ChainID, nil)
	if err == nil {
		for _, v := range mps {
			result.NPeers++
			result.Peers = append(result.Peers, convertToPeer(v))
		}
	}

	return &result, nil
}

func (n Node) CreatePeers(peers []tmtypes.Peer) error {
	for _, v := range peers {
		p := &models.Peer{
			Moniker:    v.NodeInfo.Moniker,
			PeerID:     string(v.NodeInfo.ID()),
			ListenAddr: v.NodeInfo.ListenAddr,
			Network:    v.NodeInfo.Network,
			Version:    v.NodeInfo.Version,
			Channels:   v.NodeInfo.Channels.String(),
			SendStart:  v.ConnectionStatus.SendMonitor.Start,
			RecvStart:  v.ConnectionStatus.RecvMonitor.Start,
		}

		if err := p.Insert(n.ChainID); err != nil {
			log.Printf("CreatePeers p:%+v, err:%s", v, err.Error())
		}
	}

	return nil
}
