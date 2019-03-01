// Copyright 2018 The QOS Authors

package types

type NodeType string

func (nt NodeType) String() string {
	return string(nt)
}

const (
	NodeTypeQOS    NodeType = "QOS"
	NodeTypeQSC    NodeType = "QSC"
	NodeTypeCOSMOS NodeType = "COSMOS"
)

func CheckNodeType(nodeType NodeType) bool {
	return nodeType == NodeTypeQOS || nodeType == NodeTypeQSC || nodeType == NodeTypeCOSMOS
}

type NodeTypeRouter struct {
	Route  string `json:"route"`
	Hidden bool   `json:"hidden"`
}

type AdminNodeType struct {
	Name    string           `json:"name"`
	BaseURL string           `json:"baseUrl"`
	Routers []NodeTypeRouter `json:"routers"`
}

type ResultNodeTypes struct {
	Nodes []AdminNodeType `json:"nodes"`
}

type ExplorerRoute string

const (
	ExplorerRouteBlock    ExplorerRoute = "ResultBlockBase"
	ExplorerRouteValidtor ExplorerRoute = "Validtor"
	ExplorerRouteNode     ExplorerRoute = "Node"
)

func (er ExplorerRoute) String() string {
	return string(er)
}
