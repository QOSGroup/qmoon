// Copyright 2018 The QOS Authors

package service

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func convertToConsensusState(mcs *models.ConsensusState) *types.ResultConsensusState {
	return &types.ResultConsensusState{
		ChainID:         mcs.ChainId,
		Height:          mcs.Height,
		Round:           mcs.Round,
		Step:            mcs.Step,
		PrevotesNum:     mcs.PrecommitsNum,
		PrevotesValue:   mcs.PrevotesValue,
		PrecommitsNum:   mcs.PrecommitsNum,
		PrecommitsValue: mcs.PrecommitsValue,
		StartTime:       mcs.StartTime,
	}
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

func (n Node) ConsensusState() (*types.ResultConsensusState, error) {
	mcs, err := models.RetrieveConsensusState(n.ChainID)
	if err != nil {
		return nil, err
	}

	return convertToConsensusState(mcs), nil
}

// UpdateConsensusState 更新共识状态
func (n Node) UpdateConsensusState(cs *tmctypes.ResultConsensusState) error {
	tcs, err := parseTmConsensusState(cs.RoundState)
	if err != nil {
		return err
	}

	mcs, err := models.RetrieveConsensusStateByHeight(n.ChainID, tcs.Height)
	//mcs, err := models.re
	if err != nil {
		mcs = &models.ConsensusState{
			Height:          tcs.Height,
			Round:           tcs.Round,
			Step:            tcs.Step,
			PrevotesNum:     tcs.PrevotesNum,
			PrevotesValue:   tcs.PrevotesValue,
			PrecommitsNum:   tcs.PrecommitsNum,
			PrecommitsValue: tcs.PrecommitsValue,
			StartTime:       tcs.StartTime,
		}

		return mcs.Insert(n.ChainID)
	} else {
		mcs.Height = tcs.Height
		mcs.Round = tcs.Round
		mcs.Step = tcs.Step
		mcs.PrevotesNum = tcs.PrevotesNum
		mcs.PrevotesValue = tcs.PrevotesValue
		mcs.PrecommitsNum = tcs.PrecommitsNum
		mcs.PrecommitsValue = tcs.PrecommitsValue
		mcs.StartTime = tcs.StartTime

		return mcs.Update(n.ChainID)
	}
}
