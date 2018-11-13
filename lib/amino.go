// Copyright 2018 The QOS Authors

package lib

import (
	"github.com/QOSGroup/qstars/star"
	"github.com/tendermint/go-amino"
)

var Cdc = MakeCodec()

func MakeCodec() *amino.Codec {
	cdc := star.MakeCodec()
	return cdc
}
