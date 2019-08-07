package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type Node struct {
	Id          int64  `xorm:"pk autoincr BIGINT"`
	Name        string `xorm:"unique VARCHAR(128)"`
	BaseUrl     string `xorm:"TEXT"`
	SecretKey   string `xorm:"TEXT"`
	ChainId     string `xorm:"TEXT"`
	NodeType    string `xorm:"TEXT"`
	NodeVersion string `xorm:"TEXT"`
	Sort        int

	Version       int       `xorm:"version"`
	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Node) BeforeInsert() {
	n.CreatedAtUnix = time.Now().Unix()
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Node) BeforeUpdate() {
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Node) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		n.UpdatedAt = time.Unix(n.UpdatedAtUnix, 0).Local()
	}
}

func (n *Node) Insert() error {
	_, err := basex.Insert(n)
	if err != nil {
		return err
	}

	return nil
}

func CreateNode(name, baseURL, nodeType, nodeVersion, secretKey, chainID string) (*Node, error) {
	if err := newNode(name); err != nil {
		return nil, err
	}

	n := &Node{}
	n.Name = name
	n.BaseUrl = baseURL
	n.SecretKey = secretKey
	n.NodeVersion = nodeVersion
	n.ChainId = chainID
	n.NodeType = nodeType
	n.Sort = 1

	if _, err := basex.Insert(n); err != nil {
		return nil, err
	}

	return n, nil
}

func Nodes() ([]*Node, error) {
	var nodes = make([]*Node, 0)
	return nodes, basex.Asc("sort").Find(&nodes)
}

func RetrieveNodeByName(name string) (*Node, error) {
	n := &Node{Name: name}
	has, err := basex.Get(n)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NodeNotExist{Name: name}
	}

	return n, nil
}

func DeleteNodeByName(name string) error {
	n := &Node{Name: name}
	_, err := basex.Delete(n)

	return err
}
