package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const abciQueryURI = "abci_query"

type abciQueryService service

type AbciQueryOption struct {
	Path    string `url:"path,omitempty"`
	Data    []byte `url:"data,omitempty"`
	Height  int64  `url:"height,omitempty"`
	Trusted bool   `url:"trusted,omitempty"`
}

func (s *abciQueryService) Retrieve(ctx context.Context, path string, data []byte, height int64, trusted bool) (*tmctypes.ResultABCIQuery, error) {
	u := abciQueryURI
	opt := &AbciQueryOption{
		Path:    path,
		Data:    data,
		Height:  height,
		Trusted: trusted,
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultABCIQuery
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
