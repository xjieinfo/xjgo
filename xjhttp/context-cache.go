package xjhttp

import "github.com/xjieinfo/xjgo/xjcache"

//未写完, 暂时不使用

func (c *Context) CacheNew(cacheName string, maxLifeTime int64) {
	c.CacheMgr = xjcache.NewCacheMgr(cacheName, maxLifeTime)
}

func (c *Context) CacheSet(id, key string, value interface{}) error {

	return c.CacheMgr.Set(id, key, value)
}
