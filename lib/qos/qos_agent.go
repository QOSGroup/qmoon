package qos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/lib/qos/gov/types"
)

type QosCli struct {
	remote string
}

func NewQosCli(remote string) QosCli {
	if remote == "" {
		remote = "http://localhost:19527"
	}
	return QosCli{remote: remote}
}

func (cc QosCli) QueryProposals(nodeUrl string) ([]types.Proposal, error) {
	resp, err := http.Get(cc.remote + "/gov/proposals?node_url=" + nodeUrl)
	if err != nil {
		return nil, err
	}

	var result []types.Proposal
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cc QosCli) QueryProposal(nodeUrl string, pId int64) (proposal types.Proposal, err error) {
	resp, err := http.Get(cc.remote + "/gov/proposal?node_url=" + nodeUrl + "&pId=" + strconv.FormatInt(pId, 10))
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&proposal)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryVotes(nodeUrl string, pId int64) ([]types.Vote, error) {
	resp, err := http.Get(cc.remote + "/gov/votes?node_url=" + nodeUrl + "&pId=" + strconv.FormatInt(pId, 10))
	if err != nil {
		return nil, err
	}

	var result []types.Vote
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cc QosCli) QueryDeposits(nodeUrl string, pId int64) ([]types.Deposit, error) {
	resp, err := http.Get(cc.remote + "/gov/deposits?node_url=" + nodeUrl + "&pId=" + strconv.FormatInt(pId, 10))
	if err != nil {
		return nil, err
	}

	var result []types.Deposit
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cc QosCli) QueryTally(nodeUrl string, pId int64) (result types.TallyResult, err error) {
	resp, err := http.Get(cc.remote + "/gov/tally?node_url=" + nodeUrl + "&pId=" + strconv.FormatInt(pId, 10))
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}
