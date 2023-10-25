package main

import (
	"fmt"
	ss "strings"
)

func main() {
	m := make(map[string]string)
	if m == nil {
		fmt.Println("is nil")
	}

	m["a"] = "apple"
	s := m["a"]
	fmt.Println(s)
}

func wordCount(s string) map[string]int {
	split := ss.Split(s, " ")
	result := make(map[string]int)
	for i := 0; i < len(split); i++ {
		result[split[i]] = result[split[i]] + 1
	}
	return result
}
