// Copyright 2018 The QOS Authors

package utils

import (
	"crypto/hmac"
	"crypto/sha1"
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
