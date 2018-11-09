// Copyright 2018 The QOS Authors

package hadmin

import (
	"net/http"
	"testing"

	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateNodeGin(t *testing.T) {
	body := createNodeQuery{
		Name:      "QOS",
		BaseURL:   "http://127.0.0.1:26657",
		SecretKey: "",
		Routers: []service.NodeRoute{
			{
				Route:  "Block",
				Hidden: false,
			},
			{
				Route:  "Node",
				Hidden: false,
			},
			{
				Route:  "Validtor",
				Hidden: false,
			},
		},
	}
	req, err := utils.NewPostJsonRequest(nodeTypeUrl, body)
	assert.Nil(t, err)

	_, err = handler.NewHttpTest(t, req).WithLocalIP().Do(NodesGinRegister, nil)
	assert.NotNil(t, err)
}

func TestListNodesGin(t *testing.T) {
	body := createNodeQuery{
		Name:      "QOS",
		BaseURL:   "http://127.0.0.1:26657",
		SecretKey: "",
		Routers: []service.NodeRoute{
			{
				Route:  "Block",
				Hidden: false,
			},
			{
				Route:  "Node",
				Hidden: false,
			},
			{
				Route:  "Validtor",
				Hidden: false,
			},
		},
	}
	req, err := utils.NewPostJsonRequest(nodeTypeUrl, body)
	assert.Nil(t, err)

	_, err = handler.NewHttpTest(t, req).Do(NodesGinRegister, nil)
	assert.NotNil(t, err)
	req, err = http.NewRequest(http.MethodGet, accountsUrl, nil)
	assert.Nil(t, err)
	var nts []*service.Node
	_, err = handler.NewHttpTest(t, req).Do(NodesGinRegister, &nts)
	assert.Nil(t, err)
	assert.NotEmpty(t, nts)
}
