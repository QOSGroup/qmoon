// Copyright 2018 The QOS Authors

package utils

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
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

func NewPutJsonRequest(url string, body interface{}) (*http.Request, error) {
	d, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(d)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	return req, nil
}

func CheckHttpHealthy(u string, timeout time.Duration) (ok bool, delay time.Duration, err error) {
	cli := &http.Client{
		Timeout: timeout,
	}
	s := time.Now()
	resp, err := cli.Head(u)
	if err != nil {
		return false, 0, err
	}

	e := time.Now()

	return resp.StatusCode == http.StatusOK, e.Sub(s), nil
}
