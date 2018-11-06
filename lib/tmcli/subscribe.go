package tmcli

import (
	"context"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const subscribeURI = "subscribe"

type subscribeService service

type SubscribeOption struct {
	Query string `url:"query,omitempty"`
}

func (s *subscribeService) Retrieve(ctx context.Context, query string) (*tmctypes.ResultSubscribe, error) {
	u := subscribeURI
	opt := &SubscribeOption{Query: query}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultSubscribe
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
