// Copyright 2018 The QOS Authors

package lib

import (
	"github.com/QOSGroup/kepler/cert"
	"github.com/QOSGroup/qbase/baseabci"

	qosaccount "github.com/QOSGroup/qos/account"
	qostxs "github.com/QOSGroup/qos/txs"
	qsc "github.com/QOSGroup/qstars/star"

	"github.com/tendermint/go-amino"
)

var Cdc = MakeQSCCodec()

var QOSCdc = MakeQOSCodec()
var QSCCdc = MakeQSCCodec()
var COSMOSCdc = MakeCOSMOSCodec()

func MakeCodec() *amino.Codec {
	cdc := qsc.MakeCodec()
	return cdc
}

func MakeQOSCodec() *amino.Codec {
	cdc := baseabci.MakeQBaseCodec()
	qostxs.RegisterCodec(cdc)
	qosaccount.RegisterCodec(cdc)
	cert.RegisterCodec(cdc)
	return cdc
}

// qsc
func MakeQSCCodec() *amino.Codec {
	return qsc.MakeCodec()
}

// cosmos
func MakeCOSMOSCodec() *amino.Codec {
	var cdc = amino.NewCodec()
	//bank.RegisterCodec(cdc)
	//staking.RegisterCodec(cdc)
	//distribution.RegisterCodec(cdc)
	//slashing.RegisterCodec(cdc)
	//gov.RegisterCodec(cdc)
	//auth.RegisterCodec(cdc)
	//cosmostypes.RegisterCodec(cdc)
	//cryptoAmino.RegisterAmino(cdc)
	return cdc
}
