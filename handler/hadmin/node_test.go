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

func TestCreateNodeTypeGin(t *testing.T) {
	body := createNodeTypeQuery{
		Name:      "QOS",
		BaseURL:   "http://127.0.0.1:26657",
		SecretKey: "",
		Routers: []service.NodeTypeRoute{
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

	_, err = handler.NewHttpTest(t, req).WithLocalIP().Do(NodeTypesGinRegister, nil)
	assert.NotNil(t, err)
}

func TestListNodeTypesGin(t *testing.T) {
	body := createNodeTypeQuery{
		Name:      "QOS",
		BaseURL:   "http://127.0.0.1:26657",
		SecretKey: "",
		Routers: []service.NodeTypeRoute{
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

	_, err = handler.NewHttpTest(t, req).Do(NodeTypesGinRegister, nil)
	assert.NotNil(t, err)
	req, err = http.NewRequest(http.MethodGet, accountsUrl, nil)
	assert.Nil(t, err)
	var nts []*service.NodeType
	_, err = handler.NewHttpTest(t, req).Do(NodeTypesGinRegister, &nts)
	assert.Nil(t, err)
	assert.NotEmpty(t, nts)
}
