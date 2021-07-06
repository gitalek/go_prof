package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if listItem, ok := c.items[key]; ok {
		queueItem := cacheItem{
			key:   string(key),
			value: value,
		}
		listItem.Value = queueItem
		c.queue.MoveToFront(listItem)
		c.items[key] = listItem
		return true
	}

	if len(c.items) >= c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		queueItem, ok := last.Value.(cacheItem)
		if !ok {
			// An exceptional situation, so panic.
			panic("non cacheItem value in list item")
		}
		delete(c.items, Key(queueItem.key))
	}

	queueItem := cacheItem{
		key:   string(key),
		value: value,
	}
	listItem := c.queue.PushFront(queueItem)
	c.items[key] = listItem

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if listItem, ok := c.items[key]; ok {
		c.queue.MoveToFront(listItem)
		queueItem, ok := listItem.Value.(cacheItem)
		if !ok {
			// An exceptional situation, so panic.
			panic("non cacheItem value in list item")
		}
		return queueItem.value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
