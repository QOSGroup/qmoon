package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const broadcastTxSyncURI = "broadcast_tx_sync"

type broadcastTxSyncService service

type BroadcastTxSyncOption struct {
	Tx string `url:"tx,omitempty"`
}

func (s *broadcastTxSyncService) Retrieve(ctx context.Context, opt *BroadcastTxSyncOption) (*tmctypes.ResultBroadcastTx, error) {
	u := broadcastTxSyncURI
	if opt == nil {
		opt = &BroadcastTxSyncOption{}
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
