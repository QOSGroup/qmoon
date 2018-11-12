// Copyright 2018 The QOS Authors

package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/service/block"
	"github.com/QOSGroup/qmoon/service/tx"
	"github.com/QOSGroup/qmoon/utils"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// CreateBlock 创建一个块
func CreateBlock(b *tmctypes.ResultBlock) error {
	var err error
	err = block.Save(b)
	if err != nil {
		return err
	}

	err = tx.Save(b)
	if err != nil {
		return err
	}

	return nil
}

// SyncLock 同步时锁定，同一个时间只会有一个同步协程
func SyncLock(key string) bool {
	key = "lock_" + key
	value := "1"

	qs, err := model.QmoonStatusByKey(db.Db, utils.NullString(key))
	if err != nil {
		if err == sql.ErrNoRows {
			qs = &model.QmoonStatus{
				Key:   utils.NullString(key),
				Value: utils.NullString(value),
			}
			err := qs.Insert(db.Db)
			return err == nil
		}
	}

	if qs.Value.String == "1" {
		return false
	}

	s := fmt.Sprintf("update public.qmoon_status set value='%s' where key='%s' and value='0'", value, key)
	log.Printf(s)
	r, err := db.Db.Exec(s)
	if err != nil {
		return false
	}

	num, err := r.RowsAffected()
	if err != nil {
		return false
	}

	return num == 1
}

func SyncUnlock(key string) bool {
	key = "lock_" + key

	value := "0"
	s := fmt.Sprintf("update public.qmoon_status set value='%s' where key='%s' and value='1'", value, key)
	log.Printf(s)
	r, err := db.Db.Exec(s)
	if err != nil {
		return false
	}

	num, err := r.RowsAffected()
	if err != nil {
		return false
	}

	return num == 1
}
