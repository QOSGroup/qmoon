// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errors

import "fmt"

type NodeNotExist struct {
	Name string
}

func IsNodeNotExist(err error) bool {
	_, ok := err.(NodeNotExist)
	return ok
}

func (err NodeNotExist) Error() string {
	return fmt.Sprintf("node does not exist [name: %s]", err.Name)
}
