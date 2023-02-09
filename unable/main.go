package main

import (
	"fmt"
)

func countStudents(students []int, sandwiches []int) int {

	for {

		if students[0] == sandwiches[0] {
			sandwiches = pop(sandwiches)
			students = pop(students)
			fmt.Println("st:", students)
			fmt.Println("sd:", sandwiches)
			if len(students) == 0 && len(sandwiches) == 0 {
				return len(students)
			}

		}
		if students[0] != sandwiches[0] {
			students = circle(students)
			fmt.Println("rst:", students)
			fmt.Println("rsd:", sandwiches)
			t := contains(students, sandwiches[0])
			fmt.Println(t)
			if t {
				return len(students)
			}

		}
	}
}

func pop(s []int) []int {
	// _, s = s[0], s[1:]
	s = s[1:]
	return s
}

func circle(students []int) []int {
	x, students := students[0], students[1:]
	students = append(students, x)
	return students
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return false
		}
	}

	return true
}

func main() {
	// students := []int{1, 1, 0, 0}
	// sandwichs := []int{0, 1, 0, 1}

	students := []int{1, 1, 1, 0, 0, 1}
	sandwichs := []int{1, 0, 0, 0, 1, 1}

	n := countStudents(students, sandwichs)
	fmt.Println(n)

}
