package models

import "github.com/QOSGroup/qmoon/models/errors"

type Inflation struct {
	Height int64 `xorm:"BIGINT"`
	Tokens int64 `xorm:"BIGINT"`
}

func InflationByHeight(chainID string, height int64) (*Inflation, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	inf := &Inflation{Height: height}
	has, err := x.Get(inf)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Inflation at " + string(height)}
	}

	return inf, nil
}

func (in *Inflation) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(in)
	if err != nil {
		return err
	}

	return nil
}

func (inf *Inflation) InsertOrUpdate(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = InflationByHeight(chainID, inf.Height)
	if err != nil {
		_, err = x.Insert(inf)
		if err != nil {
			return err
		}

		return nil
	}
	_, err = x.Update(inf)
	if err != nil {
		return err
	}

	return nil
}
