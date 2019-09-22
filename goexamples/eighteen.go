package main

import (
	"fmt"
	"tempconv"
)

var a tempconv.Celsius = 2465.46

func main() {
	fmt.Println(tempconv.CToF(a))
}
