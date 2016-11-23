package LRU

import (
	"testing"
	"time"
)

func unexpectValue(t *testing.T,cache *ExpireCache,key Key) bool {
	if cache == nil  {
		t.Errorf("bad expirecache")
		return false
	}
	v := cache.Get(key)
	if v == nil  {
		return true
	} else {
		return false
	}
}
func expectValue(t *testing.T,cache *ExpireCache,key Key,value interface{})bool{
	if cache == nil {
		t.Errorf("bad expirecache")
		return false
	}
	v := cache.Get(key)
	if v == nil || v != value{
		t.Errorf("can't find value")
		return false
	}
	return false
}

func TestExpireCache_Add(t *testing.T) {
	cache := NewExpireCache(10)
	cache.Add(1,1,time.Duration(10))
	expectValue(t,cache,1,1)
}

func Test_SimpleGet1(t *testing.T){
	cache := NewExpireCache(10)
	cache.Add(1,1,time.Duration(0))
	res := unexpectValue(t,cache,1)
	if res {
		t.Errorf("")
	}
}

func Test_SimpleGet2(t *testing.T){
	cache := NewExpireCache(2)
	cache.Add(1,1,time.Duration(100))
	expectValue(t,cache,1,1)
}
func Test_overflow(t *testing.T){
	cache := NewExpireCache(4)
	cache.Add(1,1,time.Duration(100))
	cache.Add(2,2,time.Duration(100))
	cache.Add(3,3,time.Duration(100))
	cache.Add(4,4,time.Duration(100))
	cache.Add(5,5,time.Duration(100))

	res:=unexpectValue(t,cache,1)
	if !res {
		t.Errorf("")
	}
	expectValue(t,cache,2,2)
	expectValue(t,cache,3,3)
	expectValue(t,cache,4,4)
	expectValue(t,cache,5,5)
}
