package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const broadcastTxCommitURI = "broadcast_tx_commit"

type BroadcastTxCommitService service

type BroadcastTxCommitOption struct {
	Tx string `url:"tx,omitempty"`
}

func (s *BroadcastTxCommitService) Retrieve(ctx context.Context, opt *BroadcastTxCommitOption) (*tmctypes.ResultBroadcastTx, error) {
	u := broadcastTxCommitURI
	if opt == nil {
		opt = &BroadcastTxCommitOption{}
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
