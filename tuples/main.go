package main

import "fmt"

func powerSeries(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func main() {
	var square int
	var cube int

	square, cube, _ = powerSeries(4)

	fmt.Println("square:", square, " cube:", cube)
}
