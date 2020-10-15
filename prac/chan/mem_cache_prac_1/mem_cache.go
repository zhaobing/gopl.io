package mem_cache_prac_1

//缓存函数的执行结果
//使用entry对函数执行进行等待，completeChan用来广播执行结果
//内存缓存对客户端缓存请求使用requestChan进行异步的调用和结果返回
//对客户端的request请求进行封装，封装内部使用resultChan来进行异步通信
//MemCache的Server不断接受客户端请求
//异步执行函数
//通过resultChan来返回执行结果
//使用entry中的completeChain来管理多个函数的并行，使用广播通知的方式来告知等待获取直接结果的goroutine

type Func func(key string) (interface{}, error)

type excResult struct {
	value interface{}
	err   error
}

type entry struct {
	excResult excResult
	complete  chan struct{}
}

func (en *entry) exec(fun Func, key string) {
	en.excResult.value, en.excResult.err = fun(key)
	close(en.complete)
}

func (en *entry) deliver(respChan chan excResult) {
	<-en.complete
	respChan <- en.excResult
}

type request struct {
	key    string
	respCh chan excResult
}

type MemCache struct {
	requests chan request
}

func (c *MemCache) Get(key string) (interface{}, error) {
	r := request{
		key:    key,
		respCh: make(chan excResult),
	}
	c.requests <- r
	result := <-r.respCh
	return result.value, result.err
}

func (c *MemCache) server(fun Func) {
	cache := make(map[string]*entry)
	for req := range c.requests {
		en, ok := cache[req.key]
		if !ok {
			en = &entry{complete: make(chan struct{})}
			cache[req.key] = en
			go en.exec(fun, req.key)
		}
		go en.deliver(req.respCh)
	}
}

func New(fun Func) *MemCache {
	cache := &MemCache{requests: make(chan request)}
	go cache.server(fun)
	return cache
}
