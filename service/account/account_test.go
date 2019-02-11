// Copyright 2018 The QOS Authors

package account

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

	a2, err := CreateAccount(mail2, pwd2)
	assert.Nil(t, err)
	assert.NotEqual(t, a2.Mail, mail1)

	assert.Equal(t, a2.Mail, mail2)
}

func TestRetrieveAccount(t *testing.T) {
	a1, err := RetrieveAccountByMail(mail1)
	assert.Nil(t, err)
	assert.Equal(t, a1.Mail, mail1)

	as, err := List(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(as))

	as, err = List(0, 100)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(as))

	a, err := RetrieveAccountByID(as[0].ID)
	assert.Nil(t, err)
	assert.Equal(t, as[0], a)
}
