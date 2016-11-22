package LRU

import "container/list"
type Key interface {}

type Cache struct{
	Evict func(key Key,value interface{})
	list * list.List
	cache map[Key]*list.Element
	MaxSize int
}

func NewCache(size int) *Cache {
	return &Cache{
		list: list.New(),
		cache: make(map[Key]*list.Element),
		MaxSize:size,
	}
}

type Entry struct {
	key Key
	value interface{}
}

func (c *Cache) Add(key Key,value interface{}){
	if c.cache == nil {
		c.list = list.New()
		c.cache = make(map[Key]*list.Element)
	}
	if elem,ok := c.cache[key] ; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*Entry).value = value
		return
	}
	elem := c.list.PushFront(&Entry{key,value})
	c.cache[key] = elem
	if c.MaxSize != 0 && c.list.Len() > c.MaxSize {
		c.RemoveOldest()
	}

}

func (c *Cache) Remove(key Key){
	if c.cache == nil {
		return
	}
	if elem,ok := c.cache[key];ok {
		c.RemoveElem(elem)
		delete(c.cache,key)
	}
}
func (c *Cache) RemoveElem(elem *list.Element){
	if elem == nil {
		return
	}
	if c.cache == nil {
		return
	}
	c.list.Remove(elem)
	key := elem.Value.(*Entry).key
	delete(c.cache,key)
}

func (c *Cache) RemoveOldest(){
	if c.cache == nil {
		return
	}
	ele := c.list.Back()
	c.RemoveElem(ele)
}

func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.list.Len()
}
