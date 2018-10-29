// Copyright 2018 The QOS Authors

package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

func CreateTestRequest(t *testing.T, f func(r *gin.Engine), req *http.Request) *httptest.ResponseRecorder {
	r := gin.Default()
	f(r)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	t.Logf("resp:%+v", w)

	return w
}

func NotHasError(t *testing.T, respBody []byte) bool {
	var resp types.RPCResponse
	err := json.Unmarshal(respBody, &resp)
	if err != nil {
		t.Logf("json Unmarshal err, %s", err.Error())
		return false
	}

	if resp.Error != nil {
		t.Logf("respBody:%s", string(respBody))
	}

	return resp.Error == nil
}
