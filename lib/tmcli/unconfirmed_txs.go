package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const unconfirmedUnconfirmedTxssURI = "unconfirmed_txs"

type unconfirmedTxsService service

type UnconfirmedTxsOption struct {
	Limit int `url:"limit,omitempty"`
}

func (s *unconfirmedTxsService) Retrieve(cunconfirmedUnconfirmedTxss context.Context, opt *UnconfirmedTxsOption) (*tmctypes.ResultUnconfirmedTxs, error) {
	u := unconfirmedUnconfirmedTxssURI

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultUnconfirmedTxs
	_, err = s.client.Do(cunconfirmedUnconfirmedTxss, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
