// Copyright 2018 The QOS Authors

package types

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	amino "github.com/tendermint/go-amino"
)

const (
	UrlPrefixNodes   = "/nodes"
	UrlNodeProxy     = UrlPrefixNodes + "/:nodeName"
	UrlPrefixPlugins = "/plugins"
)

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func NewRPCErrorResponse(id string, code int, msg string, data string) RPCResponse {
	return RPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error:   &RPCError{Code: code, Message: msg, Data: data},
	}
}

func NewRPCSuccessResponse(cdc *amino.Codec, id string, res interface{}) RPCResponse {
	var rawMsg json.RawMessage

	if res != nil {
		var js []byte
		js, err := cdc.MarshalJSON(res)
		if err != nil {
			return RPCInternalError(id, errors.Wrap(err, "Error marshalling response"))
		}
		rawMsg = json.RawMessage(js)
	}

	return RPCResponse{JSONRPC: "2.0", ID: id, Result: rawMsg}
}

func RPCParseError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, -32700, "Parse error. Invalid JSON", err.Error())
}

func RPCInvalidRequestError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, -32600, "Invalid Request", err.Error())
}

func RPCMethodNotFoundError(id string) RPCResponse {
	return NewRPCErrorResponse(id, -32601, "Method not found", "")
}

func RPCInvalidParamsError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, -32602, "Invalid params", err.Error())
}

func RPCUnauthorizedError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, http.StatusUnauthorized, "Unauthorized", err.Error())
}

func RPCForbiddenError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, http.StatusForbidden, "Forbidden", err.Error())
}

func RPCInternalError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, -32603, err.Error(), "Internal error")
}

func RPCServerError(id string, err error) RPCResponse {
	return NewRPCErrorResponse(id, -32000, "Server error", err.Error())
}
