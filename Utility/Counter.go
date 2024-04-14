package Utility

import "sync/atomic"

type Counter struct {
	Counters map[string]*atomic.Int64
}

type ICounter interface {
	Born(key string)
	Increment(key string)
	Get(key string) int64
	Add(key string, value int64)
}

func NewCounter() ICounter {
	return &Counter{
		Counters: make(map[string]*atomic.Int64),
	}
}

func (c *Counter) Born(key string) {
	c.Counters[key] = new(atomic.Int64)
}

func (c *Counter) Increment(key string) {
	c.Counters[key].Add(1)
}

func (c *Counter) Add(key string, value int64) {
	c.Counters[key].Add(value)
}

func (c *Counter) Get(key string) int64 {
	return c.Counters[key].Load()
}
