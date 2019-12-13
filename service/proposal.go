package service

import (
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/types"
)

// Proposal proposal查询
func (n Node) Proposals() (*types.ResultProposals, error) {
	var result types.ResultProposals
	result.Proposals = make([]*types.ResultProposal, 0)
	proposals, err := qos.NewQosCli("").QueryProposals(n.BaseURL)
	if err != nil {
		return nil, err
	}
	for i := len(proposals) - 1; i >= 0; i-- {
		proposal := proposals[i]
		result.Proposals = append(result.Proposals, &proposal)
	}
	return &result, err
}

func (n Node) ProposalByID(id int64) (*types.ResultProposal, error) {
	proposal, err := qos.NewQosCli("").QueryProposal(n.BaseURL, id)
	if err != nil {
		return nil, err
	}

	proposal.Deposits, err = n.DepositesByID(id)
	proposal.Votes, err = n.VotesByID(id)
	//deposite, err := strconv.ParseInt(proposal.TotalDeposit.String(), 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	return &proposal, err
}

func (n Node) VotesByID(id int64) ([]*types.ResultVote, error) {
	result := make([]*types.ResultVote, 0)
	votes, err := qos.NewQosCli("").QueryVotes(n.BaseURL, id)
	if err != nil {
		return nil, err
	}
	for _, vote := range votes {
		v := types.ResultVote{
			Voter: vote.Voter.String(),
			Option: vote.Option.String(),
		}
		result = append(result, &v)
	}
	return result, err
}

func (n Node) DepositesByID(id int64) ([]*types.ResultDeposit, error) {
	result := make([]*types.ResultDeposit, 0)
	deposits, err := qos.NewQosCli("").QueryDeposits(n.BaseURL, id)
	if err != nil {
		return nil, err
	}
	for _, deposit := range deposits {
		v := types.ResultDeposit{
			Depositor: deposit.Depositor.String(),
			Amount: deposit.Amount.String(),
		}
		result = append(result, &v)
	}
	return result, err
}