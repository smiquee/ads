// Merge sort implementations

package ms

// Classical implementation with additional arrays/slices
func Mergesort(array []int) []int {
	if len(array) > 1 {
		pivot := int(len(array) / 2)

		left := array[0:pivot]
		right := array[pivot:len(array)]

		left = Mergesort(left)
		right = Mergesort(right)

		return merge(left, right)
	}
	return array
}

func merge(left []int, right []int) []int {
	lleft := len(left)
	lright := len(right)
	ret := make([]int, lleft+lright)

	i, l, r := 0, 0, 0

	for l < lleft && r < lright {
		if left[l] < right[r] {
			ret[i] = left[l]
			l++
		} else {
			ret[i] = right[r]
			r++
		}
		i++
	}

	for l < lleft {
		ret[i] = left[l]
		l++
		i++
	}

	for r < lright {
		ret[i] = right[r]
		r++
		i++
	}

	return ret
}

// Trying an approach to reduce memory consumption by avoid intial copy
// It is a bit slower than the regular implementation
func Mergesort_less_copy(array *[]int, start int, end int) []int {
	if end-start > 1 {
		pivot := int(start + (end-start)/2)

		left := Mergesort_less_copy(array, start, pivot)
		right := Mergesort_less_copy(array, pivot, end)

		return merge(left, right)
	}
	return (*array)[start:end]
}

// Trying to implement a parallel version using goroutines and channels
// It is clearly not efficient and crash with to large array
func Mergesort_parallel(array []int) []int {
	if len(array) > 1 {
		pivot := int(len(array) / 2)

		cleft := make(chan []int)
		cright := make(chan []int)

		go parallel_task(array[0:pivot], cleft)
		go parallel_task(array[pivot:len(array)], cright)

		left := <-cleft
		right := <-cright

		return merge(left, right)
	}
	return array
}

func parallel_task(array []int, c chan []int) {
	if len(array) > 1 {
		pivot := int(len(array) / 2)

		cleft := make(chan []int)
		cright := make(chan []int)

		go parallel_task(array[0:pivot], cleft)
		go parallel_task(array[pivot:len(array)], cright)

		left := <-cleft
		right := <-cright

		c <- merge(left, right)
	}
	c <- array
}
