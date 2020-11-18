package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 4, 13}
	arr2 := []int{10, 9, 3, 2, 1}
	intersect(arr1, arr2)
}

func intersect(array1, array2 []int) {
	i, j, k := 0, 0, 0
	sort.Ints(array1)
	sort.Ints(array2)
	for i < len(array1) && j < len(array2) {
		if array1[i] < array2[j] {
			j++
		} else if array1[i] > array2[j] {
			i++
		} else {
			array1[k] = array1[i]
			i++
			j++
			k++
		}
	}
	fmt.Println(array1[:k])
}
