package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Block struct {
	Id              int64     `xorm:"pk autoincr BIGINT"`
	ChainId         string    `xorm:"-"`
	Height          int64     `xorm:"unique(blocks_height_idx) BIGINT"`
	NumTxs          int64     `xorm:"BIGINT"`
	TotalTxs        int64     `xorm:"BIGINT"`
	DataHash        string    `xorm:"TEXT"`
	ValidatorsNum   int64     `xorm:"BIGINT"`
	ValidatorsTotal int64     `xorm:"BIGINT"`
	ValidatorsHash  string    `xorm:"TEXT"`
	ProposerAddress string    `xorm:"index(proposer_idx) TEXT"`
	Time            time.Time `xorm:"-"`
	TimeUnix        int64
}

func (b *Block) BeforeInsert() {
	b.TimeUnix = b.Time.Unix()
}

func (b *Block) BeforeUpdate() {
	b.TimeUnix = b.Time.Unix()
}

func (b *Block) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "time_unix":
		b.Time = time.Unix(b.TimeUnix, 0).Local()
	}
}

func (b *Block) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(b)
	if err != nil {
		return err
	}

	return nil
}

func (b *Block) InsertIfNotExist(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}
	has, err := x.Exist(b)
	if err != nil {
		return err
	}

	if !has {
		_, err = x.Insert(b)
		if err != nil {
			return err
		}
	}
	return nil
}

type BlockOption struct {
	MinHeight     int64
	MaxHeight     int64
	Height        int64
	NumTxs        int64
	Offset, Limit int
}

func Blocks(chainID string, opt *BlockOption) ([]*Block, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var bvs = make([]*Block, 0)

	sess := x.NewSession()
	//defer sess.Close()
	if opt != nil {
		if opt.Height != 0 {
			sess = sess.Where("height = ?", opt.Height)
		}

		if opt.MinHeight != 0 && opt.MaxHeight != 0 {
			sess = sess.Where("height >= ?", opt.MinHeight).Where("height <= ?", opt.MaxHeight)
		}

		if opt.NumTxs != 0 {
			sess = sess.Where("num_txs != 0")
		}
	}

	if opt.Limit > 0 {
		sess = sess.Limit(opt.Limit, opt.Offset)
	}

	return bvs, sess.OrderBy("height desc").Find(&bvs)
}

func BlocksByProposer(chainID string, proposerAdd string) ([]*Block, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var bvs = make([]*Block, 0)
	sess := x.NewSession()
	//defer sess.Close()
	sess = sess.Where("proposer_address = ?", proposerAdd).Limit(10)
	return bvs, sess.OrderBy("height desc").Find(&bvs)
}