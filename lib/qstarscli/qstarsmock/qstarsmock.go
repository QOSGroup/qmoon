// Copyright 2018 The QOS Authors

package qstarsmock

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type QstarsMock struct {
	mux *http.ServeMux
}

func NewQstarsMock() QstarsMock {
	mux := http.NewServeMux()
	apis := []string{
		"node_version",
		"version",
		"kv/key",
		"accounts/address",
		"accounts",
		"accounts/address/send",
		"kv",
	}
	for _, v := range apis {
		vv := v
		mux.HandleFunc("/"+vv, func(w http.ResponseWriter, r *http.Request) {
			vvv := strings.Replace(vv, "/", "_", -1)
			f, err := os.Open("./qstarsmock/" + vvv + ".json")
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			defer f.Close()
			d, _ := ioutil.ReadAll(f)

			w.Write(d)
		})
	}

	return QstarsMock{
		mux: mux,
	}
}
func (tm QstarsMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm.mux.ServeHTTP(w, r)
}
