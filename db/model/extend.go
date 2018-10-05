// Copyright 2018 The QOS Authors

package model

import "database/sql"

// BlockChainByHeightRange 根据高度区间查询块信息
func BlockChainByHeightRange(db XODB, minHeight, maxHeight sql.NullInt64) ([]*BlockChain, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, height, data ` +
		`FROM public.block_chain ` +
		`WHERE height >= $1 and height <= $2`

	var bcs []*BlockChain

	// run query
	XOLog(sqlstr, minHeight, maxHeight)
	rows, err := db.Query(sqlstr, minHeight, maxHeight)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		bc := BlockChain{
			_exists: true,
		}
		err = rows.Scan(&bc.ID, &bc.Height, &bc.Data)
		if err != nil {
			return nil, err
		}
		bcs = append(bcs, &bc)
	}

	return bcs, nil
}
