package types

import "time"

type ValidatorDisplayInfo struct {
	OperatorAddress string      `json:"validator"`
	Owner           string      `json:"owner"`
	ConsAddress     string      `json:"consensusAddress"`
	ConsPubKey      string      `json:"consensusPubKey"`
	BondedTokens    string      `json:"bondTokens"`
	Description     Description `json:"description"`
	Commission      Commission  `json:"commission"`

	Status         string    `json:"status"`
	InactiveDesc   string    `json:"InactiveDesc"`
	InactiveTime   time.Time `json:"inactiveTime"`
	InactiveHeight int64     `json:"inactiveHeight"`

	MinPeriod  int64  `json:"minPeriod"`
	BondHeight int64  `json:"bondHeight"`
	SelfBond   string `json:"selfBond"`
}

// Description - description fields for a validator
type Description struct {
	Moniker string `json:"moniker"` // name
	Logo    string `json:"logo"`    // optional logo link
	Website string `json:"website"` // optional website link
	Details string `json:"details"` // optional details
}
