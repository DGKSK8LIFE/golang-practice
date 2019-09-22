package main

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return 5
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
