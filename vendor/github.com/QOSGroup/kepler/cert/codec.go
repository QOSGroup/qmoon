package cert

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

var Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(Codec)
	RegisterCodec(Codec)
}

func MustMarshalJson(obj interface{}) []byte {
	bz, err := Codec.MarshalJSON(obj)
	if err != nil {
		panic(err)
	}
	return bz
}

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Subject)(nil), nil)
	cdc.RegisterConcrete(CommonSubject{},
		CommonSubjAminoRoute, nil)
	cdc.RegisterConcrete(QSCSubject{},
		QSCSubjAminoRoute, nil)
	cdc.RegisterConcrete(QCPSubject{},
		QCPSubjAminoRoute, nil)
}

func MakeCodec() *amino.Codec {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	RegisterCodec(cdc)
	return cdc
}
