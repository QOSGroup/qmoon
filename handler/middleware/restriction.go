// Copyright 2018 The QOS Authors

package middleware

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPWhitelist(ipNets []*net.IPNet) func(c *gin.Context) {
	return func(c *gin.Context) {
		reqIP := net.ParseIP(c.ClientIP())
		var isContained bool
		for _, ip := range ipNets {
			if ip.Contains(reqIP) {
				isContained = true
				break
			}
		}

		if !isContained {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
