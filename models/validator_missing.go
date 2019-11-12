package models

import (
	"fmt"
	"github.com/QOSGroup/qmoon/types"
)

const MISSINGLENGTH = 3600

type ValidatorMissing struct {
	ValidatorAddress string `xorm:"pk TEXT"`
	Status           int8   `xorm:"INT"`
	Votes            string `xorm:"BIT"`
}

type ValidatorMissingHeight struct {
	RefreshHeight int64 `xorm:"BIGINT"`
	UpdateHeight  int64 `xorm:"BIGINT"`
}

func (val *ValidatorMissing) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(val)
	if err != nil {
		return err
	}

	return nil
}

func getValidatorMissings(chainID string, validator string, limit int) (ValidatorMissing, error) {
	var missings = ValidatorMissing{}
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return missings, err
	}

	err = x.Sql("select validator_address, status, cast(votes, bit(?)) from validator_missing where validator_address = ?", limit, validator).Find(&missings)
	return missings, err
}

func getLatestMissingHeight(chainID string) (ValidatorMissingHeight, error) {
	var heights = ValidatorMissingHeight{}
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return heights, err
	}

	err = x.Sql("select refresh_height, update_height from validator_missing_height").Find(&heights)
	return heights, err
}

func GetValidatorRecentMissingHeights(chainID string, validator string, limit int) ([]int64, error) {
	if limit > MISSINGLENGTH {
		//return nil, types.NewExceedMaxlengthError("Missing heights limit", string(limit), string(MISSINGLENGTH))
		limit = MISSINGLENGTH
	}
	missings, err := getValidatorMissings(chainID, validator, limit)
	if err != nil {
		return nil, err
	}
	lastHeight, err := getLatestMissingHeight(chainID)
	if err != nil {
		return nil, err
	}
	missed_heights := make([]int64, 0)
	for i := 0; i < limit; i++ {
		votes := missings.Votes
		voted := votes[i : i+1]
		if voted == "0" {
			missed_heights = append(missed_heights, lastHeight.UpdateHeight-int64(i))
		}
	}
	return missed_heights, nil
}

func GetMissedValidatorsByHeight(chainID string, height int64) ([]*Validator, error) {
	latest, err := getLatestMissingHeight(chainID)
	if err != nil {
		return nil, err
	}
	if height > latest.UpdateHeight {
		return nil, types.NewExceedMaxlengthError("Requested height", string(height), string(latest.UpdateHeight))
	}
	if height-latest.UpdateHeight > MISSINGLENGTH {
		return nil, types.NewExceedMaxlengthError("Requested height", string(height), string(MISSINGLENGTH))
	}
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var missings = make([]*ValidatorMissing, 0)
	err = x.Where("get_bit(votes, ?)=0", latest.UpdateHeight-height).Find(&missings)
	if err != nil {
		return nil, err
	}
	var missedValidators = make([]*Validator, 0)
	for _, miss := range missings {
		val, err := ValidatorByAddress(chainID, miss.ValidatorAddress)
		if err != nil {
			return nil, err
		}
		missedValidators = append(missedValidators, val)
	}
	return missedValidators, nil
}

func InsertMissings(chainID string, height int64, validators []*BlockValidator) error {
	x, err := GetNodeEngine(chainID)
	defer x.Close()
	if err != nil {
		return err
	}
	var full = "b'1"
	var one = "b'1"
	for i := 1; i < MISSINGLENGTH; i++ {
		full += "1"
		one += "0"
	}
	full += "'"
	one += "'"
	latest, err := getLatestMissingHeight(chainID)

	if height < latest.UpdateHeight {

		return types.NewExceedMaxlengthError("Insert height duped", string(height), string(latest.UpdateHeight))
	}
	if height-latest.UpdateHeight > 1 {
		fmt.Println("validator_missing skipped " + string(height-latest.UpdateHeight-1))
	}
	_, err = x.Exec("update validator_missing set votes=votes>>1")
	if err != nil {
		return err
	}
	var valueString = ""
	for _, bv := range validators {
		valueString += "('" + bv.ValidatorAddress + "', 0, b'" + one + "'), "
	}
	valueString = valueString[0 : len(valueString)-2]
	_, err = x.Exec("validator_missing values " + valueString + " on conflict (validator_address)	do update set votes=validator_missing.votes|excluded.votes;")
	if err != nil {
		return err
	}
	return nil
}
