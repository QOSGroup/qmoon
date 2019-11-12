// Copyright 2018 The QOS Authors

// Package pkg comments for pkg auth
// auth 请求鉴权
package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/QOSGroup/qmoon/utils"
)

type SignData struct {
	SecretID    string
	expiredTime int
	currentTime int
	rand        string

	origin string
	en     []byte
}

func ParseSign(sign []byte) (*SignData, error) {
	sd := new(SignData)
	signDecoded, err := base64.StdEncoding.DecodeString(string(sign))
	if err != nil {
		return nil, err
	}

	if len(signDecoded) < 20 {
		return nil, errors.New("1invalid sign")
	}

	sd.origin = string(signDecoded[20:])
	v, err := url.ParseQuery(sd.origin)
	if err != nil {
		return nil, errors.New("2invalid sign")
	}
	sd.SecretID = v.Get("s")

	expiredTime, err := strconv.ParseInt(v.Get("e"), 10, 32)
	if err != nil {
		return nil, errors.New("3invalid sign")
	}

	currentTime, err := strconv.ParseInt(v.Get("t"), 10, 32)
	if err != nil {
		return nil, errors.New("4invalid sign")
	}
	if currentTime >= expiredTime {
		return nil, errors.New("sign has expired. ")
	}

	if time.Now().Unix() >= expiredTime {
		return nil, errors.New("sign has expired. ")
	}

	sd.expiredTime = int(expiredTime)
	sd.currentTime = int(currentTime)
	sd.rand = v.Get("r")

	if len(sd.rand) > 10 {
		return nil, errors.New("r is too long. ")
	}

	return sd, nil
}

func (sd SignData) String() string {
	return fmt.Sprintf("s=%s&e=%d&t=%d&r=%s", sd.SecretID, sd.expiredTime, sd.currentTime, sd.rand)
}

func (sd SignData) Sign(key []byte) ([]byte, error) {
	hs, err := utils.HmacSha1(key, []byte(sd.String()))
	if err != nil {
		return nil, err
	}

	en := base64.StdEncoding.EncodeToString(append(hs, []byte(sd.String())...))

	return []byte(en), nil
}
