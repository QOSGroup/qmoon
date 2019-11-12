// Copyright 2018 The QOS Authors

package types

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestAddrPrefix(t *testing.T) {
	t.Run("cosmos", func(t *testing.T) {
		val := "cosmosvalconspub1zcjduepqmqrghhgxs6q8t6h7hlvwygawht78surjg3e47jdjxkkx8pc5n3psylm6gl"
		h := "981560BCEABD99DA0F20464118628854E3C7555D"

		d, _ := hex.DecodeString(h)
		t.Logf("--1: %v", base64.StdEncoding.EncodeToString(d))
		t.Logf("--2: %v", val)

	})
}
