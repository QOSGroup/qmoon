package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Evidence struct {
	Id               int64  `xorm:"pk autoincr BIGINT"`
	ValidatorAddress string `xorm:"TEXT"`
	Height           int64  `xorm:"BIGINT"`

	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Evidence) BeforeInsert() {
	if n.CreatedAt.IsZero() {
		n.CreatedAtUnix = time.Now().Unix()
	} else {
		n.CreatedAtUnix = n.CreatedAt.Unix()
	}
}

func (n *Evidence) BeforeUpdate() {
}

func (n *Evidence) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	}
}

func (n *Evidence) Insert(chainID string) error {
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

func RetrieveEvidences(chainID, validator string) ([]*Evidence, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	var evidences = make([]*Evidence, 0)
	return evidences, x.Where("validator_address = ?", validator).OrderBy("created_at_unix desc").Find(&evidences)
}
