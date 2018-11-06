package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const unsubscribeURI = "unsubscribe"

type unsubscribeService service

type UnsubscribeOption struct {
	Query string `url:"query,omitempty"`
}

func (s *unsubscribeService) Retrieve(ctx context.Context, query string) (*tmctypes.ResultUnsubscribe, error) {
	u := unsubscribeURI
	opt := &UnsubscribeOption{Query: query}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultUnsubscribe
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
