package xjhttp

import (
	"github.com/xjieinfo/xjgo/xjcache"
	"github.com/xjieinfo/xjgo/xjcore/xjconv"
)

func (c *Context) GetCacheInt(AccessSecret string, Caches *xjcache.CacheMgr, key string) int {
	val := c.GetCache(AccessSecret, Caches, key)
	return xjconv.InterfaceToInt(val)
}

func (c *Context) GetCacheInt64(AccessSecret string, Caches *xjcache.CacheMgr, key string) int64 {
	val := c.GetCache(AccessSecret, Caches, key)
	return xjconv.InterfaceToInt64(val)
}

func (c *Context) GetCacheString(AccessSecret string, Caches *xjcache.CacheMgr, key string) string {
	val := c.GetCache(AccessSecret, Caches, key)
	return xjconv.InterfaceToString(val)
}

func (c *Context) GetCache(AccessSecret string, Caches *xjcache.CacheMgr, key string) interface{} {
	user, ext, _ := GetAllFromBearerToken(c, AccessSecret)
	if ext != "" {
		user += ":" + ext
	}
	val := Caches.Get(user, key)
	return val
}
