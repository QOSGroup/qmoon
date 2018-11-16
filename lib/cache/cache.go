package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var c *gocache.Cache

const defaultExpiration = time.Second * 60
const cleanupInterval = time.Minute * 5

func init() {
	c = gocache.New(defaultExpiration, cleanupInterval)
}

func Set(k string, x interface{}, d time.Duration) {
	c.Set(k, x, d)
}

func Get(k string) (interface{}, bool) {
	return c.Get(k)
}
