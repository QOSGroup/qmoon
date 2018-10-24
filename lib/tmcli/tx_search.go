package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const txSearchURI = "tx_search"

type TxSearchService service

type TxSearchOption struct {
	Prove   bool `url:"prove,omitempty"`
	Page    int  `url:"page,omitempty"`
	PerPage int  `url:"perPage,omitempty"`
}

type txSearchQuery struct {
	Query   string `url:"query,omitempty"`
	Prove   bool   `url:"prove,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"perPage,omitempty"`
}

func (s *TxSearchService) Retrieve(ctx context.Context, query string, opt *TxSearchOption) (*tmctypes.ResultTxSearch, error) {
	u := txSearchURI
	q := &txSearchQuery{
		Query: query,
	}
	if opt != nil {
		q.Prove = opt.Prove
		q.Page = opt.Page
		q.PerPage = opt.PerPage
	}
	u, err := addOptions(u, q)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultTxSearch
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
