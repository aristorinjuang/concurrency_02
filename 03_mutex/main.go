package main

import (
	"log"
	"sync"
	"time"
)

type Counter struct {
	m sync.Mutex
	v int
}

func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()

	c.v++
}

func (c *Counter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()

	return c.v
}

func main() {
	c := &Counter{}
	for i := 0; i < 1000; i++ {
		go c.Increment()
	}

	time.Sleep(time.Second)
	log.Println(c.Value())
}
