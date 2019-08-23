// Copyright 2018 The QOS Authors

package lib

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/QOSGroup/qbase/store"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	qostypes "github.com/QOSGroup/qos/module/stake/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/bech32"
	"github.com/tendermint/tendermint/rpc/client"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tidwall/gjson"
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

func (tc *TmClient) RetrieveTx(txHash []byte) (*types.Tx, error) {
	var result types.Tx
	res, err := tc.Tx(txHash, true)
	if err != nil {
		log.Printf("TmClient err:%v", err)
		return nil, err
	}

	result.Hash = res.Hash.Bytes()
	result.Height = res.Height
	result.Index = res.Index
	result.Code = res.TxResult.Code
	result.GasWanted = res.TxResult.GetGasWanted()
	result.GasUsed = res.TxResult.GetGasUsed()
	result.Log = res.TxResult.Log
	if res.TxResult.IsOK() {
		result.TxStatus = types.TxStatusSuccess
	} else {
		result.TxStatus = types.TxStatusFaild
	}

	return &result, nil
}

func convertQOSValidator(chainID string, val qostypes.Validator) types.Validator {
	var pubKeyType, pubKeyValue string
	d, err := Cdc.MarshalJSON(val.ValidatorPubKey)
	if err == nil {
		j := string(d)
		pubKeyType = gjson.Get(j, "type").String()
		pubKeyValue = gjson.Get(j, "value").String()
	}

	return types.Validator{
		Name:           val.Description.Moniker,
		Logo:           val.Description.Logo,
		Website:        val.Description.Website,
		Owner:          val.Owner.String(),
		ChainID:        chainID,
		Address:        val.ValidatorPubKey.Address().String(),
		PubKeyType:     pubKeyType,
		PubKeyValue:    pubKeyValue,
		VotingPower:    int64(val.BondTokens),
		Status:         val.Status,
		InactiveCode:   val.InactiveCode,
		InactiveTime:   val.InactiveTime,
		InactiveHeight: int64(val.InactiveHeight),
		BondHeight:     int64(val.BondHeight),
	}
}

func convertQSCValidator(chainID string, val tmtypes.Validator) types.Validator {
	var pubKeyType, pubKeyValue string
	d, err := Cdc.MarshalJSON(val.PubKey)
	if err == nil {
		j := string(d)
		pubKeyType = gjson.Get(j, "type").String()
		pubKeyValue = gjson.Get(j, "value").String()
	}

	return types.Validator{
		ChainID:     chainID,
		Address:     val.Address.String(),
		PubKeyType:  pubKeyType,
		PubKeyValue: pubKeyValue,
		VotingPower: val.VotingPower,
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
		Height: height,
		Prove:  true,
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
	if err := Cdc.UnmarshalBinaryLengthPrefixed(valueBz, &vKVPair); err != nil {
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

// QOSValidator 同步最新的validator状态
func (tc *TmClient) QOSValidatorV0_0_4(height int64) ([]types.Validator, error) {
	var result []types.Validator
	if height <= 0 {
		height = 0
	}

	chainID, err := tc.ChainID()
	if err != nil {
		return nil, err
	}

	opt := client.ABCIQueryOptions{
		Height: height,
		Prove:  true,
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
	if err := Cdc.UnmarshalBinaryBare(valueBz, &vKVPair); err != nil {
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
		ProposerAddress: header.ProposerAddress.String(),
	}
}

func (tc *TmClient) RetrieveBlock(height *int64) (*types.Block, error) {
	var result types.Block

	block, err := tc.Block(height)
	if err != nil {
		return nil, err
	}

	result.Header = convertBlockHeader(block.BlockMeta.Header)
	var txs [][]byte
	for _, v := range block.Block.Txs {
		txs = append(txs, v)
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

	if block.Block.Evidence.Evidence != nil {
		result.EvidenceList = types.EvidenceList{
			Time:      block.Block.Time,
			Evidences: parseEvidence(block.Block.Evidence.Evidence),
		}
	}

	return &result, nil
}

func parseEvidence(es tmtypes.EvidenceList) []types.Evidence {
	var res []types.Evidence
	for _, v := range es {
		res = append(res, types.Evidence{
			Height:  v.Height(),
			Address: string(v.Address()),
			Hash:    string(v.Hash()),
			Data:    v.String(),
		})
	}

	return res
}

func parseVote(chainID string, v *tmtypes.CommitSig) *types.BlockValidator {
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

// ValidatorLoop 同步最新的validator状态
func (tc *TmClient) Validator(height int64) ([]types.Validator, error) {
	var result []types.Validator
	if height <= 0 {
		height = 0
	}

	vals, err := tc.Validators(&height)
	if err != nil {
		return nil, err
	}

	for _, v := range vals.Validators {
		result = append(result, convertQSCValidator(tc.chainID, *v))
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

func Bech32AddressToHex(addr string) string {
	_, edPub, err := bech32.DecodeAndConvert(addr)
	if err != nil {
		return addr
	}
	var pubKey2 ed25519.PubKeyEd25519
	Cdc.UnmarshalBinaryBare(edPub, &pubKey2)
	return pubKey2.Address().String()
}

func PubkeyToBech32Address(hrp string, t, val string) string {
	caHex := fmt.Sprintf(`{"type": "%s","value": "%s"}`, t, val)
	var pubkey ed25519.PubKeyEd25519
	err := Cdc.UnmarshalJSON([]byte(caHex), &pubkey)
	if err != nil {
		return ""
	}

	bech32Pub, _ := bech32.ConvertAndEncode(hrp, pubkey.Bytes())
	return bech32Pub
}
