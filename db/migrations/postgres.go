// Copyright 2018 The QOS Authors

// Package pkg comments for pkg migrations
// migrations 数据库表结构
package migrations

import "database/sql"

// create database qmoon ENCODING 'UTF8' TEMPLATE template0;

var postgres = []*schema{
	{
		version: "init_schema",
		up: func(db *sql.DB) error {
			s := `
CREATE TABLE IF NOT EXISTS qmoon_status(
	id bigserial PRIMARY KEY,
	key varchar(64),
	value text
);
CREATE unique index qmoon_status_key_idx on qmoon_status(key);

CREATE TABLE IF NOT EXISTS node_types(
	id bigserial PRIMARY KEY,
	name varchar(128),
	base_url text,
	secret_key text,
	created_at timestamp with time zone
);
CREATE unique index node_types_name_idx on node_types(name);

create table node_type_route(
   id bigserial primary key,
   node_type_id BIGINT REFERENCES node_types (id) ON DELETE CASCADE,
   route varchar(64),
   hidden boolean default false
);
CREATE index node_type_route_node_type_id_idx on node_type_route(node_type_id);

CREATE TABLE IF NOT EXISTS accounts(
	id bigserial PRIMARY KEY,
	mail varchar(128),
	name varchar(64),
	avatar text,
	description text,
	status int,
	password varchar(128),
	created_at timestamp with time zone
);
CREATE unique index accounts_mail_idx on accounts(mail);
insert into accounts(mail, name, status, password)values('admin@test.local', 'admin', 1, MD5('123456'));

create table login_status(
   id bigserial primary key,
   account_id BIGINT REFERENCES accounts (id) ON DELETE CASCADE,
   login_type int2 DEFAULT 0, -- 登录类型:0:webx
   token varchar(64),
   expired_at timestamp with time zone
);
CREATE unique index login_status_account_id_idx on login_status(account_id);
CREATE unique index login_status_token_idx on login_status(token);

CREATE TABLE IF NOT EXISTS apps(
	id bigserial PRIMARY KEY,
	name varchar(64),
	secret_key text,
	status int,
    account_id BIGINT REFERENCES accounts (id) ON DELETE CASCADE,
	created_at timestamp with time zone
);
CREATE unique index apps_secret_key_idx on apps(secret_key);
CREATE index apps_account_id_idx on apps(account_id);
insert into apps(name, secret_key,status,account_id)values('t1', '123456',1,1); 

CREATE TABLE IF NOT EXISTS block_chain(
	id bigserial PRIMARY KEY,
	height bigint,
	data text
);
CREATE unique index block_chain_height_idx on block_chain(height);

CREATE TABLE IF NOT EXISTS blocks(
	id bigserial PRIMARY KEY,
	height bigint,
	data text
);
CREATE unique index blocks_height_idx on blocks(height);

insert into qmoon_status(key, value)values('qmoon_version', 'init_schema');

`
			_, err := db.Query(s)
			return err

		},
		down: func(db *sql.DB) error {
			s := `
DROP TABLE qmoon_status;
DROP TABLE node_types;
DROP TABLE node_type_route;
DROP TABLE apps;
DROP TABLE login_status;
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              DROP TABLE block_chain;
DROP TABLE blocks;
DROP TABLE accounts;

`
			_, err := db.Query(s)
			return err
		},
	},
}
