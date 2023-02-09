package main

import "strconv"

type Element struct {
	elementValue int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func main() {}
