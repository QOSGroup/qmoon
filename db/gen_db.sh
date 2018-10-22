#!/bin/bash

DBUSER=postgres
DBPasS=
DBHOST=localhost
DBNAME=qmoon

DB=pgsql://$DBUSER:$DBPasS@$DBHOST/$DBNAME?sslmode=disable

set -x

mkdir -p model

# 各个表model
xo $DB -o model --suffix .go --template-path xo/templates/

