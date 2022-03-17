package main

import "log"

type Counter struct {
	v int
}

func (c *Counter) Increment() {
	c.v++
}

func (c *Counter) Value() int {
	return c.v
}

func main() {
	c := &Counter{}
	for i := 0; i < 1000; i++ {
		c.Increment()
	}

	log.Println(c.Value())
}
