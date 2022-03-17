package main

import (
	"log"
	"sync"
)

type Counter struct {
	v  int
	wg *sync.WaitGroup
}

func (c *Counter) Increment() {
	c.v++
	c.wg.Done()
}

func (c *Counter) Value() int {
	return c.v
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{wg: &wg}
	for i := 0; i < 1000; i++ {
		c.wg.Add(1)
		go c.Increment()
	}

	c.wg.Wait()
	log.Println(c.Value())
}
