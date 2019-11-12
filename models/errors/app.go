// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errors

import "fmt"

type AppNotExist struct {
	ID        int64
	SecretKey string
}

func IsAppNotExist(err error) bool {
	_, ok := err.(AppNotExist)
	return ok
}

func (err AppNotExist) Error() string {
	return fmt.Sprintf("app does not exist [id:%d, secretKey: %s]", err.ID, err.SecretKey)
}
