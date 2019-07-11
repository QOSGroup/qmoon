package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Missing struct {
	Id               int64  `xorm:"pk autoincr BIGINT"`
	ValidatorAddress string `xorm:"TEXT"`
	Height           int64  `xorm:"BIGINT"`

	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Missing) BeforeInsert() {
	if n.CreatedAt.IsZero() {
		n.CreatedAtUnix = time.Now().Unix()
	} else {
		n.CreatedAtUnix = n.CreatedAt.Unix()
	}
}

func (n *Missing) BeforeUpdate() {
}

func (n *Missing) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	}
}

func (n *Missing) Insert(chainID string) error {
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

func RetrieveMissings(chainID, validator string) ([]*Missing, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	var missings = make([]*Missing, 0)
	return missings, x.Where("validator_address = ?", validator).OrderBy("created_at_unix desc").Limit(10, 0).Find(&missings)
}
