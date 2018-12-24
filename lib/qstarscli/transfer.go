package qstarscli

import (
	"context"
	"strings"

	"github.com/QOSGroup/qstars/x/bank"
)

const transferURI = "/accounts/{address}/send"

type transferService service

type TransferBody struct {
	Address       string `json:"-"`
	Amount        string `json:"amount"`
	PirvateKey    string `json:"privatekey"`
	ChainID       string `json:"chain_id"`
	AccountNumber int64  `json:"account_number"`
	Sequence      int64  `json:"sequence"`
	Gas           int64  `json:"gas"`
}

func (s *transferService) Send(ctx context.Context, body *TransferBody) (*bank.SendResult, error) {
	u := strings.Replace(transferURI, "{address}", body.Address, -1)
	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	var res bank.SendResult
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
