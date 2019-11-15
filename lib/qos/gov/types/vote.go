package types

import "fmt"

// Vote
type Vote struct {
	Voter      string     `json:"voter"`       //  address of the voter
	ProposalID int64      `json:"proposal_id"` //  proposalID of the proposal
	Option     string `json:"option"`      //  option from OptionSet chosen by the voter
}

func (v Vote) String() string {
	return fmt.Sprintf("Voter %s voted with option %s on proposal %d", v.Voter, v.Option, v.ProposalID)
}

// Votes is a collection of Vote
type Votes []Vote

type VoteOption byte