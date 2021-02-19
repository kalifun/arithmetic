package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twosum(nums, target))
}

func twosum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
