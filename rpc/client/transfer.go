package client

import (
	"context"
	"strings"

	"github.com/QOSGroup/qmoon/plugins/transfer"
)

const transferURI = "/accounts/{address}/transfer"

type transferService service

type TransferOption struct {
	Offset int64 `url:"offset,omitempty"`
	Limit  int64 `url:"limit,omitempty"`
}

func (s *transferService) Retrieve(ctx context.Context, addr string, opt *TransferOption) ([]*transfer.TxTransfer, error) {
	u := strings.Replace(transferURI, "{address}", addr, -1)
	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res []*transfer.TxTransfer
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
