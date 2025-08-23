// Package cache Реализация кэша на min-heap
package cache

import (
	"container/heap"
)

type Cachable interface {
	Priority() int64 // учитывая величину приоритетов, можно запрашивать int64
	Key() string
	Value() interface{}
	UpdatePriority()
}

type cachableItem struct {
	item  Cachable
	index int
}

// имплементация очереди с приоритетом для Cachable, реализация интерфейса heap.Interface
type cachablePQ []*cachableItem

func (pq *cachablePQ) Len() int {
	return len(*pq)
}

func (pq *cachablePQ) Less(i, j int) bool {
	return (*pq)[i].item.Priority() > (*pq)[j].item.Priority() // куча должна распределяться от меньшего к большему
}

func (pq *cachablePQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *cachablePQ) Push(x interface{}) {
	n := len(*pq)
	item := &cachableItem{item: x.(Cachable), index: n}
	*pq = append(*pq, item)
}

func (pq *cachablePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Cache реализация кэша с приоритетом. За приоритет и содержимое должен отвечать сам Cachable объект
type Cache struct {
	capacity int
	items    map[string]*cachableItem
	pq       heap.Interface // используем min-heap для быстрого обновления кэша при перезаполнении
}

func NewCache(cap int) *Cache {
	pq := &cachablePQ{}
	heap.Init(pq)
	return &Cache{
		capacity: cap,
		items:    make(map[string]*cachableItem),
		pq:       pq,
	}
}

func (c *Cache) Get(key string) interface{} {
	item, ok := c.items[key]
	if !ok {
		return nil
	}
	item.item.UpdatePriority()
	heap.Fix(c.pq, item.index)
	return item.item.Value()
}

func (c *Cache) Put(key string, item Cachable) {
	if c.capacity == c.pq.Len() {
		popItem := c.pq.Pop().(*cachableItem)
		delete(c.items, popItem.item.Key())
	}
	cacheItem := &cachableItem{item: item}
	c.items[key] = cacheItem
	heap.Push(c.pq, item)
}
