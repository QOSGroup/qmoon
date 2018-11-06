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
    data=`cat ${i//\//_}.json`
done


declare -a posts
posts=("${posts[*]}" "accounts")
posts=("${posts[*]}" "accounts/address/send")
posts=("${posts[*]}" "kv")

for i in ${posts[@]}
do
    curl -X POST -s ${tmHost}${i} -o ${i//\//_}.json
done


echo 'package qstarsmock' > qstarsmock.go

touch qstarsdata.go
echo 'package qstarsmock' > qstarsdata.go
for f in `ls *.json`
do
    echo var $f'=`' >> qstarsdata.go
    cat $f >> qstarsdata.go
    echo '`'>>qstarsdata.go
done


touch qstarsdata.go;echo 'package qstarsmock' > qstarsdata.go; echo 'var mockdata = map[string]string{' >> qstarsdata.go; for f in `ls *.json`; do  echo '"'$f'":`' >> qstarsdata.go ;cat $f >> qstarsdata.go;echo '`,'>>qstarsdata.go ; done; echo '}'>>qstarsdata.go
gofmt -w qstarsdata.go