package main

import "fmt"

type Container[T any] struct {
	value T
}

func (c *Container[T]) Set(value T) {
	c.value = value
}

func (c *Container[T]) Get() T {
	return c.value
}

func main() {
	intContainer := Container[int]{}
	intContainer.Set(42)
	fmt.Println("Int value:", intContainer.Get())

	stringContainer := Container[string]{}
	stringContainer.Set("hello")
	fmt.Println("Int value:", stringContainer.Get())
}
