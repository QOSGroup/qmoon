#!/bin/sh

set -x

tmHost="http://127.0.0.1:1317/"

declare -a apis
apis=("${apis[*]}" "node_version")
apis=("${apis[*]}" "version")
apis=("${apis[*]}" "kv/key")
apis=("${apis[*]}" "accounts/address")

for i in ${apis[@]}
do
    curl -s ${tmHost}${i} -o ${i//\//_}.json
done


declare -a posts
posts=("${posts[*]}" "accounts")
posts=("${posts[*]}" "accounts/address/send")
posts=("${posts[*]}" "kv")

for i in ${posts[@]}
do
    curl -X POST -s ${tmHost}${i} -o ${i//\//_}.json
done


