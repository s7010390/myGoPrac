package main

import "fmt"

func main() {
	s := []string{"as", "bf", "som"}
	variadic(s...)
}

func variadic(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
