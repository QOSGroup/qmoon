package service

import (
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

// Proposal proposal查询
func (n Node) Proposals() (*types.ResultProposals, error) {
	var result types.ResultProposals

	mps, err := models.Proposals(n.ChanID)
	if err != nil {
		return nil, err
	}
	if err == nil {
		for _, v := range mps {
			result.Proposals = append(result.Proposals, convertToProposal(v))
		}
	}

	return &result, err
}

func convertToProposal(mp *models.Proposal) *types.ResultProposal {
	return &types.ResultProposal{
		ProposalID:      mp.ProposalID,
		Title:           mp.Title,
		Description:     mp.Description,
		Type:            mp.Type,
		Status:          mp.Status,
		SubmitTime:      mp.SubmitTime,
		VotingStartTime: mp.VotingStartTime,
		TotalDeposit:    mp.TotalDeposit,
	}
}
