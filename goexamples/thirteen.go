package main

import "fmt"

func main() {
	fmt.Println(Signum(-5))
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}

}
