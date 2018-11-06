package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const blockURI = "block"

type blockService service

type BlockOption struct {
	Height int `url:"height,omitempty"`
}

func (s *blockService) Retrieve(ctx context.Context, opt *BlockOption) (*tmctypes.ResultBlock, error) {
	u := blockURI
	if opt == nil {
		opt = &BlockOption{}
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultBlock
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
