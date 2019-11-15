package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type TxTransfer struct {
	Id        int64     `xorm:"pk autoincr BIGINT"`
	ChainId   string    `xorm:"-"`
	Height    int64     `xorm:"BIGINT"`
	Hash      string    `xorm:"TEXT"`
	Address   string    `xorm:"index TEXT"`
	Coin      string    `xorm:"TEXT"`
	Amount    string    `xorm:"TEXT"`
	TxStatus  int       `xorm:"default 0 INTEGER"`
	IsMulti   bool      `xorm:"default false BOOL"`
	MultiData string    `xorm:"TEXT"`
	Type      int       `xorm:"SMALLINT"`
	Time      time.Time `xorm:"-"`
	TimeUnix  int64
}

func (tt *TxTransfer) BeforeInsert() {
	tt.TimeUnix = tt.Time.Unix()
}

func (tt *TxTransfer) BeforeUpdate() {
	tt.TimeUnix = tt.Time.Unix()
}

func (tt *TxTransfer) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "time_unix":
		tt.Time = time.Unix(tt.TimeUnix, 0).Local()
	}
}
func (tt *TxTransfer) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(tt)
	if err != nil {
		return err
	}

	return nil
}

type TxTransferOption struct {
	Address       string
	Coin          string
	Offset, Limit int
}

func TxTransfers(chainID string, opt *TxTransferOption) ([]*TxTransfer, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var bvs = make([]*TxTransfer, 0)

	sess := x.NewSession()
	if opt != nil {
		if opt.Address != "" {
			sess = sess.Where("address = ?", opt.Address)
		}

		if opt.Coin != "" {
			sess = sess.Where("coin = ?", opt.Coin)
		}

		if opt.Limit > 0 {
			sess = sess.Limit(opt.Limit, opt.Offset)
		}
	}

	return bvs, sess.OrderBy("time desc").Find(&bvs)
}
