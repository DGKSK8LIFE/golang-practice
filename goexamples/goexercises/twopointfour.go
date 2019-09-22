package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountByShifting(x uint64) int {
	count := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return count
}
