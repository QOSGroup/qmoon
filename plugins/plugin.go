// Copyright 2018 The QOS Authors

package plugins

import (
	"database/sql"
	"errors"
	"fmt"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	"github.com/QOSGroup/qmoon/plugins/atm"
	"github.com/QOSGroup/qmoon/plugins/transfer"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Pluginer interface {
	DbInit(driveName string, db *sql.DB) error
	DbClear(driveName string, db *sql.DB) error

	Parse(blockHeader types.BlockHeader, itx qbasetxs.ITx) (typeName string, hit bool, err error)
	Type() string

	Doctor() error

	RegisterGin(r *gin.Engine)
}

var tps = make(map[string]Pluginer)

func init() {
	register(&transfer.TxTransferPlugin{})
	register(&atm.ATMPlugin{})
}

func RegisterGin(r *gin.Engine) {
	for _, v := range tps {
		v.RegisterGin(r)
	}
}

func register(tp Pluginer) error {
	if _, ok := tps[tp.Type()]; ok {
		return errors.New("aleady registered")
	} else {
		tps[tp.Type()] = tp
	}

	return nil
}

func DbUp(driveName string, db *sql.DB) error {
	logrus.WithField("module", "plugins").Debug(tps)
	for _, tp := range tps {
		if err := tp.DbInit(driveName, db); err != nil {
			logrus.WithField("module", "plugins").WithField("plugin", tp.Type()).Warn(err)
			return err
		}
		logrus.WithField("module", "plugins").WithField("plugin", tp.Type()).Warn("OK")
	}

	return nil
}

func DbDown(driveName string, db *sql.DB) error {
	logrus.WithField("module", "plugins").Debug(tps)
	for _, tp := range tps {
		if err := tp.DbClear(driveName, db); err != nil {
			logrus.WithField("module", "plugins").WithField("plugin", tp.Type()).Warn(err)
			return err
		}
		logrus.WithField("module", "plugins").WithField("plugin", tp.Type()).Warn("OK")
	}

	return nil
}

func Parse(blockHeader types.BlockHeader, itx qbasetxs.ITx) (name string, err error) {
	var hit bool
	for _, tp := range tps {
		if name, hit, err = tp.Parse(blockHeader, itx); hit {
			break
		}
	}

	if !hit {
		return "", errors.New("unsupport tx")
	}

	return
}

func Doctor() error {
	for _, tp := range tps {
		if err := tp.Doctor(); err != nil {
			return fmt.Errorf("Check %s fail, err:%s ", tp.Type(), err.Error())
		}
	}

	return nil
}
