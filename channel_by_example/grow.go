package channel

func Grow(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + Grow(n-1)
}
