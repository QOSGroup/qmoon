#!/bin/sh

tmHost="http://18.188.103.180:26657/"

declare -a apis
apis=("${apis[*]}" "abci_info")
apis=("${apis[*]}" "consensus_state")
apis=("${apis[*]}" "dump_consensus_state")
apis=("${apis[*]}" "genesis")
apis=("${apis[*]}" "health")
apis=("${apis[*]}" "net_info")
apis=("${apis[*]}" "num_unconfirmed_txs")
apis=("${apis[*]}" "status")
apis=("${apis[*]}" "abci_query")
apis=("${apis[*]}" "block")
apis=("${apis[*]}" "block_results")
apis=("${apis[*]}" "blockchain")
apis=("${apis[*]}" "broadcast_tx_async")
apis=("${apis[*]}" "broadcast_tx_commit")
apis=("${apis[*]}" "broadcast_tx_sync")
apis=("${apis[*]}" "commit")
apis=("${apis[*]}" "subscribe")
apis=("${apis[*]}" "tx")
apis=("${apis[*]}" "tx_search")
apis=("${apis[*]}" "unconfirmed_txs")
apis=("${apis[*]}" "unsubscribe")
apis=("${apis[*]}" "unsubscribe_all")
apis=("${apis[*]}" "validators")


for i in ${apis[@]}
do
    curl -s ${tmHost}${i} -o ${i}.json
done

