package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Proposal struct {
	ID              int64     `xorm:"pk autoincr BIGINT"`
	ProposalID      int64     `xorm:"BIGINT"`
	Title           string    `xorm:"unique(title_tx_idx) TEXT"`
	Description     string    `xorm:"TEXT"`
	Type            string    `xorm:"TEXT"`
	Status          string    `xorm:"TEXT"`
	SubmitTime      time.Time `xorm:"-"`
	VotingStartTime time.Time `xorm:"-"`
	TotalDeposit    int64     `xorm:"BIGINT"`
	ChainID         string    `xorm:"-"`
	UpdatedAt       time.Time `xorm:"-"`
	UpdatedAtUnix   int64
	CreatedAt       time.Time `xorm:"-"`
	CreatedAtUnix   int64
}

func (n *Proposal) BeforeInsert() {
	n.CreatedAtUnix = time.Now().Unix()
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Proposal) BeforeUpdate() {
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Proposal) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		n.UpdatedAt = time.Unix(n.UpdatedAtUnix, 0).Local()
	}
}

func (n *Proposal) Insert(chainID string) error {
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

func (pro *Proposal) Update(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.ID(pro.Id).Update(pro)
	if err != nil {
		return err
	}

	return nil
}

func (pro *Proposal) UpdateStatus(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.ID(pro.Id).
		Cols("status").
		Update(pro)
	if err != nil {
		return err
	}

	return nil
}

func Proposals(chainID string) ([]*Proposal, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var pros = make([]*Proposal, 0)
	return pros, x.Find(&pros)
}
