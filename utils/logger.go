// Copyright 2018 The QOS Authors

package utils

import (
	"fmt"
	"io"
	"strings"
)

func SQLLog(w io.Writer) func(d string, args ...interface{}) {
	return func(d string, args ...interface{}) {
		for k, v := range args {
			d = strings.Replace(d, fmt.Sprintf("$%d", k+1), fmt.Sprintf("%v", v), -1)
		}

		fmt.Println()
		fmt.Fprint(w, d+"\n")
	}
}
