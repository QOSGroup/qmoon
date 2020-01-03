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
	"strconv"
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

		maxId, _ := strconv.ParseInt(c.Query("maxHeight"), 10, 64)
		minId, _ := strconv.ParseInt(c.Query("minHeight"), 10, 64)
		offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		result, err := node.TxByHash(hash, minId, maxId, limit, offset)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		tx_msg, err := qos.NewQosCli("").QueryTx(node.BaseURL, hash)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		if (result.Status == 1 && tx_msg.Code != 0 ) || (result.Status == 0 && tx_msg.Code == 0 ) {
			// fmt.Println("tx_msg", tx_msg)
			if tx_msg.Code == 0 {
				result.Status = 1
				result.TxStatus = "Success";
			} else {
				result.Status = int(tx_msg.Code)
			}
			err = models.UpdateTxStatusByHash(node.ChainID, int(tx_msg.Code), hash, tx_msg.GasUsed, tx_msg.GasWanted)
		}
		if tx_msg.RawLog != "" && strings.Contains(tx_msg.RawLog, "\"message\":\"") {
			ss := strings.Split(tx_msg.RawLog, "\"message\":\"")
			result.TxStatus = ss[1][:strings.Index(ss[1], "\"")]
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}