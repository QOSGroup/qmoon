// Copyright 2018 The QOS Authors

package utils

import (
	"fmt"
	"time"
)

func DayStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// NDaysAgo N天前
func NDaysAgo(t time.Time, n int) time.Time {
	s := fmt.Sprintf("-%dh", 24*n)
	d, _ := time.ParseDuration(s)

	res := t.Add(d)

	return res
}
