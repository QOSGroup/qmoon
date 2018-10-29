// Copyright 2018 The QOS Authors

package hadmin

import (
	"net/http"
	"testing"

	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/stretchr/testify/assert"
)

func TestAccountRegisterGin(t *testing.T) {
	body := accountRegisterQuery{
		Mail:     "test@123.com",
		Password: "1234556",
	}
	req, err := utils.NewPostJsonRequest(accountRegisterUrl, body)
	assert.Nil(t, err)

	resp := handler.CreateTestRequest(t, AccountRegisterGinRegister, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.True(t, handler.NotHasError(t, resp.Body.Bytes()))
}
