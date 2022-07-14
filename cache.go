package cache

import "time"

type Cache struct {
	data map[string]string
	time map[string]time.Time
}

func NewCache() Cache {
	data := make(map[string]string)
	time := make(map[string]time.Time)
	return Cache{data, time}
}

func (c *Cache) Get(key string) (result string, isExist bool) {
	result, isExist = c.data[key]
	return
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
}

func (c *Cache) Keys() (result []string) {
	for k, _ := range c.data {
		result = append(result, k)
	}
	return
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.time[key] = deadline
}
