// Copyright 2018 The QOS Authors

package utils

import "math"

func Percent(molecular, denominator uint64) float64 {
	m := math.Float64frombits(molecular)
	d := math.Float64frombits(denominator)

	return m / d
}
