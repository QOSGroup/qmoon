// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const txHashUrl = NodeProxy + "/txs/:hash"

func init() {
	hdataHander[txHashUrl] = TxHashGinRegister
}

// TxHashGinRegister 注册tx
func TxHashGinRegister(r *gin.Engine) {
	r.GET(txHashUrl, middleware.ApiAuthGin(), txHashGin())
}

func txHashGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		hash := strings.ToUpper(c.Param("hash"))
		result, err := node.TxByHash(hash)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		if result.Status != 1 {
			tx_msg, err := qos.NewQosCli("").QueryTx(node.BaseURL, hash)
			// fmt.Println("tx_msg", tx_msg)
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}
			if tx_msg.Code == 0 && (tx_msg.GasUsed > 0 && tx_msg.GasWanted > 0) {
				result.TxStatus = "Success";
				result.Status = 1
				result.GasWanted = tx_msg.GasWanted
				result.GasUsed = tx_msg.GasUsed
				err = models.UpdateTxStatusByHash(node.ChainID, result.Status, hash, tx_msg.GasUsed, tx_msg.GasWanted)
			} else if result.GasUsed != tx_msg.GasUsed || result.GasWanted != tx_msg.GasWanted {
				result.GasWanted = tx_msg.GasWanted
				result.GasUsed = tx_msg.GasUsed
				err = models.UpdateTxStatusByHash(node.ChainID, result.Status, hash, tx_msg.GasUsed, tx_msg.GasWanted)
			}
			if tx_msg.RawLog == "" || strings.Index(tx_msg.RawLog, "\"message\":\"") < 0{
				result.TxStatus = "Can't find error message in transaction result";
				result.Status = int(tx_msg.Code)
			} else {
				result.Status = int(tx_msg.Code)
				ss := strings.Split(tx_msg.RawLog, "\"message\":\"")
				result.TxStatus = ss[1][:strings.Index(ss[1],"\"")]
			}
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}