// Copyright 2018 The QOS Authors

// Package pkg comments for pkg migrations
// migrations ...
package migrations

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

func newestSchema(driveName string) (*schema, error) {
	ss, err := getSchemasByDriver(driveName)
	if err != nil {
		return nil, err
	}

	return ss[len(ss)-1], nil
}

func CurVersion(db *sql.DB) (string, error) {
	var version string
	s := "select value from qmoon_status where key='qmoon_version'"
	r := db.QueryRow(s)
	err := r.Scan(&version)
	if err != nil {
		return "", errors.New("读取版本失败" + err.Error())
	}

	return version, nil
}

// NeedMigration 是否需要迁移
func NeedMigration(driveName string, db *sql.DB) (bool, error) {
	s, err := newestSchema(driveName)
	if err != nil {
		return false, err
	}

	version, err := CurVersion(db)
	if err != nil {
		return false, err
	}

	if s.version != version {
		return true, nil
	}

	return false, nil
}

// Init 数据库结构初始化
func Init(driveName string, db *sql.DB) error {
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

// Up 数据库结构升级
func Up(driveName string, db *sql.DB) error {
	ss, err := getSchemasByDriver(driveName)
	if err != nil {
		return err
	}

	version, err := CurVersion(db)
	if err != nil {
		return Init(driveName, db)
	}

	var start bool
	for _, v := range ss {
		if start {
			if err := v.up(db); err != nil {
				return err
			}
		}
		if v.version == version {
			start = true
		}
	}

	return nil
}

// Up 数据库结构升级
func Down(driveName string, db *sql.DB) error {
	ss, err := getSchemasByDriver(driveName)
	if err != nil {
		return err
	}

	version, err := CurVersion(db)
	if err != nil {
		return err
	}

	for _, v := range ss {
		if v.version == version {
			return v.down(db)
		}
	}

	return nil
}
