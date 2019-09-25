package types

import (
	"math/big"
	"time"
)

// Description - description fields for a validator
type Description struct {
	Moniker  string `json:"moniker" yaml:"moniker"`   // name
	Identity string `json:"identity" yaml:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website" yaml:"website"`   // optional website link
	Details  string `json:"details" yaml:"details"`   // optional details
}

type (
	// Commission defines a commission parameters for a given validator.
	Commission struct {
		CommissionRates `json:"commission_rates" yaml:"commission_rates"`
		UpdateTime      time.Time `json:"update_time" yaml:"update_time"` // the last time the commission rate was changed
	}

	// CommissionRates defines the initial commission rates to be used for creating a
	// validator.
	CommissionRates struct {
		Rate          Dec `json:"rate" yaml:"rate"`                       // the commission rate charged to delegators, as a fraction
		MaxRate       Dec `json:"max_rate" yaml:"max_rate"`               // maximum commission rate which validator can ever charge, as a fraction
		MaxChangeRate Dec `json:"max_change_rate" yaml:"max_change_rate"` // maximum daily increase of the validator commission, as a fraction
	}
)

type Dec struct {
	*big.Int `json:"int"`
}

type Validator struct {
	OperatorAddress         string      `json:"operator_address"`
	ConsPubKey              string      `json:"consensus_pubkey"`
	Jailed                  bool        `json:"jailed"`
	Status                  int         `json:"status"`
	Tokens                  string      `json:"tokens"`
	DelegatorShares         string      `json:"delegator_shares"`
	Description             Description `json:"description"`
	UnbondingHeight         string      `json:"unbonding_height"`
	UnbondingCompletionTime string      `json:"unbonding_time"`
	Commission              Commission  `json:"commission"`
	MinSelfDelegation       string      `json:"min_self_delegation"`
}
