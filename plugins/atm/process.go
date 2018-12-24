// Copyright 2018 The QOS Authors

package atm

import (
	"database/sql"
	qbasetxs "github.com/QOSGroup/qbase/txs"
	"github.com/gin-gonic/gin"
	tmtypes "github.com/tendermint/tendermint/types"
)

type ATMPlugin struct{}

func (ttp ATMPlugin) DbInit(driveName string, db *sql.DB) error {

	return DbInit(driveName, db)
}

func (ttp ATMPlugin) DbClear(driveName string, db *sql.DB) error {
	return DbClear(driveName, db)
}

func (ttp ATMPlugin) Type() string {
	return "ATM"
}

func (ttp ATMPlugin) RegisterGin(r *gin.Engine) {
	AccountWithdrawGinRegister(r)
}

func (ttp ATMPlugin) Parse(blockHeader tmtypes.Header, itx qbasetxs.ITx) (typeName string, hit bool, err error) {
	return "", false, nil
}
