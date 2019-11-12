// Copyright 2018 The QOS Authors

package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// HmacSha1 hmac sha1 加密
func HmacSha1(key, value []byte) ([]byte, error) {
	h := hmac.New(sha1.New, key)
	_, err := h.Write(value)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// EncryptPwd 密码加密
func EncryptPwd(pwd []byte) string {
	return fmt.Sprintf("%x", md5.Sum(pwd))
}

// MD5 use md5 encrypt
func MD5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Base64En encod []byte
func Base64En(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64De encod string
func Base64De(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func Sha256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func Sha1(data string) string {
	h := sha1.New()
	io.WriteString(h, data)
	md := h.Sum(nil)
	return hex.EncodeToString(md)
}
