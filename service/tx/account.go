// Copyright 2018 The QOS Authors

package tx

import (
	"github.com/QOSGroup/qbase/account"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/types"
)

// QueryAccount query account by addr
func (c Client) QueryAccount(addr string) (*types.ResultAccount, error) {
	address, err := qbasetypes.GetAddrFromBech32(addr)
	if err != nil {
		return nil, err
	}
	key := account.AddressStoreKey(address)

	acc, err := c.c.GetAccount(key, c.cdc)
	if err != nil {
		return nil, err
	}

	result := &types.ResultAccount{}
	result.Address = acc.AccountAddress.String()
	result.Nonce = acc.Nonce

	result.Coins = append(result.Coins, &qbasetypes.BaseCoin{Name: "qos", Amount: acc.QOS})
	for _, v := range acc.QSCs {
		result.Coins = append(result.Coins, v)
	}

	return result, nil
}
