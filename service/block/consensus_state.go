// Copyright 2018 The QOS Authors

package block

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func convertToConsensusState(mcs *model.ConsensusState) *types.ResultConsensusState {
	return &types.ResultConsensusState{
		ChainID:         mcs.ChainID.String,
		Height:          mcs.Height.String,
		Round:           mcs.Round.String,
		Step:            mcs.Step.String,
		PrevotesNum:     mcs.PrecommitsNum.Int64,
		PrevotesValue:   mcs.PrevotesValue.String,
		PrecommitsNum:   mcs.PrecommitsNum.Int64,
		PrecommitsValue: mcs.PrecommitsValue.String,
		StartTime:       mcs.StartTime.String,
	}
}

func retrieveConsensusState(chainID string) (*model.ConsensusState, error) {
	return model.ConsensusStateByChainID(db.Db, utils.NullString(chainID))
}

func RetrieveConsensusState(chainID string) (*types.ResultConsensusState, error) {
	mcs, err := retrieveConsensusState(chainID)
	if err != nil {
		return nil, err
	}

	return convertToConsensusState(mcs), nil
}

type tmConsensusState struct {
	Height          string `json:"-"`
	Round           string `json:"-"`
	Step            string `json:"-"`
	PrevotesNum     int64  `json:"-"`
	PrevotesValue   string `json:"-"`
	PrecommitsNum   int64  `json:"-"`
	PrecommitsValue string `json:"-"`

	HeightRoundStep string `json:"height/round/step"`
	HeightVoteSet   []struct {
		Precommits         []string `json:"precommits"`
		PrecommitsBitArray string   `json:"precommits_bit_array"`
		Prevotes           []string `json:"prevotes"`
		PrevotesBitArray   string   `json:"prevotes_bit_array"`
		Round              string   `json:"round"`
	} `json:"height_vote_set"`
	LockedBlockHash   string `json:"locked_block_hash"`
	ProposalBlockHash string `json:"proposal_block_hash"`
	StartTime         string `json:"start_time"`
	ValidBlockHash    string `json:"valid_block_hash"`
}

func parseTmConsensusState(d json.RawMessage) (res *tmConsensusState, err error) {
	res = &tmConsensusState{}
	defer func() {
		if r := recover(); r != nil {
			log.Printf("parseTmConsensusStat err：%s", r)
			err = errors.New("parse error")
		}
	}()

	err = json.Unmarshal(d, &res)
	if err != nil {
		return nil, err
	}

	hrs := strings.Split(res.HeightRoundStep, "/")
	if len(hrs) == 3 {
		res.Height = hrs[0]
		res.Round = hrs[1]
		res.Step = hrs[2]
	}

	if res.Round == "1" {
		res.PrevotesNum = int64(len(res.HeightVoteSet[1].Prevotes))
		prevotes := strings.Split(res.HeightVoteSet[1].PrevotesBitArray, "=")
		res.PrevotesValue = strings.TrimPrefix(prevotes[1], " ")

		res.PrecommitsNum = int64(len(res.HeightVoteSet[1].Precommits))
		precomms := strings.Split(res.HeightVoteSet[1].PrecommitsBitArray, "=")
		res.PrecommitsValue = strings.TrimPrefix(precomms[1], " ")
	} else {
		res.PrevotesNum = int64(len(res.HeightVoteSet[0].Prevotes))
		prevotes := strings.Split(res.HeightVoteSet[0].PrevotesBitArray, "=")
		res.PrevotesValue = strings.TrimPrefix(prevotes[1], " ")

		res.PrecommitsNum = int64(len(res.HeightVoteSet[0].Precommits))
		precomms := strings.Split(res.HeightVoteSet[0].PrecommitsBitArray, "=")
		res.PrecommitsValue = strings.TrimPrefix(precomms[1], " ")
	}

	return res, nil
}

// UpdateConsensusState 更新共识状态
func UpdateConsensusState(chainID string, cs *tmctypes.ResultConsensusState) error {
	tcs, err := parseTmConsensusState(cs.RoundState)
	if err != nil {
		return err
	}

	mcs, err := retrieveConsensusState(chainID)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		mcs = &model.ConsensusState{
			ChainID:         utils.NullString(chainID),
			Height:          utils.NullString(tcs.Height),
			Round:           utils.NullString(tcs.Round),
			Step:            utils.NullString(tcs.Step),
			PrevotesNum:     utils.NullInt64(tcs.PrevotesNum),
			PrevotesValue:   utils.NullString(tcs.PrevotesValue),
			PrecommitsNum:   utils.NullInt64(tcs.PrecommitsNum),
			PrecommitsValue: utils.NullString(tcs.PrecommitsValue),
			StartTime:       utils.NullString(tcs.StartTime),
		}

		return mcs.Insert(db.Db)
	} else {
		mcs.ChainID = utils.NullString(chainID)
		mcs.Height = utils.NullString(tcs.Height)
		mcs.Round = utils.NullString(tcs.Round)
		mcs.Step = utils.NullString(tcs.Step)
		mcs.PrevotesNum = utils.NullInt64(tcs.PrevotesNum)
		mcs.PrevotesValue = utils.NullString(tcs.PrevotesValue)
		mcs.PrecommitsNum = utils.NullInt64(tcs.PrecommitsNum)
		mcs.PrecommitsValue = utils.NullString(tcs.PrecommitsValue)
		mcs.StartTime = utils.NullString(tcs.StartTime)
		return mcs.Update(db.Db)
	}
}
