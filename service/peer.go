// Copyright 2018 The QOS Authors

package service

import (
	"database/sql"
	"log"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

func convertToPeer(mp *model.Peer) *types.ResultPeer {
	return &types.ResultPeer{
		Moniker:    mp.Moniker.String,
		ID:         mp.ID,
		PeerID:     mp.PeerID.String,
		ListenAddr: mp.ListenAddr.String,
		Network:    mp.Network.String,
		Version:    mp.Version.String,
		Channels:   mp.Channels.String,
		SendStart:  mp.SendStart.Time,
		RecvStart:  mp.RecvStart.Time,
		CreateAt:   mp.CreatedAt.Time,
	}
}

// ListPeers 查询所有peer
func ListPeers(chainID string) (*types.ResultPeers, error) {
	var result types.ResultPeers

	mps, err := model.PeersByChainID(db.Db, utils.NullString(chainID))
	if err == nil {
		for _, v := range mps {
			result.NPeers++
			result.Peers = append(result.Peers, convertToPeer(v))
		}
	}

	return &result, nil
}

func RetrieveOrInsert(chainID string, tmp tmtypes.Peer) (*model.Peer, error) {
	p, err := model.PeerByPeerID(db.Db, utils.NullString(string(tmp.ID)))
	if err != nil {
		if err == sql.ErrNoRows {
			p = &model.Peer{
				ChainID:    utils.NullString(chainID),
				Moniker:    utils.NullString(tmp.Moniker),
				PeerID:     utils.NullString(string(tmp.ID)),
				ListenAddr: utils.NullString(tmp.ListenAddr),
				Network:    utils.NullString(tmp.Network),
				Version:    utils.NullString(tmp.Version),
				Channels:   utils.NullString(tmp.Channels.String()),
				SendStart:  utils.NullTime(tmp.ConnectionStatus.SendMonitor.Start),
				RecvStart:  utils.NullTime(tmp.ConnectionStatus.RecvMonitor.Start),
				CreatedAt:  utils.NullTime(time.Now()),
			}
			err = p.Insert(db.Db)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return p, nil
}

func SavePeers(chainID string, peers []tmtypes.Peer) error {
	for _, v := range peers {
		_, err := RetrieveOrInsert(chainID, v)
		if err != nil {
			log.Printf("RetrieveOrInsert p:%+v, err:%s", v, err.Error())
		}
	}

	return nil
}
