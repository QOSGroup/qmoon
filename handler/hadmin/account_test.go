// Copyright 2018 The QOS Authors

package hadmin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/stretchr/testify/assert"
)

func TestAdminListAccountsGin(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, accountsUrl, nil)
	assert.Nil(t, err)

	var res []*account.Account
	_, err = handler.NewHttpTest(t, req).WithSession().Do(AccountGinRegister, &res)

	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestRetrieveAccountGin(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, accountsUrl, nil)
	assert.Nil(t, err)

	var accs []*account.Account
	_, err = handler.NewHttpTest(t, req).WithSession().WithLocalIP().Do(AccountGinRegister, &accs)
	assert.Nil(t, err)
	assert.NotEmpty(t, accs)

	acc := accs[0]
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", accountsUrl, acc.ID), nil)
	assert.Nil(t, err)

	var res account.Account
	_, err = handler.NewHttpTest(t, req).WithSession().Do(AccountGinRegister, &res)

	assert.Nil(t, err)
	assert.Equal(t, acc.Mail, res.Mail)
}

func TestUpdateAccountGin(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, accountsUrl, nil)
	assert.Nil(t, err)
	var accs []*account.Account
	_, err = handler.NewHttpTest(t, req).WithSession().WithLocalIP().Do(AccountGinRegister, &accs)
	assert.Nil(t, err)
	assert.NotEmpty(t, accs)

	body := updateAccountQuery{
		Name:        "new name",
		Avatar:      "new avatar",
		Description: "",
	}
	req, err = utils.NewPutJsonRequest(fmt.Sprintf("%s/%d", accountsUrl, accs[0].ID), body)
	assert.Nil(t, err)
	var res account.Account
	_, err = handler.NewHttpTest(t, req).WithSession().Do(AccountGinRegister, &res)
	assert.Nil(t, err)
	assert.NotEqual(t, accs[0].Name, res.Name)
	assert.NotEqual(t, accs[0].Avatar, res.Avatar)
	assert.NotEqual(t, accs[0].Name, res.Name)

	assert.Equal(t, body.Name, res.Name)
	assert.Equal(t, body.Avatar, res.Avatar)
	assert.Equal(t, body.Description, res.Description)
}
