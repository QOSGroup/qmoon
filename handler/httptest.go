// Copyright 2018 The QOS Authors

package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	tmltypes "github.com/tendermint/tendermint/rpc/lib/types"
)

const (
	TestInternalUser         = "test1@internal.local"
	TestInternalUserPassword = "123456"
)

func CreateTestUser() error {
	acc, err := account.CreateAccount(TestInternalUser, TestInternalUserPassword)
	if err != nil {
		panic(err)
	}
	_, err = acc.CreateApp("app")
	if err != nil {
		panic(err)
	}

	return nil
}

type HttpTest struct {
	t *testing.T
	h http.Handler

	req *http.Request
}

func NewHttpTest(t *testing.T, req *http.Request) *HttpTest {
	var ht HttpTest
	ht.t = t
	ht.req = req

	return &ht
}

func (ht *HttpTest) WithSession() *HttpTest {
	t, err := service.Login(TestInternalUser, TestInternalUserPassword)
	if err != nil {
		panic(err)
	}
	ht.req.Header.Set(types.TokenKey, t.Token)

	return ht
}

func (ht *HttpTest) WithAuth() *HttpTest {
	acc, err := account.RetrieveAccountByMail(TestInternalUser)
	if err != nil {
		panic(err)
	}
	apps, err := acc.Apps()
	if err != nil {
		panic(err)
	}

	ht.req.Header.Set(types.AuthKey, apps[0].SecretKey)

	return ht
}

func (ht *HttpTest) Do(f func(r *gin.Engine), v interface{}) (*httptest.ResponseRecorder, error) {
	r := gin.Default()
	f(r)

	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, ht.req)

	ht.t.Logf("request: %s %s, response.body:%s", ht.req.Method, ht.req.URL.String(), rw.Body.String())
	if w, ok := v.(io.Writer); ok {
		io.Copy(w, rw.Body)
	} else {
		var tmresp tmltypes.RPCResponse
		err := json.NewDecoder(rw.Body).Decode(&tmresp)
		if err != nil {
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			} else {
				return rw, err
			}
		}

		if tmresp.Error != nil {
			return rw, tmresp.Error
		}

		err = lib.Cdc.UnmarshalJSON(tmresp.Result, v)
		if err != nil {
			return rw, err
		}
	}

	return rw, nil
}
