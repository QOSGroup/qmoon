// Copyright 2018 The QOS Authors

package service

import (
	"testing"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib/tmcli"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateBlock(t *testing.T) {
	b, err := tmcli.NewClient(nil).Block.Retrieve(nil, nil)
	assert.Nil(t, err)
	err = CreateBlock(b)
	assert.Nil(t, err)

	mb, err := model.BlockByChainIDHeight(db.Db, utils.NullString(b.Block.ChainID), utils.NullInt64(b.Block.Height))
	assert.Nil(t, err)
	assert.Equal(t, b.Block.NumTxs, mb.NumTxs.Int64)
	assert.Equal(t, b.Block.DataHash.String(), mb.DataHash.String)
	assert.Equal(t, b.Block.TotalTxs, mb.TotalTxs.Int64)

	mt, err := model.TxesByChainIDHeight(db.Db, utils.NullString(b.Block.ChainID), utils.NullInt64(b.Block.Height))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(mt))
	assert.Equal(t, "TransferTx", mt[0].TxType.String)
}

func TestSyncLock(t *testing.T) {
	key := "abc"
	var ok bool

	ok = SyncLock(key)
	assert.True(t, ok)

	ok = SyncLock(key)
	assert.False(t, ok)

	ok = SyncUnlock(key)
	assert.True(t, ok)

	ok = SyncLock(key)
	assert.True(t, ok)

	ok = SyncUnlock(key)
	assert.True(t, ok)

	ok = SyncUnlock(key)
	assert.False(t, ok)

	ok = SyncLock(key)
	assert.True(t, ok)
}
