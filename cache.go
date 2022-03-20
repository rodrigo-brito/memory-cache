package cache

import "time"

type Item[T any] struct {
	TTL   time.Time
	Value T
}

type Cache[T any] struct {
	data          map[string]Item[T]
	clearInterval time.Duration
}

type Option[T any] func(c *Cache[T])

func New[T any](interval time.Duration) *Cache[T] {
	cache := &Cache[T]{
		clearInterval: interval,
		data:          make(map[string]Item[T]),
	}

	if cache.clearInterval > 0 {
		go cache.cleaner()
	}

	return cache
}

func (c *Cache[T]) cleaner() {
	for range time.NewTicker(c.clearInterval).C {
		for key, item := range c.data {
			if item.TTL.Before(time.Now()) {
				delete(c.data, key)
			}
		}
	}
}

func (c *Cache[T]) Get(key string) (value T, ok bool) {
	item, ok := c.data[key]
	if !ok {
		return
	}

	if item.TTL.Before(time.Now()) {
		delete(c.data, key)
		return value, false
	}

	return item.Value, true
}

func (c *Cache[T]) GetMulti(keys ...string) ([]T, bool) {
	result := make([]T, 0)
	for _, key := range keys {
		item, ok := c.data[key]
		if ok {
			if item.TTL.Before(time.Now()) {
				delete(c.data, key)
				continue
			}
			result = append(result, item.Value)
		}
	}
	return result, true
}

func (c *Cache[T]) Set(key string, value T, ttl time.Duration) {
	c.data[key] = Item[T]{
		TTL:   time.Now().Add(ttl),
		Value: value,
	}
}
