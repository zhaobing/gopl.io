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

func New(function Fun) *MemCache {
	return &MemCache{
		fun:   function,
		cache: make(map[string]*entry),
	}
}

func (mc *MemCache) Get(key string) (interface{}, error) {
	mc.mu.Lock()
	e := mc.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		mc.cache[key] = e
		mc.mu.Unlock()
		e.result.value, e.result.err = mc.fun(key)
		close(e.ready) //广播执行就绪
	} else {
		mc.mu.Unlock()
		<-e.ready //接收执行完成的广播,阻塞等待正在执行的对应key的结果
	}

	return e.result.value, e.result.err
}
