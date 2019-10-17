package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type QmoonStatus struct {
	Id    int64  `xorm:"pk autoincr BIGINT"`
	Key   string `xorm:"unique VARCHAR(64)"`
	Value string `xorm:"TEXT"`

	Version       int       `xorm:"version"`
	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (qs *QmoonStatus) BeforeInsert() {
	qs.CreatedAtUnix = time.Now().Unix()
	qs.UpdatedAtUnix = time.Now().Unix()
}

func (qs *QmoonStatus) BeforeUpdate() {
	qs.UpdatedAtUnix = time.Now().Unix()
}

func (qs *QmoonStatus) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		qs.CreatedAt = time.Unix(qs.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		qs.UpdatedAt = time.Unix(qs.UpdatedAtUnix, 0).Local()
	}
}

func (qs *QmoonStatus) Insert() error {
	_, err := basex.Insert(qs)
	if err != nil {
		return err
	}

	return nil
}

func (qs *QmoonStatus) Update() error {
	num, err := basex.ID(qs.Id).Cols("value", "updated_at_unix").Update(qs)
	if err != nil {
		return err
	}

	if num != 1 {
		return errors.New("no data be updated")
	}

	return nil
}

func RetrieveQmoonStatusByKey(key string) (*QmoonStatus, error) {
	qs := &QmoonStatus{Key: key}
	has, err := basex.Get(qs)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "QmoonStatus"}
	}

	return qs, nil
}



func DeleteKey(key string)  error{

	var l QmoonStatus
	_, err := basex.Where("key = ?", key).Delete(&l)
	if err != nil {
		return err
	}

	return nil
}




func DeleteKeyBySystemName(systemname string)  error {
	var l QmoonStatus
	_, err := basex.Where("system_name = ?", systemname).Delete(&l)
	if err != nil {
		return err
	}
	return nil
}