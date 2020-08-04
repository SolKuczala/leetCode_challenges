func isPowerOfTwo(n int) bool {
	result := 1
	for result < n {
		result = result * 2
	}
	return result == n
}
