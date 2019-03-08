package models

import (
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

const DefaultFeeTx = "default"

type Fee struct {
	Id        int64  `xorm:"pk autoincr BIGINT"`
	Tx        string `xorm:"unique(fee_tx_idx) TEXT"`
	GasWanted int64  `xorm:"BIGINT"`
	GasUsed   int64  `xorm:"BIGINT"`
	Fee       string `xorm:"TEXT"`
	ChainID   string `xorm:"-"`

	Version       int       `xorm:"version"`
	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Fee) BeforeInsert() {
	n.CreatedAtUnix = time.Now().Unix()
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Fee) BeforeUpdate() {
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Fee) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		n.UpdatedAt = time.Unix(n.UpdatedAtUnix, 0).Local()
	}
}

func syncDefaultFee(chainID, fee string, gasWanted, gasUsed int64) error {
	return updateFee(chainID, DefaultFeeTx, fee, gasWanted, gasUsed)
}

func (n *Fee) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(n)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFee(chainID, tx, fee string, gasWanted, gasUsed int64) error {
	if err := updateFee(chainID, tx, fee, gasWanted, gasUsed); err != nil {
		return err
	}

	if err := syncDefaultFee(chainID, fee, gasWanted, gasUsed); err != nil {
		logrus.Warnf("syncDefaultFee err:%s", err.Error())
	}

	return updateFee(chainID, tx, fee, gasWanted, gasUsed)
}

func updateFee(chainID, tx, fee string, gasWanted, gasUsed int64) error {
	tx = strings.ToLower(tx)

	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	f, err := RetrieveFeeByTx(chainID, tx)
	if err != nil {
		if errors.IsNotExist(err) {
			f = &Fee{}
			f.Tx = tx
			f.GasWanted = gasWanted
			f.GasUsed = gasUsed
			f.Fee = fee

			if err := f.Insert(chainID); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		f.GasWanted = gasWanted
		f.GasUsed = gasUsed
		f.Fee = fee
		if _, err := x.ID(f.Id).Update(f); err != nil {
			return err
		}
	}

	return nil
}

func RetrieveFeeByTx(chainID, tx string) (*Fee, error) {
	tx = strings.ToLower(tx)
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	f := &Fee{Tx: tx}
	has, err := x.Get(f)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Fee"}
	}

	f.ChainID = chainID

	return f, nil
}
