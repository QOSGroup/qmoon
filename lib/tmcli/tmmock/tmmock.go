// Copyright 2018 The QOS Authors

package tmmock

import (
	"net/http"
	"strings"
)

type TendermintMock struct {
	mux *http.ServeMux
}

func NewTendermintMock() TendermintMock {
	mux := http.NewServeMux()
	apis := []string{
		"abci_info",
		"consensus_state",
		"dump_consensus_state",
		"genesis",
		"health",
		"net_info",
		"num_unconfirmed_txs",
		"status",
		"abci_query",
		"block",
		"block_results",
		"blockchain",
		"broadcast_tx_async",
		"broadcast_tx_commit",
		"broadcast_tx_sync",
		"commit",
		"subscribe",
		"tx",
		"tx_search",
		"unconfirmed_txs",
		"unsubscribe",
		"unsubscribe_all",
		"validators",
	}
	for _, v := range apis {
		vv := v
		mux.HandleFunc("/"+v, func(w http.ResponseWriter, r *http.Request) {
			vvv := strings.Replace(vv, "/", "_", -1)
			d, ok := mockdata[vvv+".json"]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Write([]byte(d))
		})
	}

	return TendermintMock{
		mux: mux,
	}
}
func (tm TendermintMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm.mux.ServeHTTP(w, r)
}
