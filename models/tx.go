package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type Tx struct {
	Id          int64  `xorm:"pk autoincr BIGINT"`
	ChainId     string `xorm:"-"`
	Height      int64  `xorm:"index(txs_height_idx) unique(txs_height_index_idx) BIGINT"`
	TxType      string `xorm:"TEXT"`
	Index       int64  `xorm:"unique(txs_height_index_idx) BIGINT"`
	Hash        string
	Maxgas      int64 `xorm:"BIGINT"`
	GasWanted   int64
	GasUsed     int64
	QcpFrom     string    `xorm:"TEXT"`
	QcpTo       string    `xorm:"TEXT"`
	QcpSequence int64     `xorm:"BIGINT"`
	QcpTxindex  int64     `xorm:"BIGINT"`
	QcpIsresult bool      `xorm:"BOOL"`
	TxStatus    int       `xorm:"default 0 INTEGER"`
	OriginTx    string    `xorm:"TEXT"`
	JsonTx      string    `xorm:"TEXT"`
	Time        time.Time `xorm:"-"`
	TimeUnix    int64
}

func (t *Tx) BeforeInsert() {
	t.TimeUnix = t.Time.Unix()
}

func (t *Tx) BeforeUpdate() {
	t.TimeUnix = t.Time.Unix()
}

func (t *Tx) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "time_unix":
		t.Time = time.Unix(t.TimeUnix, 0).Local()
	}
}

func (t *Tx) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(t)
	if err != nil {
		return err
	}

	return nil
}

type TxOption struct {
	MinHeight     int64
	MaxHeight     int64
	Address       string
	Offset, Limit int
}

func Txs(chainID string, opt *TxOption) ([]*Tx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var txs = make([]*Tx, 0)

	sess := x.NewSession()
	if opt.MinHeight != 0 && opt.MaxHeight != 0 {
		sess = sess.Where("height >= ?", opt.MinHeight).Where("height <= ?", opt.MaxHeight)
	}

	if opt.Address != "" {
		sess = sess.Where("json_tx like ?", "%"+opt.Address+"%")
	}

	if opt.Limit > 0 {
		sess = sess.Limit(opt.Limit, opt.Offset)
	}

	return txs, sess.OrderBy("height desc").Find(&txs)
}

func TxByHeightIndex(chainID string, height, index int64) (*Tx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	tx := &Tx{Height: height, Index: index}
	has, err := x.Get(tx)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Tx"}
	}

	return tx, nil
}
