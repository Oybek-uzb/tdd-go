package main

import "fmt"

func main() {
	counter := Counter{43}
	counter.Inc()

	fmt.Println(counter.Value())
}

type Counter struct {
	Val int
}

func (c *Counter) Inc() {
	c.Val++
}

func (c *Counter) Value() int {
	return c.Val
}
