package qstarscli

import (
	"context"

	"github.com/QOSGroup/qstars/x/kvstore"
)

const kvURI = "kv"

type KVService service

type CreateKVQuery struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	PrivateKey string `json:"privatekey"`
	ChainID    string `json:"chainid"`
}

func (s *KVService) Create(ctx context.Context, key string, value string, privatekey string, chainid string) (
	*kvstore.ResultSendKV, error) {
	u := kvURI
	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}
	b := CreateKVQuery{
		Key:        key,
		Value:      value,
		PrivateKey: privatekey,
		ChainID:    chainid,
	}
	req, err := s.client.NewRequest("POST", u, b)
	if err != nil {
		return nil, err
	}

	var res kvstore.ResultSendKV
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *KVService) Retrieve(ctx context.Context, key string) (*kvstore.ResultGetKV, error) {
	u := kvURI + "/" + key
	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res kvstore.ResultGetKV
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
