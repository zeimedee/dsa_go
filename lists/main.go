package main

import (
	"container/list"
	"fmt"
)

func main() {

	var Intlist list.List

	Intlist.PushBack(11)
	Intlist.PushBack(12)
	Intlist.PushBack(13)

	for element := Intlist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
