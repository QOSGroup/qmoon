package types

// Deposit
type Deposit struct {
	Depositor  string `json:"depositor"`   //  Address of the depositor
	ProposalID int64  `json:"proposal_id"` //  proposalID of the proposal
	Amount     string `json:"amount"`      //  Deposit amount
}

type Deposits []Deposit
