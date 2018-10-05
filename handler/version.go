// Copyright 2018 The QOS Authors

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/QOSGroup/qmoon/version"
)

type VersionRes struct {
	Version string `json:"version"`
}

func ServerInfoHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		si := &VersionRes{
			Version: version.Version,
		}
		d, _ := json.Marshal(si)
		res.Write(d)
	}
}
