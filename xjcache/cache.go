package xjcache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cacheID  string
	lastTime time.Time
	values   map[string]interface{}
}

type CacheMgr struct {
	cacheName   string
	mLock       sync.RWMutex
	maxLifeTime int64
	caches      map[string]*Cache
}

func NewCacheMgr(cacheName string, maxLifeTime int64) *CacheMgr {
	mgr := &CacheMgr{cacheName: cacheName, maxLifeTime: maxLifeTime, caches: make(map[string]*Cache)}
	go mgr.CacheGC()
	return mgr
}

func (mgr *CacheMgr) CacheGC() {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	for id, cache := range mgr.caches {
		if cache.lastTime.Unix()+mgr.maxLifeTime < time.Now().Unix() {
			delete(mgr.caches, id)
		}
	}
	time.AfterFunc(time.Duration(mgr.maxLifeTime)*time.Second, func() {
		mgr.CacheGC()
	})
}

func (mgr *CacheMgr) Set(id, key string, value interface{}) error {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	if session, ok := mgr.caches[id]; ok {
		session.values[key] = value
		return nil
	} else {
		values := make(map[string]interface{})
		values[key] = value
		cache := &Cache{
			cacheID:  id,
			lastTime: time.Now().Add(time.Second * time.Duration(mgr.maxLifeTime)),
			values:   values,
		}
		mgr.caches[id] = cache
	}
	return errors.New("invalid cache ID")
}

// GetSessionValue get value fo session
func (mgr *CacheMgr) Get(id string, key string) interface{} {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()
	if session, ok := mgr.caches[id]; ok {
		if val, ok := session.values[key]; ok {
			return val
		}
	}
	return nil
}

func (mgr *CacheMgr) DelGroup(cacheGroup string) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	delete(mgr.caches, cacheGroup)
}

func (mgr *CacheMgr) Del(cacheGroup, key string) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	group := mgr.caches[cacheGroup]
	delete(group.values, key)
}
