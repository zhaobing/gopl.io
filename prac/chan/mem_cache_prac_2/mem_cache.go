package mem_cache_prac_2

import "sync"

//缓存函数的执行结果
//entry 包装key对应的func执行结果，同时包装了chan来进行实现通知机制
//使用entry来避免多个goroutine同时运行相同的方法
//容器结构为map[key]*entry,同时增加一个mutex来保证map访问的并发安全
//get请求时，先加锁，防止线程并发，然后查询map是否有已经标记在执行的key的对应entry
//如果请求key已经标记为在执行（或者已经执行完成）,则解锁，并且对entry的complete进行等待，等待后获取到执行结果
//如果请求key没有被标记为在执行，则生成这个key对应的entry并且与key绑定，然后解锁，再执行具体方法，具体方法执行完成后使用该entry的complete进行通知

type Func func(key string) (interface{}, error)

type excResult struct {
	value interface{}
	err   error
}

type entry struct {
	excResult excResult
	complete  chan struct{}
}

type MemCache struct {
	fun       Func
	container map[string]*entry
	mu        sync.Mutex
}

func (c *MemCache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	e, ok := c.container[key]
	if !ok {
		e = &entry{
			complete: make(chan struct{}),
		}
		c.container[key] = e
		c.mu.Unlock()
		e.excResult.value, e.excResult.err = c.fun(key)
		close(e.complete)
	} else {
		c.mu.Unlock()
		<-e.complete
	}
	return e.excResult.value, e.excResult.err
}

func New(fun Func) *MemCache {
	cache := &MemCache{
		fun:       fun,
		container: make(map[string]*entry),
	}
	return cache
}
