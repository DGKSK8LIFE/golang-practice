package main

import (
	"fmt"
)

func main() {
	fmt.Println(basename("b.c"))
}
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[1] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
