// Copyright 2018 The QOS Authors

package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func init() {

}

const netInfoURI = "net_info"

type netInfoService service

func (s *netInfoService) Retrieve(ctx context.Context) (*tmctypes.ResultNetInfo, error) {
	u := netInfoURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultNetInfo
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
