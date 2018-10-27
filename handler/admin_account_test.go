// Copyright 2018 The QOS Authors

package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAdminAccountCRUD(t *testing.T) {
	r := gin.Default()
	AdminAccountGinRegister(r)

	s := httptest.NewServer(r)
	defer s.Close()
	t.Logf("url:%s", s.URL)
}
