package main

import "fmt"

func twosum(nums []int, k int) []int {
	for a, _ := range nums {
		for b, _ := range nums[:a] {
			if nums[a]+nums[b] == k {
				return []int{a, b}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	numss := []int{3, 2, 4}
	numsss := []int{3, 3}
	fmt.Println(twosum(nums, 9))
	fmt.Println(twosum(numss, 6))
	fmt.Println(twosum(numsss, 6))

}
