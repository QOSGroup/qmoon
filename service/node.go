// Copyright 2018 The QOS Authors

package service

import (
	"errors"
	"fmt"
	"strings"

	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/service/metric"
	"github.com/QOSGroup/qmoon/types"
	"github.com/hashicorp/go-version"
	"github.com/tendermint/go-amino"
)

type Route struct {
	Route  string `json:"route"`
	Hidden bool   `json:"hidden"`
}

var defaultRoute = []Route{
	{
		Route:  "ResultBlockBase",
		Hidden: false,
	},
	{
		Route:  "Validtor",
		Hidden: false,
	},
	{
		Route:  "Node",
		Hidden: false,
	},
}

type Node struct {
	Name        string          `json:"name"`    // name
	BaseURL     string          `json:"baseUrl"` // base_url
	SecretKey   string          `json:"-"`       // secret_key
	ChanID      string          `json:"chanId"`
	NodeType    string          `json:"nodeType"`
	NodeVersion version.Version `json:"nodeVersion"`
	Routers     []Route         `json:"routers"`
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

func covertToNode(mnt *models.Node) *Node {
	nv, _ := version.NewVersion(mnt.NodeVersion)
	return &Node{
		Name:        mnt.Name,
		BaseURL:     mnt.BaseUrl,
		SecretKey:   mnt.SecretKey,
		ChanID:      mnt.ChainId,
		NodeType:    mnt.NodeType,
		NodeVersion: *nv,
		Routers:     defaultRoute,
	}
}

func AllNodes() ([]*Node, error) {
	ns, err := models.Nodes()

	if err != nil {
		return nil, err
	}

	var res []*Node
	for _, v := range ns {
		res = append(res, covertToNode(v))
	}

	return res, nil
}

func CreateNode(name, baseURL, nodeType, nodeVersion, secretKey string) error {
	if strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL[:len(baseURL)-1]
	}

	tmc := lib.TendermintClient(baseURL)
	status, err := tmc.Status()
	if err != nil {
		return fmt.Errorf("retrieve node netinfo err:%s", err)
	}

	chainID := status.NodeInfo.Network
	if _, err := models.CreateNode(name, baseURL, nodeType, nodeVersion, secretKey, chainID); err != nil {
		return err
	}

	if err := metric.CreateDatabase(chainID); err != nil {
		return errors.New("CreateDatabase:" + err.Error())
	}

	return nil
}

func GetNodeByName(name string) (*Node, error) {
	mn, err := models.RetrieveNodeByName(name)
	if err != nil {
		return nil, err
	}

	return covertToNode(mn), nil
}

func DeleteNodeByName(name string) error {
	if err := models.DropDatabase(name); err != nil {
		return err
	}

	if err := metric.DeleteDatabase(name); err != nil {
		return err
	}

	return models.DeleteNodeByName(name)
}

func (n Node) Bech32PrefixConsPub() string {
	switch types.NodeType(n.NodeType) {
	case types.NodeTypeQOS:
		return ""
	case types.NodeTypeQSC:
		return ""
	case types.NodeTypeCOSMOS:
		return "cosmosvalconspub"
	default:
		return ""
	}
}
