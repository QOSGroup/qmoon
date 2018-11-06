package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const txURI = "tx"

type txService service

type TxOption struct {
	Hash  []byte `url:"hash,omitempty"`
	Prove bool   `url:"prove,omitempty"`
}

func (s *txService) Retrieve(ctx context.Context, hash []byte, prove bool) (*tmctypes.ResultTx, error) {
	u := txURI
	opt := &TxOption{
		Hash:  hash,
		Prove: prove,
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultTx
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
