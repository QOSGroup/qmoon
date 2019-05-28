// Copyright 2019 The QOS Authors

package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

const (
	NetworkStatusValid   = 1
	NetworkStatusInvalid = 2
)

type Network struct {
	ID     int64  `json:"id"`
	Remote string `json:"remote"`
	Delay  int64  `json:"delay"`
	Status int    `json:"status"` // 1: 有效  2: 无效

	LastTestAt     time.Time `xorm:"-"`
	LastTestAtUnix int64

	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Network) BeforeInsert() {
	now := time.Now()
	n.LastTestAtUnix = now.Unix()
	n.CreatedAtUnix = now.Unix()
}

func (n *Network) BeforeUpdate() {
	n.LastTestAtUnix = time.Now().Unix()
}

func (n *Network) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	case "last_test_at_unix":
		n.LastTestAt = time.Unix(n.LastTestAtUnix, 0).Local()
	}
}

func (n *Network) InsertOrUpdate(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	old := &Network{Remote: n.Remote}
	has, err := x.Get(old)
	if err != nil {
		return err
	}

	if has {
		old.LastTestAt = time.Now()
		old.Status = n.Status
		old.Delay = n.Delay
		_, err = x.ID(old.ID).Update(old)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = x.Insert(n)
	if err != nil {
		return err
	}

	return nil
}

func CreateNetworkOrUpdate(chainID string, n *Network) error {
	return n.InsertOrUpdate(chainID)
}

type NetworkOption struct {
	Status        int
	Offset, Limit int
}

func Networks(chainID string, opt *NetworkOption) ([]*Network, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var ns = make([]*Network, 0)

	sess := x.NewSession()
	if opt != nil {
		if opt.Status != 0 {
			sess = sess.Where("status = ?", opt.Status)
		}
		if opt.Limit > 0 {
			sess = sess.Limit(opt.Limit, opt.Offset)
		}
	}

	return ns, sess.OrderBy("delay").Find(&ns)
}
