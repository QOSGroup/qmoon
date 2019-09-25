// Copyright 2018 The QOS Authors

package lib

import (
	"encoding/json"
	"github.com/QOSGroup/qmoon/lib/cosmos/staking/types"
	"net/http"
	"strings"
)

type Tx struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ResultCosmosTx struct {
	Status int    `json:"status"`
	Txs    []Tx   `json:"txs"`
	Fee    string `json:"fee"`
	IsOK   bool   `json:"isOk"`
	Err    string `json:"err"`
}

type CosmosCli struct {
	remote string
}

func NewCosmosCli(remote string) CosmosCli {
	if remote == "" {
		remote = "http://localhost:19527"
	}
	return CosmosCli{remote: remote}
}

func (cc CosmosCli) Txs(txs []string) ([]ResultCosmosTx, error) {
	body := struct {
		Txs []string `json:"txs"`
	}{
		Txs: txs,
	}

	d, _ := json.Marshal(body)
	resp, err := http.Post(cc.remote+"/txs", "application/json", strings.NewReader(string(d)))
	if err != nil {
		return nil, err
	}

	var result []ResultCosmosTx
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cc CosmosCli) Validators() ([]types.Validator, error) {
	resp, err := http.Get(cc.remote+"/stake/validators")
	if err != nil {
		return nil, err
	}

	var result []types.Validator
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
