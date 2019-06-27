// Copyright 2018 The QOS Authors

package atm

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	qbaseaccount "github.com/QOSGroup/qbase/client/account"
	"github.com/QOSGroup/qbase/client/context"
	"github.com/QOSGroup/qbase/txs"
	"github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/utils"
	txtransfer "github.com/QOSGroup/qos/module/transfer"
	transtypes "github.com/QOSGroup/qos/module/transfer/types"
	qostypes "github.com/QOSGroup/qos/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const AtmIPLimit = 5

var atmPrikey crypto.PrivKey

func getAtmIPLimit() int64 {
	limit := os.Getenv("ATM_IPLimit")

	if d, err := strconv.ParseInt(limit, 10, 64); err == nil {
		return d
	}

	return AtmIPLimit
}

func getBank() (crypto.PrivKey, error) {
	if atmPrikey != nil {
		return atmPrikey, nil
	}
	var pri ed25519.PrivKeyEd25519
	prikStr := strings.TrimSpace(os.Getenv("ATM_KEY"))
	if prikStr == "" {
		return pri, errors.New("atm ed25519 private key is empty")
	}
	privBytes, err := base64.StdEncoding.DecodeString(prikStr)
	if err != nil {
		return pri, err
	}
	copy(pri[:], privBytes)

	atmPrikey = pri

	return pri, nil
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

func ipCheck(ip, chainid string) error {
	t := utils.DayStart(time.Now())
	ir, err := models.RetrieveAtmIpRecordByIpCreateat(chainid, ip, t)
	if err != nil {
		return nil
	}

	if ir.Amount >= int(getAtmIPLimit()) {
		return errors.New("超过今日领取上限，请明天再来")
	}

	return nil
}

func ipWithdraw(ip, chainid string) error {
	t := utils.DayStart(time.Now())
	ir, err := models.RetrieveAtmIpRecordByIpCreateat(chainid, ip, t)
	if err != nil {
		ir = &models.AtmIpRecord{
			Ip:            ip,
			Amount:        1,
			CreatedAtUnix: t.Unix(),
		}
		return ir.Insert(chainid)
	} else {
		ir.Amount = ir.Amount + 1
		return ir.Update(chainid)
	}

	return nil
}

func check(addr, chainid string) error {
	t := utils.DayStart(time.Now())
	_, err := models.RetrieveAtmRecordByAddressCreateat(chainid, addr, t)
	if err == nil {
		return errors.New("今天已经领取")
	}

	return nil
}

func Withdraw(addr, chainid, nodeUrl string) (*ctypes.ResultBroadcastTxCommit, error) {
	pri, err := getBank()
	if err != nil {
		logrus.WithField("pri", pri).Warnf("invalid pri")
		return nil, errors.New("服务异常")
	}
	acc := types.Address(pri.PubKey().Address().Bytes())

	if err := check(addr, chainid); err != nil {
		return nil, err
	}

	maxGas := int64(2000000)
	// maxGas := viper.GetInt64("max-gas")
	// if maxGas < 0 {
	// 	return nil, errors.New("max-gas flag not correct")
	// }

	coin := "qos"
	amount := getAmount()
	res, err := send(nodeUrl, pri, maxGas, fmt.Sprintf("%s,%d%s", acc.String(), amount, coin),
		fmt.Sprintf("%s,%d%s", addr, amount, coin))
	logrus.Debugf("send:res:%+v,err:%v", res, err)
	if err != nil {
		return nil, err
	}

	ar := &models.AtmRecord{}
	ar.Address = addr
	ar.Coin = coin
	ar.Amount = fmt.Sprintf("%d", amount)
	ar.Createat = utils.DayStart(time.Now())
	ar.Height = res.Height
	ar.Hash = res.Hash.String()

	if err := ar.Insert(chainid); err != nil {
		return nil, err
	}

	return res, nil
}

func send(remote string, pri crypto.PrivKey, maxGas int64, senderStr, receiverStr string) (
	*ctypes.ResultBroadcastTxCommit, error) {
	cdc := lib.Cdc
	tmcli := client.NewHTTP(remote, "/websocket")
	genesis, err := tmcli.Genesis()
	if err != nil {
		return nil, err
	}

	chainid := genesis.Genesis.ChainID
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithClient(client.NewHTTP(remote, "/websocket"))

	tx, err := transferBuilder(cliCtx, senderStr, receiverStr)
	if err != nil {
		return nil, err
	}
	txStd := txs.NewTxStd(tx, chainid, types.NewInt(maxGas))

	nonce, err := getDefaultAccountNonce(cliCtx, tx.Senders[0].Address.Bytes())
	if err != nil || nonce < 0 {
		return nil, err
	}
	nonce = nonce + 1

	sig, err := pri.Sign(txStd.BuildSignatureBytes(nonce, ""))
	if err != nil {
		return nil, err
	}
	txStd.Signature = append(txStd.Signature, txs.Signature{
		Pubkey:    pri.PubKey(),
		Signature: sig,
		Nonce:     nonce,
	})

	return cliCtx.BroadcastTx(cdc.MustMarshalBinaryBare(txStd))
}

func transferBuilder(ctx context.CLIContext, senderStr, receiverStr string) (*txtransfer.TxTransfer, error) {
	senders, err := parseTransItem(ctx, senderStr)
	if err != nil {
		return nil, err
	}

	receivers, err := parseTransItem(ctx, receiverStr)
	if err != nil {
		return nil, err
	}

	return &txtransfer.TxTransfer{
		Senders:   senders,
		Receivers: receivers,
	}, nil
}

// parseTransItem  aadr,1qos,100aoe
func parseTransItem(cliCtx context.CLIContext, str string) (transtypes.TransItems, error) {
	items := make(transtypes.TransItems, 0)
	tis := strings.Split(str, ";")
	for _, ti := range tis {
		if ti == "" {
			continue
		}

		addrAndCoins := strings.Split(ti, ",")
		if len(addrAndCoins) < 2 {
			return nil, fmt.Errorf("`%s` not match rules", ti)
		}

		addr, err := qbaseaccount.GetAddrFromValue(cliCtx, addrAndCoins[0])
		if err != nil {
			return nil, err
		}
		qos, qscs, err := qostypes.ParseCoins(strings.Join(addrAndCoins[1:], ","))
		if err != nil {
			return nil, err
		}
		items = append(items, transtypes.TransItem{
			Address: addr,
			QOS:     qos,
			QSCs:    qscs,
		})
	}

	return items, nil
}

func getDefaultAccountNonce(ctx context.CLIContext, address []byte) (int64, error) {
	if ctx.NonceNodeURI == "" {
		return qbaseaccount.GetAccountNonce(ctx, address)
	}

	//NonceNodeURI不为空,使用NonceNodeURI查询account nonce值
	rpc := lib.TendermintClient(ctx.NonceNodeURI)
	newCtx := context.NewCLIContext().WithClient(rpc).WithCodec(ctx.Codec)

	return qbaseaccount.GetAccountNonce(newCtx, address)
}
