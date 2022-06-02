package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 50
	result := []string{}
	// newArr := []string{}
	for i := 0; i <= n; i++ {
		if i%3 == 0 && i > 2 {
			result = append(result, "Frontend")
		} else if i%5 == 0 && i > 4 {
			result = append(result, "Backend")
		} else if i > 0 {
			conv := strconv.Itoa(i)
			result = append(result, conv)
		}
	}
	// for i := 0; i <= len(newArr)*2; i++ {
	// }
	fmt.Println(result)
}
