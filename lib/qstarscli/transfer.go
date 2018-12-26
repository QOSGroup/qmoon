package qstarscli

import (
	"context"
	"strings"

	"github.com/QOSGroup/qstars/x/bank"
	"github.com/sirupsen/logrus"
)

const transferURI = "/accounts/{address}/send"

type transferService service

type TransferBody struct {
	Address       string `json:"-"`
	Amount        string `json:"amount"`
	PirvateKey    string `json:"privatekey"`
	ChainID       string `json:"chain_id"`
	AccountNumber int64  `json:"-"`
	Sequence      int64  `json:"-"`
	Gas           int64  `json:"-"`
}

func (s *transferService) Send(ctx context.Context, body *TransferBody) (*bank.SendResult, error) {
	u := strings.Replace(transferURI, "{address}", body.Address, -1)
	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, body)
	logrus.WithField("module", "qstarscli").WithField("url", s.client.host+u).WithField("body", body).Debug()
	if err != nil {
		return nil, err
	}

	var res bank.SendResult
	resp, err := s.client.Do(ctx, req, &res)
	logrus.WithField("module", "qstarscli").WithField("resp", resp).WithField("err", err).Debugln()
	if err != nil {
		return nil, err
	}

	return &res, nil
}
