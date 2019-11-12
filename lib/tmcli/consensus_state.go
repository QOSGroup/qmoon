// Copyright 2018 The QOS Authors

package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func init() {

}

const consensusStateURI = "consensus_state"

type consensusStateService service

func (s *consensusStateService) Retrieve(ctx context.Context) (*tmctypes.ResultConsensusState, error) {
	u := consensusStateURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultConsensusState
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
