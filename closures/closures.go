// we are showing that an inner function can capture
// its scope's variables.
package main

import "fmt"

func createFunc() func(i int) int {
	a := 1
	b := 1

	innerFunc := func(i int) int {
		fmt.Printf("[%d] a = %d\n", i, a)
		fmt.Printf("[%d] b = %d\n", i, b)

		b++

		return 0
	}

	a++

	return innerFunc
}

func main() {
	innerFunc := createFunc()

	// at the end innerFunc read `a` and `b` on the time of its execution.
	// please note that innerFunc hold a reference to `a` and `b`.
	for i := 1; i < 3; i++ {
		innerFunc(i)
	}
}
