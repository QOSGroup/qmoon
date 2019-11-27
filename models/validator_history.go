package models

import (
	"fmt"
	"github.com/QOSGroup/qmoon/types"
	"github.com/shopspring/decimal"
	"strconv"
)

type ValidatorHistoryRecord struct {
	Id                 int64  `xorm:"pk autoincr BIGINT"`
	Height			int64	`xorm:"unique(height_address_idx) BIGINT"`
	Address            string    `xorm:"unique(height_address_idx) index(address_idx) TEXT"`
	VotingPower        int64     `xorm:"BIGINT"`
	TotalPower			int64 `xorm:"BIGINT"`
	Status             int       `xorm:"INTEGER"`
	UptimePercent	  string     `xorm:"uptime_percent"`
}

type ValidatorUptimeResult struct {
	TimeUnix int64
	UptimePercent string
}

type ValidatorVotingPowerResult struct {
	TimeUnix int64
	VotingPower int64
	TotalPower int64
}

type validatorUptimeCurrent struct {
	StatusCnt int64
	TotalCnt int64
}

func (vh *ValidatorHistoryRecord) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}
	sess := x.NewSession()
	defer sess.Close()

	vh1 := ValidatorHistoryRecord{Address:vh.Address}
	totalCnt, err := sess.Count(vh1)
	statusCnt, err := sess.SumInt(vh1, "status")
	//err = sess.SQL("select sum(1-status) as \"status_cnt\", count(1) as \"total_cnt\" from validator_history_record where address= ? ", vh.Address).Find(&uptimeCurrent)
	if err == nil && totalCnt > 0{
		vh.UptimePercent = strconv.FormatFloat(float64((totalCnt-statusCnt)*100)/float64(totalCnt), 'f', -10, 64)
	}

	_, err = x.Insert(vh)
	if err != nil {
		return err
	}

	return nil
}

func ValidatorHistoryByAddress(chainID string, address string, limit int) ([]*ValidatorHistoryRecord, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	sess := x.NewSession()
	defer sess.Close()
	sess = sess.Where("address = ?", address).Desc("height").Limit(limit)

	var bvs = make([]*ValidatorHistoryRecord, 0)

	return bvs, sess.Find(&bvs)
}

func PurgeOldValidatorHistory(chainID string, condition string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	sess := x.NewSession()
	defer sess.Close()
	var bvs = make([]*ValidatorHistoryRecord, 0)
	n, err := sess.Where(condition).Delete(&bvs)
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

	var bvs = make([]*ValidatorVotingPowerResult, 0)
	var result = make([]types.Matrix, 0)
	// err = sess.Where(" address = ? ", address).Find(&bvs)
	err = sess.SQL("select b.time_unix, vh.voting_power, vh.total_power from block b, validator_history_record vh where vh.address= ? and b.height=vh.height", address).Find(&bvs)
	if err != nil {
		return nil, err
	}
	for _, vh :=  range bvs {
		a,_ := decimal.NewFromString(strconv.FormatInt(vh.VotingPower * 100, 10))
		b,_ := decimal.NewFromString(strconv.FormatInt(vh.TotalPower,10))
		y := "Not Available"
		if !b.IsZero() {
			percent,_:= a.Div(b).Float64()
			y = strconv.FormatFloat(percent,'f', -2, 64)
		}
		result = append(result, types.Matrix{
			X:strconv.FormatInt(vh.TimeUnix, 10),
			Y:y,
		})
	}
	return result, err
}

func QueryValidatorVotingPower(chainID string, address string, intervals int64, limit int) ([]types.Matrix, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	sess := x.NewSession()
	defer sess.Close()

	var bvs = make([]*ValidatorVotingPowerResult, 0)
	var result = make([]types.Matrix, 0)
	if intervals > 0 {
		err = sess.SQL("select b.time_unix, vh.voting_power, vh.total_power from block b, validator_history_record vh where vh.address= ? and vh.height % ? = 0 and b.height=vh.height order by b.id desc limit "+strconv.FormatInt(int64(limit), 10), address, intervals).Find(&bvs)
	} else {
		err = sess.SQL("select b.time_unix, vh.voting_power, vh.total_power from block b, validator_history_record vh where vh.address= ? and b.height=vh.height order by b.id desc limit "+strconv.FormatInt(int64(limit), 10), address).Find(&bvs)
	}

	if err != nil {
		return nil, err
	}
	for _, vh :=  range bvs {
		result = append(result, types.Matrix{
			X:strconv.FormatInt(vh.TimeUnix, 10),
			Y:strconv.FormatFloat(float64(vh.VotingPower * 100)/float64(vh.TotalPower), 'f', -10,64),
		})
	}
	return result, err
}

func QueryValidatorUptime(chainID string, address string, intervals int64, limit int)(result []types.Matrix, err error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	sess := x.NewSession()
	defer sess.Close()
	var bvs = make([]*ValidatorUptimeResult, 0)
	result = make([]types.Matrix, 0)
	if intervals > 0 {
		err = sess.SQL("select b.time_unix, vh.uptime_percent from block b, validator_history_record vh where vh.address= ? and vh.height % ? = 0 and b.height=vh.height order by b.id desc limit "+strconv.FormatInt(int64(limit), 10), address, intervals).Find(&bvs)
	} else {
		err = sess.SQL("select b.time_unix, vh.uptime_percent from block b, validator_history_record vh where vh.address= ? and b.height=vh.height order by b.id desc limit "+strconv.FormatInt(int64(limit), 10), address).Find(&bvs)
	}
	if err != nil {
		return nil, err
	}
	for _, vh :=  range bvs {
		result = append(result, types.Matrix{
			X:strconv.FormatInt(vh.TimeUnix, 10),
			Y:vh.UptimePercent,
		})
	}
	return result, err
}