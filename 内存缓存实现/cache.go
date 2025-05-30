package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	value  interface{}
	expire int64
}

type Cache struct {
	cacheItems map[string]*CacheItem
	m          sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		cacheItems: make(map[string]*CacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}, expire int64) {
	c.m.Lock()
	defer c.m.Unlock()

	c.cacheItems[key] = &CacheItem{
		value:  value,
		expire: expire,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	if value, ok := c.cacheItems[key]; ok {
		if value.expire == 0 || value.expire > time.Now().UnixNano() {
			return value.value, true
		} else {
			delete(c.cacheItems, key)
			return nil, false
		}
	}
	return nil, false
}

func main() {
	cache := NewCache()
	cache.Set("key1", "value1", time.Now().Add(time.Second).UnixNano())
	cache.Set("key2", "value2", 0)

	if value, ok := cache.Get("key1"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("key 已过期或不存在")
	}
}
