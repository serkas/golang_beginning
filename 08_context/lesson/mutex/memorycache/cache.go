package memorycache

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Cache struct {
	data map[string]string
}

func New() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key string, value string) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (string, error) {
	val, exist := c.data[key]
	if !exist {
		return "", ErrNotFound
	}

	return val, nil
}
