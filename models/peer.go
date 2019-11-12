// Copyright 2018 The QOS Authors

package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Peer struct {
	ID            int64     `json:"id"`
	ChainID       string    `json:"chain_id"`
	Moniker       string    `json:"moniker"`
	PeerID        string    `xorm:"unique" json:"peer_id"`
	ListenAddr    string    `json:"listen_addr"`
	Network       string    `json:"network"`
	Version       string    `json:"version"`
	Channels      string    `json:"channels"`
	SendStart     time.Time `xorm:"-" json:"send_start"`
	SendStartUnix int64
	RecvStart     time.Time `xorm:"-" json:"recv_start"`
	RecvStartUnix int64
}

func (p *Peer) BeforeInsert() {
	p.SendStartUnix = p.SendStart.Unix()
	p.RecvStartUnix = p.RecvStart.Unix()
}

func (p *Peer) BeforeUpdate() {
	p.SendStartUnix = p.SendStart.Unix()
	p.RecvStartUnix = p.RecvStart.Unix()
}

func (p *Peer) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "send_start_unix":
		p.SendStart = time.Unix(p.SendStartUnix, 0).Local()
	case "recv_start_unix":
		p.RecvStart = time.Unix(p.RecvStartUnix, 0).Local()
	}
}

func (p *Peer) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	has, err := x.Exist(&Peer{PeerID: p.PeerID})
	if err != nil {
		return err
	}

	if has {
		return nil
	}

	_, err = x.Insert(p)
	if err != nil {
		return err
	}

	return nil
}

type PeerOption struct {
	Offset, Limit int
}

func Peers(chainID string, opt *PeerOption) ([]*Peer, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var ps = make([]*Peer, 0)

	sess := x.NewSession()
	if opt != nil {
		if opt.Limit > 0 {
			sess = sess.Limit(opt.Limit, opt.Offset)
		}
	}

	return ps, sess.OrderBy("created_at desc").Find(&ps)
}
