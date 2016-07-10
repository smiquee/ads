// Merge sort algorithms implementation benchmark
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"ms"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFile(fname string) (array []int, err error) {
	buffer, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buffer), "\n")
	// Assign cap to avoid resize on every append.
	array = make([]int, 0, len(lines))

	for _, line := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(line) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		array = append(array, n)
	}
	return array, nil
}

func main() {

	var array []int
	var err error

	if len(os.Args) == 2 {
		array, err = readFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		array = make([]int, 10000)
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < 10000; i++ {
			array[i] = rand.Intn(1000)
		}
	}

	//fmt.Println(array)

	start := time.Now()
	//narray := ms.Mergesort(array)
	ms.Mergesort(array)
	stop := time.Now()
	//fmt.Println(narray)
	var duration time.Duration = stop.Sub(start)
	fmt.Printf("mergesort:          %.9fs\n", duration.Seconds())
}
