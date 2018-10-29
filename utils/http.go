// Copyright 2018 The QOS Authors

package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

func NewPostJsonRequest(url string, body interface{}) (*http.Request, error) {
	d, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(d)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	return req, nil
}
