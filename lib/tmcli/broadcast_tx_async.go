package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const broadcastTxAsyncURI = "broadcast_tx_async"

type BroadcastTxAsyncService service

type BroadcastTxAsyncOption struct {
	Tx string `url:"tx,omitempty"`
}

func (s *BroadcastTxAsyncService) Retrieve(ctx context.Context, opt *BroadcastTxAsyncOption) (*tmctypes.ResultBroadcastTx, error) {
	u := broadcastTxAsyncURI
	if opt == nil {
		opt = &BroadcastTxAsyncOption{}
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultBroadcastTx
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
