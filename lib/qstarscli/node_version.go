package qstarscli

import (
	"context"

	"github.com/QOSGroup/qstars/client/lcd"
)

func init() {
}

const nodeVersionURI = "node_version"

type NodeVersionService service

func (s *NodeVersionService) Retrieve(ctx context.Context) (*lcd.ResultNodeVersion, error) {
	u := nodeVersionURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res lcd.ResultNodeVersion
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
