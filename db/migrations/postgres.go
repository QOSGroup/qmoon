// Copyright 2018 The QOS Authors

// Package pkg comments for pkg migrations
// migrations 数据库表结构
package migrations

import "database/sql"

// create database qmoon ENCODING 'UTF8' TEMPLATE template0;

var postgres = []*schema{{
	version: "2018-10-01-0800_init_schema",
	up: func(db *sql.DB) error {
		s := `
CREATE TABLE IF NOT EXISTS qmoon_status(
	id bigserial PRIMARY KEY,
	key varchar(64),
	value text,
	unique (key)
);

CREATE TABLE IF NOT EXISTS block_chain(
	id bigserial PRIMARY KEY,
	height bigint,
	data text,
	unique (height)
);

CREATE TABLE IF NOT EXISTS blocks(
	id bigserial PRIMARY KEY,
	height bigint,
	data text,
	unique (height)
);

insert into qmoon_status(key, value)values('qmoon_version', '2018-10-01-0800_init_schema');

`
		_, err := db.Query(s)
		return err

	},
	down: func(db *sql.DB) error {
		s := `
DROP TABLE qmoon_status;
DROP TABLE block_chain;
DROP TABLE blocks;
`
		_, err := db.Query(s)
		return err
	},
},
	{
		version: "2018-10-22-1800",
		up: func(db *sql.DB) error {
			s := `
CREATE TABLE IF NOT EXISTS accounts(
	id bigserial PRIMARY KEY,
	secret_id varchar(64),
	secret_key text,
	description text,
	created_at timestamp with time zone,
	unique (secret_id)
);

update qmoon_status set value = '2018-10-22-1800' where key='qmoon_version';
`
			_, err := db.Query(s)
			return err

		},
		down: func(db *sql.DB) error {
			s := `
DROP TABLE accounts;
update qmoon_status set value = '2018-10-01-0800_init_schema' where key='qmoon_version';
`
			_, err := db.Query(s)
			return err
		},
	},
}
