package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flow", "flower", "flight"}
	fmt.Println(SearchPrefix(strs))
}

func SearchPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, val := range strs {
		// for strings.Index(val, prefix) != 0 {
		// 	if len(prefix) == 0 {
		// 		return ""
		// 	}
		// 	prefix = prefix[:len(prefix)-1]
		// }
		fmt.Println("val:", val)
		for strings.Index(val, prefix) != 0 {
			fmt.Println("prefix:", prefix)
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
