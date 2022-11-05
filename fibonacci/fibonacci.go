package fibonacci

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	// nolint: gomnd
	return Fibonacci(n-1) + Fibonacci(n-2)
}
