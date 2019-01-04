// Copyright 2018 The QOS Authors

package atm

import "database/sql"

var postgres = []*schema{
	{
		up: func(db *sql.DB) error {
			s := `
CREATE TABLE IF NOT EXISTS atm_ip_record(
    id bigserial PRIMARY KEY,
    chainid text,
    ip text,
    amount int,
    createat timestamp with time zone
);
CREATE unique index IF NOT EXISTS atm_ip_record_ip_chainid_createat_idx on atm_ip_record(ip, chainid, createat);


CREATE TABLE IF NOT EXISTS atm_record(
	id bigserial PRIMARY KEY,
	address text,
	chainid text,
	coin text,
	amount text,
	height text,
	hash text,
	createat timestamp with time zone 
);
CREATE unique index IF NOT EXISTS atm_record_address_chainid_idx on atm_record(address, chainid, createat);
`
			_, err := db.Query(s)
			return err

		},
		down: func(db *sql.DB) error {
			s := `
DROP TABLE atm_ip_record;
DROP TABLE atm_record;
`
			_, err := db.Query(s)
			return err
		},
	},
}
