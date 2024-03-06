package main

import "fmt"

//Встраивание интерфейса

type counter struct {
	val uint
}

func (c *counter) increment() {
	c.val++
}

func (c *counter) value() uint {
	return c.val
}

type Counter interface {
	increment()
	value() uint
}

type usage struct {
	service string
	Counter
}

func newUsage(service string) *usage {
	return &usage{service, &counter{}}
}

func main() {

	usage := newUsage("find")
	usage.increment()
	usage.increment()
	usage.increment()
	usage.increment()
	fmt.Printf("%s usage: %d\n", usage.service, usage.value())

}
