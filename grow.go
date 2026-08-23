package main

func grow(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + grow(n-1)
}
