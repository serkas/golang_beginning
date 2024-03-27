package memorycache

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

type Cache struct {
	data map[string]string
	mx   sync.Mutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key string, value string) {
	c.mx.Lock()
	c.data[key] = value
	c.mx.Unlock()
}

func (c *Cache) Get(key string) (string, error) {
	c.mx.Lock()
	val, exist := c.data[key]
	c.mx.Unlock()

	if !exist {
		return "", ErrNotFound
	}

	return val, nil
}
