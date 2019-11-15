package qstarscli

//
//
//import (
//	"context"
//
//	qosacc "github.com/QOSGroup/qos/account"
//)
//
//const accountsURI = "accounts"
//
//type accountsService service
//
//func (s *accountsService) Create(ctx context.Context) (*qosacc.QOSAccount, error) {
//	u := accountsURI
//	u, err := addOptions(u, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req, err := s.client.NewRequest("POST", u, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var res qosacc.QOSAccount
//	_, err = s.client.Do(ctx, req, &res)
//	if err != nil {
//		return nil, err
//	}
//
//	return &res, nil
//}
//
//func (s *accountsService) Retrieve(ctx context.Context, addr string) (*qosacc.QOSAccount, error) {
//	u := accountsURI + "/" + addr
//	u, err := addOptions(u, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req, err := s.client.NewRequest("GET", u, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var res qosacc.QOSAccount
//	_, err = s.client.Do(ctx, req, &res)
//	if err != nil {
//		return nil, err
//	}
//
//	return &res, nil
//}
