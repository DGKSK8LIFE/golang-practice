package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
