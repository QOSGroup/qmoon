// Copyright 2018 The QOS Authors

package service

import (
	"fmt"
	"strings"
	"time"

	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/tendermint/go-amino"
)

type Node struct {
	mnt   *model.Node
	mntrs []*model.NodeRoute

	Name      string       `json:"name"`    // name
	BaseURL   string       `json:"baseUrl"` // base_url
	SecretKey string       `json:"-"`       // secret_key
	Routers   []*NodeRoute `json:"routers"`
	ChanID    string       `json:"chanId"`
}

type NodeRoute struct {
	mntr   *model.NodeRoute
	Route  string `json:"route"`  // route
	Hidden bool   `json:"hidden"` // hidden
}

func covertToNode(mnt *model.Node) *Node {
	nmtrs, _ := model.NodeRoutesByNodeID(db.Db, utils.NullInt64(mnt.ID))
	var routers []*NodeRoute
	for _, v := range nmtrs {
		routers = append(routers, covertToNodeRoute(v))
	}

	g, _ := mnt.Genesi(db.Db)

	return &Node{
		mnt:       mnt,
		mntrs:     nmtrs,
		Name:      mnt.Name.String,
		BaseURL:   mnt.BaseURL.String,
		SecretKey: mnt.SecretKey.String,
		ChanID:    g.ChainID.String,
		Routers:   routers,
	}
}

func covertToNodeRoute(mntr *model.NodeRoute) *NodeRoute {
	return &NodeRoute{
		mntr:   mntr,
		Route:  mntr.Route.String,
		Hidden: mntr.Hidden.Bool,
	}
}

func (n Node) AppState(cdc *amino.Codec) (*qbasetypes.GenesisState, error) {
	tmc := lib.TendermintClient(n.BaseURL)
	genesis, err := tmc.Genesis()
	if err != nil {
		return nil, fmt.Errorf("retrieve node genesis err:%s", err)
	}

	gs := qbasetypes.GenesisState{}
	err = cdc.UnmarshalJSON(genesis.Genesis.AppState, &gs)
	if err != nil {
		return nil, err
	}

	return &gs, nil
}

func AllNodes() ([]*Node, error) {
	var res []*Node
	var offset, limit int64 = 0, 50
	for {
		nts, err := model.NodeFilter(db.Db, "", offset, limit)
		if err != nil {
			break
		}

		if len(nts) == 0 {
			break
		}

		for _, v := range nts {
			res = append(res, covertToNode(v))
		}
		offset += limit
	}

	return res, nil
}

func DefaultRoute(nodeID int64) []NodeRoute {
	var rs []NodeRoute
	rs = append(rs, *covertToNodeRoute(&model.NodeRoute{
		NodeID: utils.NullInt64(nodeID),
		Route:  utils.NullString(types.ExplorerRouteBlock.String()),
		Hidden: utils.NullBool(false),
	}))

	rs = append(rs, *covertToNodeRoute(&model.NodeRoute{
		NodeID: utils.NullInt64(nodeID),
		Route:  utils.NullString(types.ExplorerRouteValidtor.String()),
		Hidden: utils.NullBool(false),
	}))

	rs = append(rs, *covertToNodeRoute(&model.NodeRoute{
		NodeID: utils.NullInt64(nodeID),
		Route:  utils.NullString(types.ExplorerRouteNode.String()),
		Hidden: utils.NullBool(false),
	}))

	return rs
}

func CreateNode(name, baseURL, secretKey string, routes []NodeRoute) error {
	if strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL[:len(baseURL)-1]
	}

	g, err := AddGenesis(baseURL)
	if err != nil {
		return err
	}

	mnt := &model.Node{}
	mnt.Name = utils.NullString(name)
	mnt.BaseURL = utils.NullString(baseURL)
	mnt.SecretKey = utils.NullString(secretKey)
	mnt.CreatedAt = utils.NullTime(time.Now())
	mnt.GenesisID = utils.NullInt64(g.ID)
	mnt.ChainID = utils.NullString(g.ChainID.String)
	if err := mnt.Insert(db.Db); err != nil {
		return err
	}

	if routes == nil && len(routes) == 0 {
		routes = DefaultRoute(mnt.ID)
	}

	for _, v := range routes {
		mntr := &model.NodeRoute{}
		mntr.NodeID = utils.NullInt64(mnt.ID)
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

func GetNodeByName(name string) (*Node, error) {
	mnt, err := model.NodeByName(db.Db, utils.NullString(name))
	if err != nil {
		return nil, err
	}

	return covertToNode(mnt), nil
}

func DeleteNodeByName(name string) error {
	mnt, err := model.NodeByName(db.Db, utils.NullString(name))
	if err != nil {
		return err
	}

	err = mnt.Delete(db.Db)

	return nil
}
