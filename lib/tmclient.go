// Copyright 2018 The QOS Authors

package lib

import (
	"errors"
	"sync"

	"github.com/QOSGroup/qbase/store"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	qostypes "github.com/QOSGroup/qos/types"
	"github.com/tendermint/tendermint/rpc/client"
	tmtypes "github.com/tendermint/tendermint/types"
)

type TmClient struct {
	*client.HTTP
	chainID string
}

func (tc *TmClient) ChainID() (string, error) {
	if tc.chainID == "" {
		g, err := tc.Genesis()
		if err != nil {
			return "", err
		}

		tc.chainID = g.Genesis.ChainID
	}

	return tc.chainID, nil
}

func (tc *TmClient) RetrieveTxResult(tx []byte) types.TxStatus {
	txHash := tmtypes.Tx(tx).Hash()

	res, err := tc.Tx(txHash, true)
	if err != nil {
		return types.TxStatusInit
	}

	if res.TxResult.IsOK() {
		return types.TxStatusSuccess
	} else {
		return types.TxStatusFaild
	}
}

func convertQOSValidator(chainID string, val qostypes.Validator) types.Validator {
	return types.Validator{
		Name:           val.Name,
		Owner:          val.Owner.String(),
		ChainID:        chainID,
		Address:        val.ValidatorPubKey.Address().String(),
		PubKeyValue:    string(val.ValidatorPubKey.Bytes()),
		VotingPower:    int64(val.BondTokens),
		Status:         val.Status,
		InactiveCode:   val.InactiveCode,
		InactiveTime:   val.InactiveTime,
		InactiveHeight: int64(val.InactiveHeight),
		BondHeight:     int64(val.BondHeight),
	}
}

func convertQSCValidator(chainID string, val tmtypes.Validator) types.Validator {
	return types.Validator{
		ChainID:     chainID,
		Address:     val.Address.String(),
		PubKeyValue: val.PubKey.Address().String(),
		VotingPower: val.VotingPower,
		Accum:       val.Accum,
	}
}

var (
	//keys see docs/spec/staking.md
	validatorKey            = []byte{0x01} // 保存Validator信息. key: ValidatorAddress
	validatorByOwnerKey     = []byte{0x02} // 保存Owner与Validator的映射关系. key: OwnerAddress, value : ValidatorAddress
	validatorByInactiveKey  = []byte{0x03} // 保存处于`inactive`状态的Validator. key: ValidatorInactiveTime + ValidatorAddress
	validatorByVotePowerKey = []byte{0x04} // 按VotePower排序的Validator地址,不包含`pending`状态的Validator. key: VotePower + ValidatorAddress

	currentValidatorAddressKey = []byte("currentValidatorAddressKey")
)

// QOSValidator 同步最新的validator状态
func (tc *TmClient) QOSValidator(height int64) ([]types.Validator, error) {
	var result []types.Validator
	if height <= 0 {
		height = 0
	}

	chainID, err := tc.ChainID()
	if err != nil {
		return nil, err
	}

	opt := client.ABCIQueryOptions{
		Height:  height,
		Trusted: true,
	}
	path := "/store/validator/subspace"
	resp, err := tc.ABCIQueryWithOptions(path, validatorKey, opt)
	if err != nil {
		return nil, err
	}

	valueBz := resp.Response.GetValue()
	if len(valueBz) == 0 {
		return nil, errors.New("")
	}

	var vKVPair []store.KVPair
	if err := Cdc.UnmarshalBinary(valueBz, &vKVPair); err != nil {
		return nil, err
	}
	for _, kv := range vKVPair {
		var val qostypes.Validator
		if err := Cdc.UnmarshalBinaryBare(kv.Value, &val); err == nil {
			result = append(result, convertQOSValidator(chainID, val))
		}
	}

	return result, nil
}

func convertBlockHeader(header tmtypes.Header) types.BlockHeader {
	return types.BlockHeader{
		ChainID:         header.ChainID,
		Height:          header.Height,
		Time:            header.Time,
		NumTxs:          header.NumTxs,
		TotalTxs:        header.TotalTxs,
		LastCommitHash:  header.LastCommitHash.String(),
		DataHash:        header.DataHash.String(),
		ValidatorsHash:  header.ValidatorsHash.String(),
		ConsensusHash:   header.ConsensusHash.String(),
		AppHash:         header.AppHash.String(),
		LastResultsHash: header.LastResultsHash.String(),
		EvidenceHash:    header.EvidenceHash.String(),
	}
}

func (tc *TmClient) RetrieveBlock(height *int64) (*types.Block, error) {
	var result types.Block

	block, err := tc.Block(height)
	if err != nil {
		return nil, err
	}

	result.Header = convertBlockHeader(block.BlockMeta.Header)
	var txs []types.Tx
	for _, v := range block.Block.Txs {

		txs = append(txs, []byte(v))
	}
	result.Txs = txs

	var precommit []*types.BlockValidator
	for _, v := range block.Block.LastCommit.Precommits {
		if v == nil {
			continue
		}
		precommit = append(precommit, parseVote(block.Block.ChainID, v))
	}
	result.Precommits = precommit

	return &result, nil
}

func parseVote(chainID string, v *tmtypes.Vote) *types.BlockValidator {
	return &types.BlockValidator{
		ChainID:          chainID,
		Height:           v.Height,
		ValidatorAddress: v.ValidatorAddress.String(),
		ValidatorIndex:   int64(v.ValidatorIndex),
		Round:            int64(v.Round),
		Type:             int64(v.Type),
		Signature:        utils.Base64En(v.Signature),
		Timestamp:        v.Timestamp,
	}
}

// Validator 同步最新的validator状态
func (tc *TmClient) Validator(height int64) ([]types.Validator, error) {
	var result []types.Validator
	if height <= 0 {
		height = 0
	}

	chainID, err := tc.ChainID()
	if err != nil {
		return nil, err
	}

	vals, err := tc.Validators(&height)
	if err != nil {
		return nil, err
	}

	for _, v := range vals.Validators {
		result = append(result, convertQSCValidator(chainID, *v))
	}

	return result, nil
}

type tmClients struct {
	tmcs map[string]*TmClient
	lock *sync.Mutex
}

var tmcs *tmClients

func init() {
	tmcs = &tmClients{
		tmcs: make(map[string]*TmClient),
		lock: new(sync.Mutex),
	}
}

func (tc *tmClients) Get(remote string) *TmClient {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	c, ok := tc.tmcs[remote]
	if ok {
		return c
	}

	tmc := &TmClient{
		HTTP: client.NewHTTP(remote, "/websocket"),
	}

	tc.tmcs[remote] = tmc

	return tmc
}

func TendermintClient(remote string) *TmClient {
	return tmcs.Get(remote)
}
