package service

import (
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/types"
)

// Proposal proposal查询
func (n Node) Proposals() (*types.ResultProposals, error) {
	var result types.ResultProposals
	proposals, err := qos.NewQosCli("").QueryProposals(n.BaseURL)
	if err != nil {
		return nil, err
	}
	for _, proposal := range proposals {
		//deposite, err := strconv.ParseInt(proposal.TotalDeposit.String(), 10, 64)
		//if err != nil {
		//	return nil, err
		//}
		result.Proposals = append(result.Proposals, &proposal)
	}
	return &result, err
}

func (n Node) ProposalByID(id int64) (*types.ResultProposal, error) {
	proposal, err := qos.NewQosCli("").QueryProposal(n.BaseURL, id)
	if err != nil {
		return nil, err
	}
	//deposite, err := strconv.ParseInt(proposal.TotalDeposit.String(), 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	return &proposal, err
}

func (n Node) VotesByID(id int64) (*types.ResultVotes, error) {
	var result types.ResultVotes
	votes, err := qos.NewQosCli("").QueryVotes(n.BaseURL, id)
	if err != nil {
		return nil, err
	}
	for _, vote := range votes {
		v := types.ResultVote{
			Voter: vote.Voter.String(),
			Option: vote.Option.String(),
		}
		result.Votes = append(result.Votes, &v)
	}
	return &result, err
}

func (n Node) DepositesByID(id int64) (*types.ResultDeposits, error) {
	var result types.ResultDeposits
	deposits, err := qos.NewQosCli("").QueryDeposits(n.BaseURL, id)
	if err != nil {
		return nil, err
	}
	for _, deposit := range deposits {
		v := types.ResultDeposit{
			Depositor: deposit.Depositor.String(),
			Amount: deposit.Amount.String(),
		}
		result.Deposits = append(result.Deposits, &v)
	}
	return &result, err
}


//func convertToProposal(mp *models.Proposal) *types.ResultProposal {
//	return &types.ResultProposal{
//		ProposalID:      mp.ProposalID,
//		Title:           mp.Title,
//		Description:     mp.Description,
//		Type:            mp.Type,
//		Status:          mp.Status,
//		SubmitTime:      types.ResultTime(mp.SubmitTime),
//		VotingStartTime: types.ResultTime(mp.VotingStartTime),
//		VotingEndTime:   types.ResultTime(mp.VotingEndTime),
//		TotalDeposit:    mp.TotalDeposit,
//	}
//}
