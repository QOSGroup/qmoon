package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const blockResultsURI = "block_results"

type BlockResultsService service

type BlockResultsOption struct {
	Height int `url:"height,omitempty"`
}

func (s *BlockResultsService) Retrieve(ctx context.Context, opt *BlockResultsOption) (*tmctypes.ResultBlockResults, error) {
	u := blockResultsURI
	if opt == nil {
		opt = &BlockResultsOption{}
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultBlockResults
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
