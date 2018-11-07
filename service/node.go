// Copyright 2018 The QOS Authors

package service

import (
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
)

type NodeType struct {
	mnt   *model.NodeType
	mntrs []*model.NodeTypeRoute

	Name      string           `json:"name"`    // name
	BaseURL   string           `json:"baseUrl"` // base_url
	SecretKey string           `json:"-"`       // secret_key
	Routers   []*NodeTypeRoute `json:"routers"`
}

type NodeTypeRoute struct {
	mntr   *model.NodeTypeRoute
	Route  string `json:"route"`  // route
	Hidden bool   `json:"hidden"` // hidden
}

func covertToNodeType(mnt *model.NodeType) *NodeType {
	nmtrs, _ := model.NodeTypeRoutesByNodeTypeID(db.Db, utils.NullInt64(mnt.ID))
	var routers []*NodeTypeRoute
	for _, v := range nmtrs {
		routers = append(routers, covertToNodeTypeRoute(v))
	}

	return &NodeType{
		mnt:       mnt,
		mntrs:     nmtrs,
		Name:      mnt.Name.String,
		BaseURL:   mnt.BaseURL.String,
		SecretKey: mnt.SecretKey.String,
		Routers:   routers,
	}
}

func covertToNodeTypeRoute(mntr *model.NodeTypeRoute) *NodeTypeRoute {
	return &NodeTypeRoute{
		mntr:   mntr,
		Route:  mntr.Route.String,
		Hidden: mntr.Hidden.Bool,
	}
}

func AllNodeTypes() ([]*NodeType, error) {
	var res []*NodeType
	var offset, limit int64 = 0, 50
	for {

		nts, err := model.NodeTypeFilter(db.Db, "", offset, limit)
		if err != nil {
			break
		}

		if len(nts) == 0 {
			break
		}

		for _, v := range nts {
			res = append(res, covertToNodeType(v))
		}
		offset += limit

	}

	return res, nil
}

func DefaultRoute(nodeID int64) []NodeTypeRoute {
	var rs []NodeTypeRoute
	rs = append(rs, *covertToNodeTypeRoute(&model.NodeTypeRoute{
		NodeTypeID: utils.NullInt64(nodeID),
		Route:      utils.NullString(types.ExplorerRouteBlock.String()),
		Hidden:     utils.NullBool(false),
	}))

	rs = append(rs, *covertToNodeTypeRoute(&model.NodeTypeRoute{
		NodeTypeID: utils.NullInt64(nodeID),
		Route:      utils.NullString(types.ExplorerRouteValidtor.String()),
		Hidden:     utils.NullBool(false),
	}))

	rs = append(rs, *covertToNodeTypeRoute(&model.NodeTypeRoute{
		NodeTypeID: utils.NullInt64(nodeID),
		Route:      utils.NullString(types.ExplorerRouteNode.String()),
		Hidden:     utils.NullBool(false),
	}))

	return rs
}

func CreateNodeType(name, baseURL, secretKey string, routes []NodeTypeRoute) error {
	mnt := &model.NodeType{}
	mnt.Name = utils.NullString(name)
	mnt.BaseURL = utils.NullString(baseURL)
	mnt.SecretKey = utils.NullString(secretKey)
	mnt.CreatedAt = utils.NullTime(time.Now())
	err := mnt.Insert(db.Db)
	if err != nil {
		return err
	}

	if routes == nil && len(routes) == 0 {
		routes = DefaultRoute(mnt.ID)
	}

	for _, v := range routes {
		mntr := &model.NodeTypeRoute{}
		mntr.NodeTypeID = utils.NullInt64(mnt.ID)
		mntr.Route = utils.NullString(v.Route)
		mntr.Hidden = utils.NullBool(v.Hidden)

		err := mntr.Insert(db.Db)
		if err != nil {
			mnt.Delete(db.Db)
			return err
		}
	}

	return nil
}

func GetNodeTypeByName(name string) (*NodeType, error) {
	mnt, err := model.NodeTypeByName(db.Db, utils.NullString(name))
	if err != nil {
		return nil, err
	}

	return covertToNodeType(mnt), nil
}

func DeleteNodeTypeByName(name string) error {
	mnt, err := model.NodeTypeByName(db.Db, utils.NullString(name))
	if err != nil {
		return err
	}

	err = mnt.Delete(db.Db)

	return nil
}
