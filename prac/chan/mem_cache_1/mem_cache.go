package mem_cache_1

import (
	"sync"
)

//缓存函数的执行结果

type result struct {
	value interface{}
	err   error
}

type MemCache struct {
	fun   Fun
	cache map[string]result
	mu    sync.Mutex
}

type Fun func(string) (interface{}, error)

func New(function Fun) *MemCache {
	return &MemCache{
		fun:   function,
		cache: make(map[string]result),
	}
}

func (mc *MemCache) Get(key string) (interface{}, error) {
	mc.mu.Lock()
	result, ok := mc.cache[key]
	mc.mu.Unlock()
	if !ok {
		result.value, result.err = mc.fun(key)
		mc.mu.Lock()
		mc.cache[key] = result
		mc.mu.Unlock()
	}
	return result.value, result.err
}
