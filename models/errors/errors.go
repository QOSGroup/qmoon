// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"
)

var InternalServerError = errors.New("internal server error")

// New is a wrapper of real errors.New function.
func New(text string) error {
	return errors.New(text)
}

type NotExist struct {
	Obj string
}

func IsNotExist(err error) bool {
	_, ok := err.(NotExist)
	return ok
}

func (err NotExist) Error() string {
	return fmt.Sprintf("%s does not exist", err.Obj)
}
