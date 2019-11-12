// Copyright 2018 The QOS Authors

package qstarsmock

import (
	"fmt"
	"net/http"
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
			d, ok := mockdata[vvv+".json"]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Write([]byte(d))
		})
	}

	return QstarsMock{
		mux: mux,
	}
}
func (tm QstarsMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("---url:%s\n", r.RequestURI)
	tm.mux.ServeHTTP(w, r)
}
