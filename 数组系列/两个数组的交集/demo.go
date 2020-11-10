package main

import "fmt"

func main() {
	// 定义两个数组
	a := []int{1, 2, 2, 5, 8}
	b := []int{2, 7, 9, 2}
	intersect(a, b)
	// fmt.Println(a, b)
}

func intersect(array1, array2 []int) {
	// 1.计算出具体的k出现了几次
	result1 := map[int]int{}
	for _, val := range array1 {
		result1[val]++
	}
	fmt.Println(result1)
	k := 0
	for _, va := range array2 {
		if result1[va] > 0 {
			result1[va]--
			array2[k] = va
			k++
		}
	}
	fmt.Println(array2)
	fmt.Println(array2[0:k])
}
