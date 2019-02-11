package models

type NodeRoute struct {
	Id     int64  `xorm:"pk autoincr BIGINT"`
	NodeId int64  `xorm:"index BIGINT"`
	Route  string `xorm:"VARCHAR(64)"`
	Hidden bool   `xorm:"default false BOOL"`
}
