package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const validatorsURI = "validators"

type validatorsService service

type ValidatorsOption struct {
	Height int `url:"height,omitempty"`
}

func (s *validatorsService) Retrieve(ctx context.Context, opt *ValidatorsOption) (*tmctypes.ResultValidators, error) {
	u := validatorsURI
	if opt == nil {
		opt = &ValidatorsOption{}
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultValidators
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
