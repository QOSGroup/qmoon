package models

import (
	"fmt"
	"github.com/QOSGroup/qmoon/types"
	"github.com/shopspring/decimal"
	"strconv"
)

type ValidatorHistoryRecord struct {
	Id                 int64  `xorm:"pk autoincr BIGINT"`
	RecordTime			int64	`xorm:"BIGINT"`
	Address            string    `xorm:"unique TEXT"`
	VotingPower        int64     `xorm:"BIGINT"`
	TotalPower			int64 `xorm:"BIGINT"`
	Status             int       `xorm:"INTEGER"`
}


func (vh *ValidatorHistoryRecord) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}
	total, err := TotalVotingPower(chainID)
	if err!=nil || total==0{
		return err
	}
	vh.TotalPower = total

	_, err = x.Insert(vh)
	if err != nil {
		return err
	}

	return nil
}

func ValidatorHistoryByAddress(chainID string, address string) ([]*ValidatorHistoryRecord, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	sess := x.NewSession()
	defer sess.Close()
	sess = sess.Where("address = ?", address).Desc("record_time").Limit(24*12, 0)

	var bvs = make([]*ValidatorHistoryRecord, 0)

	return bvs, sess.Find(&bvs)
}

func PurgeOldValidatorHistory(chainID string, purgeTime int64) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	sess := x.NewSession()
	defer sess.Close()
	var bvs = make([]*ValidatorHistoryRecord, 0)
	n, err := sess.Where("record_time < ?", purgeTime).Delete(&bvs)
	fmt.Println("Purged old validators' history: ", n)
	return err
}


func QueryValidatorVotingPowerPercent(chainID string, address string) ([]types.Matrix, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	sess := x.NewSession()
	defer sess.Close()

	var bvs = make([]*ValidatorHistoryRecord, 0)
	var result = make([]types.Matrix, 0)
	err = sess.Where(" address = ? ", address).Find(&bvs)
	if err != nil {
		return nil, err
	}
	for _, vh :=  range bvs {
		a,_ := decimal.NewFromString(strconv.FormatInt(vh.VotingPower * 100, 10))
		b,_ := decimal.NewFromString(strconv.FormatInt(vh.TotalPower,10))
		percent,_:= a.Div(b).Float64()
		result = append(result, types.Matrix{
			X:strconv.FormatInt(vh.RecordTime, 10),
			Y:strconv.FormatFloat(percent,'f', -4, 32),
		})
	}
	return result, err
}

func QueryValidatorUptime(chainID string, address string)([]types.Matrix, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	sess := x.NewSession()
	defer sess.Close()
	var bvs = make([]*ValidatorHistoryRecord, 0)
	var result = make([]types.Matrix, 0)
	err = sess.Where(" address = ? ", address).Find(&bvs)
	if err != nil {
		return nil, err
	}
	for _, vh :=  range bvs {
		result = append(result, types.Matrix{
			X:strconv.FormatInt(vh.RecordTime, 10),
			Y:strconv.FormatInt(int64(1-vh.Status), 10),
		})
	}
	return result, err
}