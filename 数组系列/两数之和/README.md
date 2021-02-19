# 两数之和  

## 题目  

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。  

示例 1：

```
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
```  

示例 2：

```
输入：nums = [3,2,4], target = 6
输出：[1,2]
```

示例 3：

```
输入：nums = [3,3], target = 6
输出：[0,1]
```  

## 解题  
首先我们拿到题目一看，马上可以想到暴力题解。我们只需要 “遍历每个元素 x，并查找是否存在一个值与 target - x 相等的目标元素。”  

``` golang
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
```  
 
1. 首先，我们还是先遍历数组 nums，i 为当前下标。我们需要将每一个遍历的值放入 map 中作为 key。 
2. 同时，对每个值都判断 map 中是否存在 target-nums[i] 的 key 值。在这里就是 9-7=2。我们可以看到 2 在 map 中已经存在。
3. 所以，2 和 7 所在的 key 对应的 value，也就是 [0,1]。就是我们要找的两个数组下标。  

```golang
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
```