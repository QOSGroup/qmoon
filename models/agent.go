package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Agent struct {
	Id       int64  `xorm:"pk autoincr BIGINT"`
	NodeType string `xorm:"TEXT"`
	BaseUrl  string `xorm:"TEXT"`

	Version       int       `xorm:"version"`
	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Agent) BeforeInsert() {
	n.CreatedAtUnix = time.Now().Unix()
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Agent) BeforeUpdate() {
	n.UpdatedAtUnix = time.Now().Unix()
}

func (n *Agent) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		n.UpdatedAt = time.Unix(n.UpdatedAtUnix, 0).Local()
	}
}

func (n *Agent) Insert() error {
	_, err := basex.Insert(n)
	if err != nil {
		return err
	}

	return nil
}

func CreateAgent(agentType, baseURL string) (*Agent, error) {
	n := &Agent{}
	n.BaseUrl = baseURL
	n.NodeType = agentType

	if _, err := basex.Insert(n); err != nil {
		return nil, err
	}

	return n, nil
}

func Agents() ([]*Agent, error) {
	var agents = make([]*Agent, 0)
	return agents, basex.Find(&agents)
}

func AgentsByNodeType(nodeType string) ([]*Agent, error) {
	var agents = make([]*Agent, 0)
	return agents, basex.Where("node_type = ?", nodeType).Find(&agents)
}
