package tmcli

import (
	"context"

	"github.com/QOSGroup/qstars/client/lcd"
)

func init() {
}

const versionURI = "version"

type VersionService service

func (s *VersionService) Retrieve(ctx context.Context) (*lcd.ResultCLIVersion, error) {
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
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
