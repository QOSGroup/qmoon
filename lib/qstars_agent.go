// Copyright 2018 The QOS Authors

package lib

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ResultQstarsAgentTx struct {
	Status int    `json:"status"`
	Tx     Tx     `json:"txs"`
	IsOK   bool   `json:"isOk"`
	Err    string `json:"err"`
}

type QstarsAgentCli struct {
	remote string
}

func NewQstarsAgentCli(remote string) QstarsAgentCli {
	if remote == "" {
		remote = "http://localhost:19528"
	}
	return QstarsAgentCli{remote: remote}
}

func (cc QstarsAgentCli) Txs(txs []string) ([]ResultQstarsAgentTx, error) {
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

	var result []ResultQstarsAgentTx
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
