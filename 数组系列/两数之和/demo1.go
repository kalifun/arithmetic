package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twosum(nums, target))
}

func twosum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
		fmt.Println(m)
	}
	return result
}
