// Copyright 2018 The QOS Authors

package example

import "database/sql"

var postgres = []*schema{
	{
		up: func(db *sql.DB) error {
			return nil
		},
		down: func(db *sql.DB) error {
			return nil
		},
	},
}
