package qstarscli

import (
	"context"
	"fmt"

	"github.com/QOSGroup/qstars/client/lcd"
)

func init() {
}

const versionURI = "version"

type versionService service

func (s *versionService) Retrieve(ctx context.Context) (*lcd.ResultCLIVersion, error) {
	u := versionURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res lcd.ResultCLIVersion
	resp, err := s.client.Do(ctx, req, &res)
	fmt.Printf("--resp:%+v,err:%+v\n", resp, err)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
