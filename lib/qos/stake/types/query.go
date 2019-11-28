package types

import qtypes "github.com/QOSGroup/qmoon/types"

type DelegationQueryResult struct {
	DelegatorAddr            string `json:"delegator_address"`
	ValidatorAddr            string `json:"validator_address"`
	ValidatorConsensusPubKey string `json:"validator_cons_pub_key"`
	Amount                   string `json:"delegate_amount"`
	IsCompound               bool   `json:"is_compound"`
}

type Delegations []DelegationQueryResult

type DelegationQueryWithValidatorResult struct {
	DelegatorAddr            string `json:"delegator_address"`
	Validator           	 *qtypes.Validator `json:"validator"`
	ValidatorConsensusPubKey string `json:"validator_cons_pub_key"`
	Amount                   string `json:"delegate_amount"`
	IsCompound               bool   `json:"is_compound"`
}

type DelegationsWithValidator []DelegationQueryWithValidatorResult
