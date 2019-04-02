package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var c *gocache.Cache

const cacheDuration = time.Second * 60
const cacheDurationLong = time.Minute * 5

func init() {
	c = gocache.New(3*time.Minute, 10*time.Minute)
}

func Set(k string, x interface{}, d time.Duration) {
	c.Set(k, x, d)
}

func Get(k string) (interface{}, bool) {
	return c.Get(k)
}
