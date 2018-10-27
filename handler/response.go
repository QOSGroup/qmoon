// Copyright 2018 The QOS Authors

package handler

import "github.com/QOSGroup/qmoon/types"

func WrapperRPCResponse(id string, result interface{}, err *types.RPCError) types.RPCResponse {
	return types.RPCResponse{
		JSONRPC: "2.0",
		ID:      "",
		Result:  result,
		Error:   err,
	}
}

func NewRPCError(code int, message, data string) *types.RPCError {
	return &types.RPCError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
