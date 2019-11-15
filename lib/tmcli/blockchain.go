package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const blockchainURI = "blockchain"

type blockchainService service

type BlockchainOption struct {
	MinHeight int64 `url:"minHeight,omitempty"`
	MaxHeight int64 `url:"maxHeight,omitempty"`
}

func (s *blockchainService) List(ctx context.Context, opt *BlockchainOption) (*tmctypes.ResultBlockchainInfo, error) {
	u := blockchainURI
	if opt == nil {
		opt = &BlockchainOption{}
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultBlockchainInfo
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
