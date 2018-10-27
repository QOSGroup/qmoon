// Copyright 2018 The QOS Authors

package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mail1 = "test1@test.local"
	pwd1  = "11111"

	mail2 = "test2@test.local"
	pwd2  = "22222"

	appName1 = "app1"
	appName2 = "app2"
	appName3 = "app3"
)

func TestCreateAccount(t *testing.T) {
	a1, err := CreateAccount(mail1, pwd1)
	assert.Nil(t, err)
	assert.Equal(t, a1.Mail, mail1)
	assert.True(t, a1.CheckPassword(pwd1))

	a2, err := CreateAccount(mail2, pwd2)
	assert.Nil(t, err)
	assert.NotEqual(t, a2.Mail, mail1)
	assert.False(t, a2.CheckPassword(pwd1))

	assert.Equal(t, a2.Mail, mail2)
	assert.True(t, a2.CheckPassword(pwd2))
}

func TestRetrieveAccount(t *testing.T) {
	a1, err := RetrieveAccountByMail(mail1)
	assert.Nil(t, err)
	a1.CheckPassword(pwd1)

	as, err := Accounts(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(as))

	as, err = Accounts(0, 100)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(as))

	a, err := RetrieveAccountByID(as[0].ID)
	assert.Nil(t, err)
	assert.Equal(t, as[0], a)

}

func TestCreateApp(t *testing.T) {
	a1, err := RetrieveAccountByMail(mail1)
	assert.Nil(t, err)
	a1.CheckPassword(pwd1)

	app1, err := a1.CreateApp(appName1)
	assert.Nil(t, err)
	assert.Equal(t, appName1, app1.Name)

	app2, err := a1.CreateApp(appName2)
	assert.Nil(t, err)
	assert.NotEqual(t, appName2, app1.Name)
	assert.Equal(t, appName2, app2.Name)

	a1.CreateApp(appName3)
}

func TestRetrieveApp(t *testing.T) {
	a1, err := RetrieveAccountByMail(mail1)
	assert.Nil(t, err)
	a1.CheckPassword(pwd1)

	apps, err := a1.Apps()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(apps))

	app1, err := a1.AppByID(apps[0].ID)
	assert.Nil(t, err)
	assert.Equal(t, appName1, app1.Name)

	tmpapp, err := AppBySecretKey(app1.SecretKey)
	assert.Nil(t, err)
	assert.Equal(t, app1, tmpapp)

	acc, err := tmpapp.Account()
	assert.Nil(t, err)
	assert.Equal(t, a1, acc)

}

func TestDeleteApp(t *testing.T) {
	a1, err := RetrieveAccountByMail(mail1)
	assert.Nil(t, err)
	a1.CheckPassword(pwd1)

	apps, err := a1.Apps()
	assert.Nil(t, err)

	app2, err := AppBySecretKey(apps[1].SecretKey)
	assert.Nil(t, err)
	assert.Equal(t, appName2, app2.Name)

	err = a1.DeleteByID(app2.ID)
	assert.Nil(t, err)

	apps, err = a1.Apps()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(apps))
	assert.Equal(t, appName1, apps[0].Name)
	assert.Equal(t, appName3, apps[1].Name)
}
