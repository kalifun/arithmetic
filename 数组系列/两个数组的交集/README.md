# 两个数组的交集
## 一、题目
给定两个数组，编写一个函数来计算它们的交集。
## 二、解题
> 刚开始学习可能会想，我用双重for循环就可以了，emmmmmm 貌似也太费时了吧～

### 2.1 思路
2.1.1 我们需要先统计出数组中元素出现的次数。
``` golang
// 1.计算出具体的k出现了几次
result1 := map[int]int{}
for _, val := range array1 {
    result1[val]++
}
```
2.1.2 另一个数组是否在这个数组中出现了？
``` golang
k := 0
for _, va := range array2 {
    if result1[va] > 0 {
        result1[va]--
        array2[k] = va
        k++
    }
}
```
当存在则重置数组，将按照顺序替换数组。
##三、示例
``` golang
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
```