package main

import "testing"

func Benchmark(b *testing.B) {
	c := &Counter{}
	for i := 0; i < 1000; i++ {
		c.Increment()
	}
}
