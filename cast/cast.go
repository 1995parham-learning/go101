package main

import "fmt"

type example interface {
	notImplemented()
}

type empty interface{}

func main() {
	one := 1

	var i empty = one

	float := float32(one)

	switch i.(type) {
	default:
		fmt.Printf("Type error\n")
	case int:
		fmt.Printf("%d\n", i)
	}

	fmt.Printf("%f\n", float)

	// This will panic at runtime
	e := i.(example)
	fmt.Printf("%d\n", e.(empty).(int))
}
