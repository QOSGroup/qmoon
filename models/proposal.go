package models

import "time"

type proposal struct {
	Id        int64  `xorm:"pk autoincr BIGINT"`
	Tx        string `xorm:"unique(fee_tx_idx) TEXT"`
	GasWanted int64  `xorm:"BIGINT"`
	GasUsed   int64  `xorm:"BIGINT"`
	Fee       string `xorm:"TEXT"`
	ChainID   string `xorm:"-"`

	Version       int       `xorm:"version"`
	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}
