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
	data sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, value interface{}, expire int64) {
	c.data.Store(key, &CacheItem{
		value:  value,
		expire: expire,
	})
}

func (c *Cache) Get(key string) (interface{}, bool) {
	itemRaw, ok := c.data.Load(key)
	if !ok {
		return nil, false
	}

	item := itemRaw.(*CacheItem)
	if item.expire == 0 || item.expire > time.Now().UnixNano() {
		return item.value, true
	}

	// 已过期，删除
	c.data.Delete(key)
	return nil, false
}

// 可选：启动后台定时清理过期项
func (c *Cache) StartGC(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			now := time.Now().UnixNano()
			c.data.Range(func(key, value interface{}) bool {
				item := value.(*CacheItem)
				if item.expire != 0 && item.expire <= now {
					c.data.Delete(key)
				}
				return true
			})
		}
	}()
}

func main() {
	cache := NewCache()
	cache.Set("key1", "value1", time.Now().Add(time.Second).UnixNano())
	cache.Set("key2", "value2", 0)

	// 启动自动清理协程
	cache.StartGC(time.Second)

	// 模拟等待
	time.Sleep(2 * time.Second)

	if value, ok := cache.Get("key1"); ok {
		fmt.Println("key:", value)
	} else {
		fmt.Println("key 已过期或不存在")
	}

	if value, ok := cache.Get("key2"); ok {
		fmt.Println("key:", value)
	}
}
