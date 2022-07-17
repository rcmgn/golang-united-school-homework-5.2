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
	if v, ok := c.time[key]; ok {
		if v.Before(time.Now()) {
			delete(c.time, key)
			delete(c.data, key)
			return "", false
		}
	}
	result, isExist = c.data[key]
	return
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
}

func (c *Cache) Keys() (result []string) {
	for key, _ := range c.data {
		if v, ok := c.time[key]; ok {
			if v.Before(time.Now()) {
				delete(c.time, key)
				delete(c.data, key)
				continue
			}
		}
		result = append(result, key)
	}
	return
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.time[key] = deadline
}
