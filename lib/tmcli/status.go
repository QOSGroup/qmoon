package tmcli

import (
	"context"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func init() {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		"tendermint/PubKeyEd25519", nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{},
		"tendermint/PubKeySecp256k1", nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{},
		"tendermint/PrivKeyEd25519", nil)
	cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{},
		"tendermint/PrivKeySecp256k1", nil)
}

const statusURI = "status"

type statusService service

func (s *statusService) Retrieve(ctx context.Context) (*tmctypes.ResultStatus, error) {
	u := statusURI

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var res tmctypes.ResultStatus
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
