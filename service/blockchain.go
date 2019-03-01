// Copyright 2018 The QOS Authors

package service

import (
	"encoding/json"
	"errors"

	qbaseaccount "github.com/QOSGroup/qbase/client/account"

	"github.com/QOSGroup/qbase/client/context"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/tendermint/tendermint/rpc/client"
)

func QueryAccount(node *Node, addr string) (json.RawMessage, error) {
	switch node.NodeType {
	case types.NodeTypeQOS.String():
		return queryQOSAccount(node, addr)
	case types.NodeTypeQSC.String():
		return queryQSCAccount(node, addr)
	case types.NodeTypeCOSMOS.String():
		return queryCOSMOSAccount(node, addr)
	default:
		return nil, errors.New("不支持链类型")
	}
}

func queryQOSAccount(node *Node, addr string) (json.RawMessage, error) {
	cdc := lib.Cdc
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithClient(client.NewHTTP(node.BaseURL, "/websocket"))
	accAddr, err := qbaseaccount.GetAddrFromValue(cliCtx, addr)
	if err != nil {
		return nil, err
	}

	acc, err := qbaseaccount.GetAccount(cliCtx, accAddr)
	if err != nil {
		return nil, err
	}

	res, err := cdc.MarshalJSON(acc)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(res), nil
}

func queryQSCAccount(node *Node, addr string) (json.RawMessage, error) {
	cdc := lib.Cdc
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithClient(client.NewHTTP(node.BaseURL, "/websocket"))
	accAddr, err := qbaseaccount.GetAddrFromValue(cliCtx, addr)
	if err != nil {
		return nil, err
	}

	acc, err := qbaseaccount.GetAccount(cliCtx, accAddr)
	if err != nil {
		return nil, err
	}

	res, err := cdc.MarshalJSONIndent(acc, "", "  ")
	if err != nil {
		return nil, err
	}

	return json.RawMessage(res), nil
}

func queryCOSMOSAccount(node *Node, addr string) (json.RawMessage, error) {
	return nil, nil
}
