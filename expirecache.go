package LRU

import "time"
import (
	"sync"
)

type ExpireCache struct {
	cache *Cache
	lock sync.Mutex
}

func NewExpireCache(size int) *ExpireCache{
	return &ExpireCache{
		cache: NewCache(size),
	}
}
type TimeEntity struct{
	time time.Time
	value interface{}
}

func (cache *ExpireCache) Add(key Key,value interface{},ttl time.Duration){
	cache.lock.Lock()
	defer cache.lock.Unlock()
	var entity TimeEntity
	entity.time = time.Now().Add(ttl)
	entity.value = value
	cache.cache.AddItem(key,&entity)
	time.AfterFunc(ttl,func(){
		cache.remove(key)
	})
}

func (cache *ExpireCache) remove(key Key){
	cache.lock.Lock()
	defer cache.lock.Unlock()
	cache.cache.Remove(key)
}

func (cache *ExpireCache) Get(key Key) interface{}{
	cache.lock.Lock()
	defer cache.lock.Unlock()
	item,ok := cache.cache.GetElem(key)
	if !ok {
		return nil
	}
	if time.Now().After(item.(*TimeEntity).time) {
		go cache.remove(key)
		return nil
	}
	return item.(*TimeEntity).value
}
