package mem_cache_4

//缓存函数的执行结果
//使用entry对函数执行进行等待，completeChan用来广播执行结果
//内存缓存对客户端缓存请求使用requestChan进行异步的调用和结果返回
//对客户端的request请求进行封装，封装内部使用resultChan来进行异步通信
//MemCache的Server不断接受客户端请求
//异步执行函数
//通过resultChan来返回执行结果
//使用entry中的completeChain来管理多个函数的并行，使用广播通知的方式来告知等待获取直接结果的goroutine

//要执行的函数
type Fun func(key string) (interface{}, error)

//函数执行结果
type result struct {
	value interface{}
	err   error
}

//执行结果的封装
type entry struct {
	excResult result
	complete  chan struct{}
}

func (e *entry) exec(key string, fun Fun) {
	e.excResult.value, e.excResult.err = fun(key)
	close(e.complete)
}

func (e *entry) deliver(respCh chan<- result) {
	<-e.complete
	respCh <- e.excResult
}

//缓存请求，封装了客户端对缓存的请求信息和响应通道
type req struct {
	key    string
	respCh chan result
}

//内存缓存
type MemCache struct {
	reqCh chan req
}

//初始化内存缓存
func New(fun Fun) *MemCache {
	m := &MemCache{reqCh: make(chan req)}
	go m.server(fun)
	return m
}

func (m *MemCache) Get(key string) (interface{}, error) {
	r := req{
		key:    key,
		respCh: make(chan result),
	}
	m.reqCh <- r
	result := <-r.respCh
	return result.value, result.err
}

func (m *MemCache) server(fun Fun) {
	cache := make(map[string]*entry)
	for req := range m.reqCh {
		e, ok := cache[req.key]
		if !ok {
			e = &entry{complete: make(chan struct{})}
			cache[req.key] = e
			go e.exec(req.key, fun)
		}
		go e.deliver(req.respCh)
	}
}
