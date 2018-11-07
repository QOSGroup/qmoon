// Copyright 2018 The QOS Authors

package hadmin

import (
	"net/http"
	"testing"

	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/stretchr/testify/assert"
)

func TestRegisterGinEmptyMail(t *testing.T) {

	body := registerQuery{
		Email:    "",
		Password: "1234556",
	}
	req, err := utils.NewPostJsonRequest(registerUrl, body)
	assert.Nil(t, err)

	var res account.Account
	_, err = handler.NewHttpTest(t, req).Do(RegisterGinRegister, &res)
	assert.NotNil(t, err)
}

func TestRegisterGinEmptyPassword(t *testing.T) {
	body := registerQuery{
		Email:    "test@123.com",
		Password: "",
	}
	req, err := utils.NewPostJsonRequest(registerUrl, body)
	assert.Nil(t, err)

	var res account.Account
	_, err = handler.NewHttpTest(t, req).Do(RegisterGinRegister, &res)
	assert.NotNil(t, err)
}

func TestRegisterGinSuccess(t *testing.T) {

	body := registerQuery{
		Email:    "test@123.com",
		Password: "1234556",
	}
	req, err := utils.NewPostJsonRequest(registerUrl, body)
	assert.Nil(t, err)

	var res account.Account
	resp, err := handler.NewHttpTest(t, req).Do(RegisterGinRegister, &res)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
}
