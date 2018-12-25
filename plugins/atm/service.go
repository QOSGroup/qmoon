// Copyright 2018 The QOS Authors

package atm

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/lib/qstarscli"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/QOSGroup/qstars/x/bank"
	"github.com/sirupsen/logrus"
)

func getBank() string {
	return os.Getenv("ATM_KEY")
}

func getChain() string {
	return os.Getenv("ATM_CHAIN_ID")
}

func getAmount() int64 {
	d := os.Getenv("ATM_AMOUNT")
	if d != "" {
		amount, err := strconv.ParseInt(d, 10, 64)
		if err == nil {
			return amount
		}
	}

	return 2000 * 10000
}

func check(addr, chainid string) error {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	_, err := AtmRecordByAddressChainidCreateat(db.Db, utils.NullString(addr), utils.NullString(chainid), utils.NullTime(t))
	if err == nil {
		return errors.New("今天已经领取")
	}

	return nil
}

func Withdraw(addr, chainid string) (*bank.SendResult, error) {
	prikey := getBank()
	if prikey == "" {
		logrus.WithField("prikey", prikey).Warnf("invalid prikey")
		return nil, nil
	}

	if err := check(addr, chainid); err != nil {
		return nil, err
	}

	opt, err := qstarscli.NewOption(qstarscli.SetOptionHost(os.Getenv("Qstars")))
	if err != nil {
		return nil, err
	}

	coin := "QOS"
	amount := getAmount()
	qcli := qstarscli.NewClient(opt)
	sr, err := qcli.TransferService.Send(nil, &qstarscli.TransferBody{
		Address:    addr,
		Amount:     fmt.Sprintf("%d%s", amount, coin),
		PirvateKey: getBank(),
		ChainID:    chainid,
	})

	if err != nil {
		return nil, err
	}

	if sr.Code != "0" {
		return nil, errors.New(sr.Error)
	}

	ar := &AtmRecord{}
	ar.Address = utils.NullString(addr)
	ar.Chainid = utils.NullString(chainid)
	ar.Coin = utils.NullString(coin)
	ar.Amount = utils.NullString(fmt.Sprintf("%d", amount))
	ar.Createat = utils.NullTime(time.Now())
	ar.Height = utils.NullString(sr.Heigth)
	ar.Hash = utils.NullString(sr.Hash)

	return sr, ar.Insert(db.Db)
}
