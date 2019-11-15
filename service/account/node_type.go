// Copyright 2018 The QOS Authors

package account

import "github.com/QOSGroup/qmoon/types"

type adminNodeTypeOption struct {
}

type setAdminNodeTypeOption func(*adminNodeTypeOption) error

func CreateAdminNodeType(opt *adminNodeTypeOption) error {
	return nil
}

func AdminNodeTypes() (*types.ResultNodeTypes, error) {
	return nil, nil
}
