// Quicksort algorithms implementation benchmark
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

	//fmt.Println(array)

	start := time.Now()
	//narray := qs.Quicksort(&array)
	qs.Quicksort(&array)
	stop := time.Now()
	//fmt.Println(narray)
	var duration time.Duration = stop.Sub(start)
	fmt.Printf("quicksort        : %.9fs\n", duration.Seconds())

	start = time.Now()
	qs.Quicksort_inplace(&array, 0, len(array)-1)
	stop = time.Now()
	duration = stop.Sub(start)
	fmt.Printf("quicksort_inplace: %.9fs\n", duration.Seconds())
	//fmt.Println(array)
}
