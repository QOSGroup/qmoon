// Copyright 2018 The QOS Authors

package transfer

import "database/sql"

var postgres = []*schema{
	{
		version: "init_schema",
		up: func(db *sql.DB) error {
			s := `
CREATE TABLE IF NOT EXISTS tx_transfer(
	id bigserial PRIMARY KEY,
	chain_id text,
	height bigint,
	hash text,
	address text,
	coin text,
	amount text,
	type int2, --0:send 1:recieve
	time timestamp with time zone 
);
CREATE index IF NOT EXISTS tx_transfer_address_idx on tx_transfer(address);
`
			_, err := db.Query(s)
			return err

		},
		down: func(db *sql.DB) error {
			s := `
DROP TABLE tx_transfer;
`
			_, err := db.Query(s)
			return err
		},
	},
}
