// Copyright 2018 The QOS Authors

package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func init() {

}

const genesisURI = "genesis"

type GenesisService service

func (s *GenesisService) Retrieve(ctx context.Context) (*tmctypes.ResultGenesis, error) {
	u := genesisURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultGenesis
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
