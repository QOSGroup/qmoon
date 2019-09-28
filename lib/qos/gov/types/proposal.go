package types

import (
	"math/big"
	"time"
)

type Proposal struct {
	ProposalContent `json:"proposal_content"` // Proposal content interface

	ProposalID int64 `json:"proposal_id"` //  ID of the proposal

	Status           ProposalStatus `json:"proposal_status"`    //  Status of the Proposal {Pending, Active, Passed, Rejected}
	FinalTallyResult TallyResult    `json:"final_tally_result"` //  Result of Tallys

	SubmitTime     time.Time `json:"submit_time"`      //  Time of the block where TxGovSubmitProposal was included
	DepositEndTime time.Time `json:"deposit_end_time"` // Time that the Proposal would expire if deposit amount isn't met
	TotalDeposit   BigInt    `json:"total_deposit"`    //  Current deposit on this proposal. Initial value is set at InitialDeposit

	VotingStartTime   time.Time `json:"voting_start_time"` //  Time of the block where MinDeposit was reached. -1 if MinDeposit is not reached
	VotingStartHeight int64     `json:"voting_start_height"`
	VotingEndTime     time.Time `json:"voting_end_time"` // Time that the VotingPeriod for this proposal will end and votes will be tallied
}

type ProposalContent struct {
	Type  string               `json:"type"`
	Value ProposalContentValue `json:"value"`
}
type ProposalContentValue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deposit     string `json:"deposit"`
}

// Tally Results
type TallyResult struct {
	Yes        BigInt `json:"yes"`
	Abstain    BigInt `json:"abstain"`
	No         BigInt `json:"no"`
	NoWithVeto BigInt `json:"no_with_veto"`
}

type BigInt struct {
	i *big.Int
}

// Int64 converts BigInt to int64
// Panics if the value is out of range
func (i BigInt) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}

type ProposalStatus byte

//nolint
const (
	StatusNil           ProposalStatus = 0x00
	StatusDepositPeriod ProposalStatus = 0x01
	StatusVotingPeriod  ProposalStatus = 0x02
	StatusPassed        ProposalStatus = 0x03
	StatusRejected      ProposalStatus = 0x04
)

func ValidProposalStatus(status ProposalStatus) bool {
	if status == StatusDepositPeriod ||
		status == StatusVotingPeriod ||
		status == StatusPassed ||
		status == StatusRejected {
		return true
	}
	return false
}

// Turns VoteOption byte to String
func (ps ProposalStatus) String() string {
	switch ps {
	case StatusDepositPeriod:
		return "Deposit"
	case StatusVotingPeriod:
		return "Voting"
	case StatusPassed:
		return "Passed"
	case StatusRejected:
		return "Rejected"
	default:
		return ""
	}
}
