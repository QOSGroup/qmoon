package qos

import (
	"encoding/json"
	// "github.com/cosmos/cosmos-sdk/x/gov"
	//"github.com/QOSGroup/qmoon/lib/qos/gov/types"
	mint_types "github.com/QOSGroup/qmoon/lib/qos/mint/types"
	stake_types "github.com/QOSGroup/qmoon/lib/qos/stake/types"
	base_types "github.com/QOSGroup/qmoon/lib/qos/types"
	gov_types "github.com/QOSGroup/qmoon/lib/qos/gov/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tm_types "github.com/tendermint/tendermint/types"
	"github.com/QOSGroup/qmoon/models/errors"
	"net/http"
	"strconv"
)

type QosCli struct {
	remote string
}

func NewQosCli(remote string) QosCli {
	if remote == "" {
		remote = "http://localhost:19528"
	}
	return QosCli{remote: remote}
}

func (cc QosCli) QueryTx(nodeUrl, tx string) (result base_types.TxResponse, err error) {
	resp, err := http.Get(cc.remote + "/block/tx?node_url=" + nodeUrl + "&tx=" + tx)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (cc QosCli) QueryProposals(nodeUrl string) (result []gov_types.Proposal, err error) {
	resp, err := http.Get(cc.remote + "/gov/proposals?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (cc QosCli) QueryProposal(nodeUrl string, pId int64) (proposal gov_types.Proposal, err error) {
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

func (cc QosCli) QueryVotes(nodeUrl string, pId int64) ([]gov_types.Vote, error) {
	resp, err := http.Get(cc.remote + "/gov/votes?node_url=" + nodeUrl + "&pId=" + strconv.FormatInt(pId, 10))
	if err != nil {
		return nil, err
	}

	var result []gov_types.Vote
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cc QosCli) QueryDeposits(nodeUrl string, pId int64) ([]gov_types.Deposit, error) {
	resp, err := http.Get(cc.remote + "/gov/deposits?node_url=" + nodeUrl + "&pId=" + strconv.FormatInt(pId, 10))
	if err != nil {
		return nil, err
	}

	var result []gov_types.Deposit
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cc QosCli) QueryTally(nodeUrl string, pId int64) (result gov_types.TallyResult, err error) {
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

func (cc QosCli) QueryCommunityFeePool(nodeUrl string) (result string, err error) {
	resp, err := http.Get(cc.remote + "/distribution/community/fee/pool?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryInflationPhrases(nodeUrl string) (result []mint_types.InflationPhrase, err error) {
	resp, err := http.Get(cc.remote + "/mint/inflation/phrases?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryTotal(nodeUrl string) (result string, err error) {
	resp, err := http.Get(cc.remote + "/mint/total?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryApplied(nodeUrl string) (result string, err error) {
	resp, err := http.Get(cc.remote + "/mint/Aapplied?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryValidators(nodeUrl string) (result []stake_types.ValidatorDisplayInfo, err error) {
	resp, err := http.Get(cc.remote + "/stake/validators?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryTotalValidatorBondTokens(nodeUrl string) (result string, err error) {
	resp, err := http.Get(cc.remote + "/stake/validators/total/bond/tokens?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryDelegationsWithValidator(nodeUrl, validator string) (result []stake_types.DelegationQueryResult, err error) {
	resp, err := http.Get(cc.remote + "/stake/validator/delegations?node_url=" + nodeUrl + "&validator=" + validator)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryBlockByHeight(nodeUrl string, height int64) (result *tm_types.Block, err error) {
	resp, err := http.Get(cc.remote + "/block?node_url=" + nodeUrl + "&height=" + strconv.FormatInt(height, 10))
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (cc QosCli) QueryStatus(nodeUrl string) (result *ctypes.SyncInfo, err error) {
	status, err := http.Get(cc.remote + "/status?node_url=" + nodeUrl)
	if err != nil {
		return
	}

	err = json.NewDecoder(status.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}
