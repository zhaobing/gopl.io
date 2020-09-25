package mem_cache_1

import (
	"github.com/zhaobing/gopl.io/prac/mem_cache_test_1"
	"testing"
)

func Test(t *testing.T) {
	mCache := New(mem_cache_test_1.HTTPGetBody)
	mem_cache_test_1.Sequential(t, mCache)
}

func TestConcurrent(t *testing.T) {
	mCache := New(mem_cache_test_1.HTTPGetBody)
	mem_cache_test_1.Concurrent(t, mCache)
}
