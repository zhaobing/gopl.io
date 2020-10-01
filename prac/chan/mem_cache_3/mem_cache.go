package mem_cache_3

//缓存函数的执行结果

type result struct {
	value interface{}
	err   error
}

type entry struct {
	result        result
	completeBroad chan struct{}
}

type request struct {
	key      string
	response chan result
}

type MemCache struct {
	requests chan request
}

type Fun func(string) (interface{}, error)

func New(fun Fun) *MemCache {
	mem := &MemCache{
		requests: make(chan request),
	}
	go mem.server(fun)
	return mem
}

func (m *MemCache) Get(key string) (interface{}, error) {
	response := make(chan result)
	req := request{
		key:      key,
		response: response,
	}
	m.requests <- req
	res := <-req.response
	return res.value, res.err
}

func (m *MemCache) server(fun Fun) {
	cache := make(map[string]*entry)
	for request := range m.requests {
		e, ok := cache[request.key]
		if !ok {
			e = &entry{completeBroad: make(chan struct{})}
			cache[request.key] = e
			go e.call(fun, request.key)
		}
		go e.deliver(request.response)
	}
}

func (e *entry) call(fun Fun, key string) {
	e.result.value, e.result.err = fun(key)
	close(e.completeBroad)
}

func (e *entry) deliver(response chan<- result) {
	<-e.completeBroad
	response <- e.result
}
