package main

import (
	"fmt"
)

func main() {
	context := Context{}
	a := 10
	b := 5

	context.SetAlgorithm(&Sum{})
	fmt.Println(context.Execute(&a, &b))

	context.SetAlgorithm(&AntiSum{})
	fmt.Println(context.Execute(&a, &b))

	context.SetAlgorithm(&Multiply{})
	fmt.Println(context.Execute(&a, &b))

}

type SelectedAlgo interface {
	execute(a, b *int) int
}

type Context struct {
	selectedAlgo SelectedAlgo
}

func (c *Context) SetAlgorithm(e SelectedAlgo) {
	c.selectedAlgo = e
}

func (c *Context) Execute(a, b *int) int {
	return c.selectedAlgo.execute(a, b)
}

type Multiply struct {
}

func (s *Multiply) execute(a, b *int) int {
	return *a * *b
}

type Sum struct {
}

func (s *Sum) execute(a, b *int) int {
	return *a + *b
}

type AntiSum struct {
}

func (m *AntiSum) execute(a, b *int) int {
	return *a - *b
}
