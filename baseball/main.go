package main

import (
	"fmt"
)

func baseball(ops []any) {

	// var arr []any
	for _, i := range ops {
		fmt.Println(i)
	}

}

func main() {

	ops := []any{5, 2, "D", "C", "+"}

	baseball(ops)

}
