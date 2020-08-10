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
	var vh1 = make([]*validatorUptimeCurrent,0)
	if vh.Height > 10000 {
		err = sess.SQL("select sum(1-status) as \"status_cnt\", count(1) as \"total_cnt\" from validator_history_record where height >= " + strconv.FormatInt(vh.Height-10000, 10) + " and address= '" + vh.Address + "';").Find(&vh1)
	} else {
		err = sess.SQL("select sum(1-status) as \"status_cnt\", count(1) as \"total_cnt\" from validator_history_record where height >= 0 and address= '" + vh.Address + "';").Find(&vh1)
	}
	//err = sess.SQL("select sum(1-status) as \"status_cnt\", count(1) as \"total_cnt\" from validator_history_record where address= ? ", vh.Address).Find(&uptimeCurrent)
	if err == nil && vh1[0].TotalCnt > 0{
		vh.UptimePercent = strconv.FormatFloat(float64((vh1[0].StatusCnt)*100)/float64(vh1[0].TotalCnt), 'f', -10, 64)
	}

	_, err = x.Insert(vh)
	if err != nil {
		return err
	}

	return nil
}

func InsertHistoryAsLast(chainID string, height int64) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}
	sess := x.NewSession()
	defer sess.Close()

	_, err = x.Exec("insert into validator_history_record (height, address, voting_power, total_power, status, uptime_percent) (select "+strconv.FormatInt(height, 10)+", address, voting_power, total_power, status, uptime_percent from validator_history_record where height = " + strconv.FormatInt(height - 1, 10) + ");")
	if err != nil {
		return err
	}

	return nil
}

func ValidatorHistoryByAddress(chainID string, address string, maxId int64, minId int64, limit int) ([]*ValidatorHistoryRecord, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	sess := x.NewSession()
	defer sess.Close()
	if maxId > 0 {
		sess = sess.Where("id < ?", maxId)
	}
	if minId > 0 {
		sess = sess.Where("id > ?", maxId)
	}
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

	// var bvs = make([]*ValidatorHistoryRecord, 0)
	n, err := sess.Exec("delete from validator_history_record " + condition)
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
		// When the result is too small( the network just started) return the most recent records
		if len(bvs) < 2 {
			return QueryValidatorVotingPower(chainID, address, 0, limit)
		}
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
		// When the result is too small( the network just started) return the most recent records
		if len(bvs) < 2 {
			return QueryValidatorUptime(chainID, address, 0, limit)
		}
	} else {
		// err = sess.SQL("select b.time_unix, vh.uptime_percent from block b, validator_history_record vh where vh.address= ? and b.height=vh.height order by b.id desc limit "+strconv.FormatInt(int64(limit), 10), address).Find(&bvs)
		err = sess.SQL("select vh.uptime_percent, b.time_unix from block b, (select height, address, uptime_percent from validator_history_record vh where address= ? order by height desc limit " + strconv.FormatInt(int64(limit), 10) + ") vh where b.height=vh.height;", address).Find(&bvs)
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
