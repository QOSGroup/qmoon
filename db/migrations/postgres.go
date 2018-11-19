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

--- 
CREATE TABLE IF NOT EXISTS genesis(
	id bigserial PRIMARY KEY,
	chain_id text,
	genesis_time timestamp with time zone,
	data text,
	created_at timestamp with time zone
);
CREATE unique index genesis_chain_id_idx on genesis(chain_id);

CREATE TABLE IF NOT EXISTS nodes(
	id bigserial PRIMARY KEY,
	name varchar(128),
	base_url text,
	secret_key text,
	chain_id text,
	genesis_id BIGINT REFERENCES genesis(id),
	created_at timestamp with time zone
);
CREATE unique index nodes_name_idx on nodes(name);
CREATE index nodes_genesis_id_idx on nodes(genesis_id);

create table node_route(
   id bigserial primary key,
   node_id BIGINT REFERENCES nodes (id) ON DELETE CASCADE,
   route varchar(64),
   hidden boolean default false
);
CREATE index node_route_node_id_idx on node_route(node_id);

CREATE TABLE IF NOT EXISTS tm_block_chain(
	id bigserial PRIMARY KEY,
	chain_id varchar(64),
	height bigint,
    num_txs bigint,
	data text,
	time timestamp with time zone,
	created_at timestamp with time zone
);
CREATE index tm_block_chain_chain_id_idx on tm_block_chain(chain_id);
CREATE unique index tm_block_chain_chain_id_height_idx on tm_block_chain(chain_id, height);

CREATE TABLE IF NOT EXISTS tm_blocks(
	id bigserial PRIMARY KEY,
	chain_id text,
	height bigint,
	data text,
	created_at timestamp with time zone
);
CREATE unique index tm_blocks_chain_id_height_idx on tm_blocks(chain_id, height);
CREATE index tm_blocks_chain_id_idx on tm_blocks(chain_id);

CREATE TABLE IF NOT EXISTS blocks(
	id bigserial PRIMARY KEY,
	chain_id text,
	height bigint,
    num_txs bigint,
	total_txs bigint,
	data_hash varchar(40),
	validators_num bigint,
	validators_total bigint,
	validators_hash varchar(40),
	time timestamp with time zone,
	created_at timestamp with time zone
);
CREATE unique index blocks_chain_id_height_idx on blocks(chain_id, height);
CREATE index blocks_chain_id_idx on blocks(chain_id);

CREATE TABLE IF NOT EXISTS validators(
	id bigserial PRIMARY KEY,
	chain_id text,
	address varchar(40),
	pub_key_type text,
	pub_key_value text,
	voting_power bigint,
	accum bigint,
	first_block_height bigint,
	first_block_time timestamp with time zone,
	created_at timestamp with time zone
);
CREATE unique index validators_address_idx on validators(address);
CREATE index validators_chain_id_idx on validators(chain_id);

CREATE TABLE IF NOT EXISTS block_validators(
	id bigserial PRIMARY KEY,
	chain_id text,
	height bigint,
	validator_address varchar(40),
	validator_index bigint,
	type bigint,
	round bigint,
	signature text,
	voting_power bigint,
	accum bigint,
	time timestamp with time zone,
	created_at timestamp with time zone
);
CREATE index block_validators_chain_id_height_idx on block_validators(chain_id, height);
CREATE index block_validators_address_idx on block_validators(validator_address);

CREATE TABLE IF NOT EXISTS txs(
	id bigserial PRIMARY KEY,
	chain_id text,
	height bigint,
	tx_type text,
	index bigint,
	maxgas bigint,
	qcp_from text,
	qcp_to text,
	qcp_sequence bigint,
	qcp_txindex bigint,
	qcp_isresult boolean,
	origin_tx text,
	json_tx text,
	time timestamp with time zone,
	created_at timestamp with time zone
);
CREATE unique index txs_chain_id_height_index_idx on txs(chain_id, height, index);
CREATE index txs_chain_id_height_idx on txs(chain_id, height);
CREATE index txs_chain_id_idx on txs(chain_id);

CREATE TABLE IF NOT EXISTS peers(
	id bigserial PRIMARY KEY,
	chain_id text,
	moniker text,
	peer_id varchar(40),
	listen_addr text,
	network text,
	version text,
	channels text,
	send_start timestamp with time zone,
	recv_start timestamp with time zone,
	created_at timestamp with time zone
);
CREATE index peers_chain_id_idx on peers(chain_id);
CREATE unique index peers_peer_id_idx on peers(peer_id);

CREATE TABLE IF NOT EXISTS consensus_state(
	id bigserial PRIMARY KEY,
	chain_id text,
	height text,
	round text,
	step text,
	prevotes_num bigint,
	prevotes_value text,
	precommits_num bigint,
	precommits_value text,
	start_time text
);
CREATE unique index consensus_state_chain_id_idx on consensus_state(chain_id);


insert into qmoon_status(key, value)values('qmoon_version', 'init_schema');
`
			_, err := db.Query(s)
			return err

		},
		down: func(db *sql.DB) error {
			s := `
DROP TABLE qmoon_status;
DROP TABLE nodes;
DROP TABLE node_route;
DROP TABLE apps;
DROP TABLE genesis;
DROP TABLE login_status;
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              DROP TABLE block_chain;
DROP TABLE blocks;
DROP TABLE block_validators;
DROP TABLE txs;
DROP TABLE accounts;
DROP TABLE peers;
DROP TABLE consensus_state;

`
			_, err := db.Query(s)
			return err
		},
	},
}
