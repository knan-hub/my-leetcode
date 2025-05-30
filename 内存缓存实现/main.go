package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用 golang 实现一个内存缓存功能，需要满足任意类型存储、过期时间设置以及过期清除功能

// cacheItem 表示缓存中的一项
type cacheItem struct {
	value      any
	expiration int64 // Unix 时间戳，单位：纳秒
}

// 判断是否过期
func (item *cacheItem) IsExpired() bool {
	if item.expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.expiration
}

// Cache 是一个线程安全的内存缓存结构
type Cache struct {
	items map[string]*cacheItem
	lock  sync.RWMutex
	ttl   time.Duration
	stop  chan struct{}
}

// New 创建一个新的 Cache 实例，ttl 是默认过期时间，cleanupInterval 是清理频率
func New(ttl, cleanupInterval time.Duration) *Cache {
	c := &Cache{
		items: make(map[string]*cacheItem),
		ttl:   ttl,
		stop:  make(chan struct{}),
	}
	go c.cleanupLoop(cleanupInterval)
	return c
}

// Set 设置缓存键值对，可指定过期时间（为0则使用默认TTL）
func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	var exp int64
	if ttl <= 0 {
		ttl = c.ttl
	}
	if ttl > 0 {
		exp = time.Now().Add(ttl).UnixNano()
	}

	c.items[key] = &cacheItem{
		value:      value,
		expiration: exp,
	}
}

// Get 获取缓存项，返回值和是否存在标记
func (c *Cache) Get(key string) (any, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}
	if item.IsExpired() {
		return nil, false
	}
	return item.value, true
}

// Delete 删除缓存项
func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, key)
}

// Stop 停止后台清理任务
func (c *Cache) Stop() {
	close(c.stop)
}

// cleanupLoop 定期清理过期缓存项
func (c *Cache) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.stop:
			return
		}
	}
}

// cleanup 执行过期项的实际清理逻辑
func (c *Cache) cleanup() {
	c.lock.Lock()
	defer c.lock.Unlock()
	now := time.Now().UnixNano()
	for key, item := range c.items {
		if item.expiration > 0 && now > item.expiration {
			delete(c.items, key)
		}
	}
}

func main() {
	// 创建一个缓存，默认过期时间 5s，清理间隔 2s
	c := New(5*time.Second, 2*time.Second)
	defer c.Stop()

	// 设置一个键值对，3秒后过期
	c.Set("name", "xx", 3*time.Second)

	// 获取值
	if val, ok := c.Get("name"); ok {
		fmt.Println("获取 name:", val)
	} else {
		fmt.Println("未找到 name")
	}

	// 等待 4 秒让它过期
	time.Sleep(4 * time.Second)

	// 再次尝试获取
	if val, ok := c.Get("name"); ok {
		fmt.Println("获取 name（仍有效）:", val)
	} else {
		fmt.Println("name 已过期")
	}
}
