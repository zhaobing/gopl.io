package mem_cache_2

import (
	"sync"
)

//缓存函数的执行结果
type result struct {
	value interface{}
	err   error
}

type entry struct {
	result result
	ready  chan struct{}
}

type MemCache struct {
	fun   Fun
	cache map[string]*entry
	mu    sync.Mutex
}

type Fun func(string) (interface{}, error)

func New(fun Fun) *MemCache {
	return &MemCache{
		fun:   fun,
		cache: make(map[string]*entry),
	}
}

func (mc *MemCache) Get(key string) (interface{}, error) {
	mc.mu.Lock()
	e, ok := mc.cache[key]
	if !ok {
		e = &entry{ready: make(chan struct{})}
		mc.cache[key] = e
		mc.mu.Unlock()
		e.result.value, e.result.err = mc.fun(key)
		close(e.ready)
	} else {
		mc.mu.Unlock()
		<-e.ready
	}
	return e.result.value, e.result.err
}
