// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errors

import "fmt"

type AccountNotExist struct {
	Id   int64
	Mail string
}

func IsAccountNotExist(err error) bool {
	_, ok := err.(AccountNotExist)
	return ok
}

func (err AccountNotExist) Error() string {
	return fmt.Sprintf("account does not exist [id:%d, mail: %s]", err.Id, err.Mail)
}
