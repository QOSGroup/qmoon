// Copyright 2018 The QOS Authors

package tx

import (
	"github.com/QOSGroup/qbase/account"
	"github.com/QOSGroup/qbase/types"
	qosaccount "github.com/QOSGroup/qos/account"
)

// QueryAccount query account by addr
func (c Client) QueryAccount(addr string) (*qosaccount.QOSAccount, error) {
	address, err := types.GetAddrFromBech32(addr)
	if err != nil {
		return nil, err
	}
	key := account.AddressStoreKey(address)

	acc, err := c.c.GetAccount(key, c.cdc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
