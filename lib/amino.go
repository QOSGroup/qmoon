// Copyright 2018 The QOS Authors

package lib

import (
	"github.com/QOSGroup/qos/app"

	"github.com/tendermint/go-amino"
)

var Cdc = MakeCodec()
var QOSCdc = MakeCodec()

func MakeCodec() *amino.Codec {
	return app.MakeCodec()
}
