// Copyright 2018 The QOS Authors

package example

import (
	"database/sql"
	"errors"
)

type upFunc func(db *sql.DB) error
type downFunc func(db *sql.DB) error

type schema struct {
	version string
	up      upFunc
	down    downFunc
}

var schemaMap map[string][]*schema

func init() {
	schemaMap = make(map[string][]*schema)
	schemaMap["postgres"] = postgres
}

func getSchemasByDriver(driveName string) ([]*schema, error) {
	ss, ok := schemaMap[driveName]
	if !ok {
		return nil, errors.New("存储类型不支持")
	}

	l := len(ss)
	if l == 0 {
		return nil, errors.New("存储类型不支持")
	}

	return ss, nil
}

// DbInit 数据库结构初始化
func DbInit(driveName string, db *sql.DB) error {
	ss, err := getSchemasByDriver(driveName)
	if err != nil {
		return err
	}

	for _, v := range ss {
		if err := v.up(db); err != nil {
			return err
		}
	}

	return nil
}

// DbClear 数据库清除
func DbClear(driveName string, db *sql.DB) error {
	ss, err := getSchemasByDriver(driveName)
	if err != nil {
		return err
	}

	l := len(ss)
	for i := l; i >= 0; i-- {
		if err := ss[i].down(db); err != nil {
			return err
		}
	}

	return nil
}
