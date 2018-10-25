// Copyright 2018 The QOS Authors

package types

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
