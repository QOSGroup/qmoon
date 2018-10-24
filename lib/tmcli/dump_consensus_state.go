// Copyright 2018 The QOS Authors

package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func init() {

}

const dumpDumpConsensusStateURI = "dump_consensus_state"

type DumpConsensusStateService service

func (s *DumpConsensusStateService) Retrieve(ctx context.Context) (*tmctypes.ResultDumpConsensusState, error) {
	u := dumpDumpConsensusStateURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultDumpConsensusState
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
