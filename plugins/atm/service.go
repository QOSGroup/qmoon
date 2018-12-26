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
	t := utils.DayStart(time.Now())
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

	coin := "qos"
	amount := getAmount()
	sr, err := send(addr, chainid, fmt.Sprintf("%d%s", amount, coin))
	if err != nil {
		return nil, err
	}

	ar := &AtmRecord{}
	ar.Address = utils.NullString(addr)
	ar.Chainid = utils.NullString(chainid)
	ar.Coin = utils.NullString(coin)
	ar.Amount = utils.NullString(fmt.Sprintf("%d", amount))
	ar.Createat = utils.NullTime(utils.DayStart(time.Now()))
	ar.Height = utils.NullString(sr.Heigth)
	ar.Hash = utils.NullString(sr.Hash)

	if err := ar.Insert(db.Db); err != nil {
		return nil, err
	}

	return sr, nil
}

func send(addr, chainid string, amount string) (*bank.SendResult, error) {
	opt, err := qstarscli.NewOption(qstarscli.SetOptionHost(os.Getenv("Qstars")))
	if err != nil {
		return nil, err
	}

	qcli := qstarscli.NewClient(opt)
	sr, err := qcli.TransferService.Send(nil, &qstarscli.TransferBody{
		Address:    addr,
		Amount:     amount,
		PirvateKey: getBank(),
		ChainID:    chainid,
	})

	if err != nil {
		return nil, err
	}

	if sr.Hash == "" {
		return nil, errors.New(sr.Error)
	}

	return sr, nil
}
