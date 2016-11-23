package LRU

import "testing"

func Test_NewLRUCache(t *testing.T){
	var size int = 1
	cache := NewLRUCache(size)
	if cache == nil {
		t.Errorf("error")
	}
}

func ExpectValue(t *testing.T,cache *LRU,key uint,value interface{})bool{
	v,ok := cache.Get(key)
	if !ok {
		t.Errorf("can't find value")
		return false
	} else if value != v {
		t.Errorf("not equal")
		return false
	}
	return true
}
func Test_Add(t *testing.T){
	cache := NewLRUCache(10)
	cache.Add(1,10)
	ExpectValue(t,cache,1,10)
}

func Test_Override(t *testing.T){
	cache := NewLRUCache(10)
	cache.Add(1,10)
	ExpectValue(t,cache,1,10)
	cache.Add(1,11)
	ExpectValue(t,cache,1,11)
	cache.Add(1,12)
	ExpectValue(t,cache,1,12)
}

func Test_Overflow(t *testing.T){
	cache := NewLRUCache(shardsCount)
	for i := 0 ; i < shardsCount + 1 ; i ++ {
		cache.Add(uint(i),"fuck")
	}
	exp := make([]uint,0)
	for i := 0 ; i < shardsCount+1 ; i ++ {
		_,ok := cache.Get(uint(i))
		if ok {
			exp = append(exp,uint(i))
		}
	}
	if len(exp) != shardsCount {
		t.Errorf("fuck")
	}
}
