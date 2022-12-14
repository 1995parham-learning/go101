package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const (
	threadNum = 4

	width  = 1000
	height = 1000
)

func main() {
	var N int

	ch := make(chan int, threadNum)
	inner := 0

	fmt.Scanf("%d", &N)

	for i := 0; i < threadNum; i++ {
		go func(n int) {
			inner := 0

			for i := 0; i < n; i++ {
				xbig, err := rand.Int(rand.Reader, big.NewInt(width))
				if err != nil {
					log.Fatal(err)
				}

				x := xbig.Int64()

				ybig, err := rand.Int(rand.Reader, big.NewInt(height))
				if err != nil {
					log.Fatal(err)
				}

				y := ybig.Int64()

				if (x-width/2)*(x-width/2)+(y-height/2)*(y-height/2) <= (width/2)*(height/2) {
					inner++
				}
			}
			ch <- inner
		}(N / threadNum)
	}

	for i := 0; i < threadNum; i++ {
		inner += <-ch
	}

	fmt.Printf("%f\n", 4.0*float64(inner)/float64(N))
}
