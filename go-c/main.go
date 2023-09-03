package main

// #include "hello.h"
import "C"

import (
	"fmt"
	"time"
)

const (
	waitTime = 1 * time.Hour
	period   = 100 * time.Millisecond
)

func main() {
	fmt.Println("Hello world... in Go")
	C.say_hello(C.CString("Parham Alvani"))
	C.setup_timer()
	C.setup_signal_handler()

	go func() {
		timer := time.Tick(period)

		for {
			<-timer
			fmt.Println("100 miliseconds in Go")
		}
	}()

	// wait for a signal to handle it in C
	time.Sleep(waitTime)
}
