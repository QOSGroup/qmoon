package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"time"
)

type Proposal struct {
	ProposalContent   ProposalContent `json:"proposal_content"`   // Proposal content interface
	ProposalID        int64           `json:"proposal_id"`        //  ID of the proposal
	Type		string `json:"type"`
	Status            string          `json:"proposal_status"`    //  Status of the Proposal {Pending, Active, Passed, Rejected}
	FinalTallyResult  TallyResult     `json:"final_tally_result"` //  Result of Tallys
	SubmitTime        time.Time       `json:"submit_time"`        //  Time of the block where TxGovSubmitProposal was included
	DepositEndTime    time.Time       `json:"deposit_end_time"`   // Time that the Proposal would expire if deposit amount isn't met
	TotalDeposit      string          `json:"total_deposit"`      //  Current deposit on this proposal. Initial value is set at InitialDeposit
	VotingStartTime   time.Time       `json:"voting_start_time"`  //  Time of the block where MinDeposit was reached. -1 if MinDeposit is not reached
	VotingStartHeight int64           `json:"voting_start_height"`
	VotingEndTime     time.Time       `json:"voting_end_time"` // Time that the VotingPeriod for this proposal will end and votes will be tallied
}

type ProposalContent struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	Deposit     string `json:"deposit"`
}

type ProposalContentValue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deposit     string `json:"deposit"`
}

// Tally Results
type TallyResult struct {
	Yes        string `json:"yes"`
	Abstain    string `json:"abstain"`
	No         string `json:"no"`
	NoWithVeto string `json:"no_with_veto"`
}