package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type BlockValidator struct {
	Id               int64     `xorm:"pk autoincr BIGINT"`
	ChainId          string    `xorm:"-"`
	Height           int64     `xorm:"unique(height_val_address_idx) BIGINT"`
	ValidatorAddress string    `xorm:"unique(height_val_address_height_idx) index(val_address_idx) TEXT"`
	ValidatorIndex   int64     `xorm:"BIGINT"`
	Type             int64     `xorm:"BIGINT"`
	Round            int64     `xorm:"BIGINT"`
	Signature        string    `xorm:"TEXT"`
	VotingPower      int64     `xorm:"BIGINT"`
	Accum            int64     `xorm:"BIGINT"`
	Time             time.Time `xorm:"-"`
	TimeUnix         int64
}

func (bv *BlockValidator) BeforeInsert() {
	bv.TimeUnix = bv.Time.Unix()
}

func (bv *BlockValidator) BeforeUpdate() {
	bv.TimeUnix = bv.Time.Unix()
}

func (bv *BlockValidator) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "time_unix":
		bv.Time = time.Unix(bv.TimeUnix, 0).Local()
	}
}

func (bv *BlockValidator) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(bv)
	if err != nil {
		return err
	}

	return nil
}

type BlockValidatorOption struct {
	MinHeight        int64
	MaxHeight        int64
	Height           int64
	ValidatorAddress string
	Offset, Limit    int
}

func BlockValidators(chainID string, opt *BlockValidatorOption) ([]*BlockValidator, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var bvs = make([]*BlockValidator, 0)

	sess := x.NewSession()
	if opt != nil {
		if opt.Height != 0 {
			sess = sess.Where("height = ?", opt.Height)
		}

		if opt.MinHeight != 0 && opt.MaxHeight != 0 {
			sess = sess.Where("height >= ?", opt.MinHeight).Where("height <= ?", opt.MaxHeight)
		}

		if opt.ValidatorAddress != "" {
			sess = sess.Where("json_tx = ?", opt.ValidatorAddress)
		}

		if opt.Limit > 0 {
			sess = sess.Limit(opt.Limit, opt.Offset)
		}
	}

	return bvs, sess.OrderBy("height desc").Find(&bvs)
}
