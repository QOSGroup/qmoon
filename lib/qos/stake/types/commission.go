package types

import "time"

type (
	// Commission defines a commission parameters for a given validator.
	Commission struct {
		CommissionRates `json:"commission_rates" yaml:"commission_rates"`
		UpdateTime      time.Time `json:"update_time"` // the last time the commission rate was changed
	}

	// CommissionRates defines the initial commission rates to be used for creating a
	// validator.
	CommissionRates struct {
		Rate          string `json:"rate"`            // the commission rate charged to delegators, as a fraction
		MaxRate       string `json:"max_rate"`        // maximum commission rate which validator can ever charge, as a fraction
		MaxChangeRate string `json:"max_change_rate"` // maximum daily increase of the validator commission, as a fraction
	}
)
