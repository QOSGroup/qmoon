package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/QOSGroup/qmoon/types"
	"github.com/go-xorm/xorm"
)

type Evidence struct {
	Id               int64  `xorm:"pk autoincr BIGINT"`
	Height           int64  `xorm:"BIGINT"`
	ValidatorAddress string `xorm:"TEXT"`
	Hash             string `xorm:"TEXT"`
	Data             string `xorm:"TEXT"`

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

func createEvidence(chainID string, t time.Time, e types.Evidence) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	var evidences Evidence
	has, err := x.Where("validator_address = ?", e.Address).Where("height = ?", e.Height).Exist(&evidences)
	if err == nil && has {
		return nil
	}

	me := &Evidence{
		Height:           e.Height,
		ValidatorAddress: e.Address,
		Hash:             e.Hash,
		Data:             e.Data,
		CreatedAt:        t,
	}

	return me.Insert(chainID)
}

func CreateEvidences(chainID string, es types.EvidenceList) (err error) {
	for _, v := range es.Evidences {
		err = createEvidence(chainID, es.Time, v)
	}

	return err
}

func RetrieveEvidenceByHeight(chainID, validator string, height int64) (*Evidence, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	var evidences Evidence
	has, err := x.Where("validator_address = ?", validator).Where("height = ?", height).Get(&evidences)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Evidence"}
	}

	return &evidences, nil
}

func RetrieveEvidences(chainID, validator string) ([]*Evidence, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	var evidences = make([]*Evidence, 0)
	return evidences, x.Where("validator_address = ?", validator).OrderBy("created_at_unix desc").Find(&evidences)
}
