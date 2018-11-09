// Copyright 2018 The QOS Authors

package tx

import (
	"os"

	"github.com/QOSGroup/qmoon/lib"
	qstarscontext "github.com/QOSGroup/qstars/client/context"
	"github.com/tendermint/go-amino"
)

type Client struct {
	c   *qstarscontext.CLIContext
	cdc *amino.Codec
}

func NewClient(nodeURI string) Client {
	cdc := lib.MakeCodec()
	c := qstarscontext.NewCLIContext1(nodeURI).
		WithCodec(cdc).
		WithLogger(os.Stdout)

	return Client{c: &c, cdc: cdc}
}

func (c Client) MarshalJSONIndent(o interface{}, prefix, indent string) ([]byte, error) {
	return c.cdc.MarshalJSONIndent(o, prefix, indent)
}

func (c Client) UnmarshalBinaryBare(bz []byte, ptr interface{}) error {
	return c.cdc.UnmarshalBinaryBare(bz, ptr)
}

func (c Client) UnmarshalJSON(bz []byte, ptr interface{}) error {
	return c.cdc.UnmarshalJSON(bz, ptr)
}

func (c Client) MarshalJSON(o interface{}) ([]byte, error) {
	return c.cdc.MarshalJSON(o)
}
