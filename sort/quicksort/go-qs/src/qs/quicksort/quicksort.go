// desc
package main

import (
	"fmt"
	"math/rand"
	"qs"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {

	array := make([]int, 100)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 100; i++ {
		array[i] = rand.Intn(100)
	}

	start := time.Now()
	// narray := qs.Quicksort(&array)
	qs.Quicksort(&array)
	stop := time.Now()
	var duration time.Duration = stop.Sub(start)
	fmt.Printf("quicksort: %.9fs\n", duration.Seconds())
}
