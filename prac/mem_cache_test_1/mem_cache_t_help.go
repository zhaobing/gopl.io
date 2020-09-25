package mem_cache_test_1

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

type M interface {
	Get(key string) (interface{}, error)
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

var urls = []string{
	"http://www.gusuwang.com",
	"https://www.liansuo.com/",
	"https://www.iyiou.com/",
	"http://www.sina.com",
	"http://www.sohu.com",
	"http://www.17ok.com",
	"http://www.6.cn",
	"http://www.eastday.com",
	"http://www.efu.com.cn/",
	"http://www.gusuwang.com",
	"https://www.liansuo.com/",
	"https://www.iyiou.com/",
	"http://www.sina.com",
	"http://www.sohu.com",
	"http://www.17ok.com",
	"http://www.6.cn",
	"http://www.eastday.com",
	"http://www.efu.com.cn/",
	"http://www.eastday.com",
	"http://www.efu.com.cn/",
	"http://www.gusuwang.com",
	"https://www.liansuo.com/",
	"https://www.iyiou.com/",
	"http://www.sina.com",
	"http://www.sohu.com",
	"http://www.17ok.com",
	"http://www.6.cn",
	"http://www.eastday.com",
	"http://www.efu.com.cn/",
}

func incomingUrls() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func Sequential(t *testing.T, m M) {
	for url := range incomingUrls() {
		start := time.Now()
		val, err := m.Get(url)
		if err != nil {
			logs.Error(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(val.([]byte)))
	}
}

func Concurrent(t *testing.T, m M) {
	var wg sync.WaitGroup
	i := 0
	for url := range incomingUrls() {
		wg.Add(1)
		i++
		go func(href string, idx int) {
			defer wg.Done()
			start := time.Now()
			val, err := m.Get(href)
			if err != nil {
				logs.Error(err)
				return
			}
			fmt.Printf("%d,%s, %s, %d bytes\n", idx, href, time.Since(start), len(val.([]byte)))
		}(url, i)
	}

	wg.Wait()

}
