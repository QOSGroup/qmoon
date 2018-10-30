// Copyright 2018 The QOS Authors

package hdata

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/lib/qstarscli"
	"github.com/QOSGroup/qmoon/lib/tmcli"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	dbTest := db.NewTestDb(m)
	defer dbTest.Close()

	err := handler.CreateTestUser()
	if err != nil {
		panic(err)
	}

	tq := qstarscli.NewTestQstarsServer()
	defer tq.Close()

	tts := tmcli.NewTestTmServer()
	defer tts.Close()

	m.Run()
}

func createTestRequest(t *testing.T, f func(r *gin.Engine), req *http.Request) *httptest.ResponseRecorder {
	r := gin.Default()
	f(r)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	t.Logf("resp:%+v", w)

	return w
}

func notHasError(t *testing.T, respBody []byte) bool {
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
