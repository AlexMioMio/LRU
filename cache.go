package LRU

import (
	"sync"
)

const shardsCount = 32

type Cache []*cacheShard

func NewCache(size uint) Cache{
	if size < shardsCount {
		size = shardsCount
	}
	cache := make(Cache,shardsCount)
	for i := 0 ; i < size ; i++ {
		cache[i] = &cacheShard{
			size: size/shardsCount,
			mp: make(map[uint]interface{}),
		}
	}
	return cache
}

func (c Cache) getShard(index uint) *cacheShard{
	return c[index%(uint(shardsCount))]
}

func (c Cache) Get(index uint) (obj interface{},found bool){
	return c.getShard(index).get(index)
}
func (c Cache) Add(index uint,obj interface{})bool {
	return c.getShard(index).add(index,obj)
}
type cacheShard struct {
	mp map[uint]interface{}
	size int
	sync.RWMutex
}

func (c *cacheShard) add(index uint,value interface{}) bool{
	c.Lock()
	defer c.Unlock()
	_, isOverride := c.mp[index]
	if ! isOverride && len(c.mp) >= c.size {
		var random uint
		for random = range c.mp {
			break
		}
		delete(c.mp,random)
	}
	c.mp[index] = value
	return isOverride
}

func (c *cacheShard) get(index uint) (obj interface{},found bool){
	c.RLock()
	defer c.RUnlock()
	obj,ok := c.mp[index]
	return obj,ok
}

