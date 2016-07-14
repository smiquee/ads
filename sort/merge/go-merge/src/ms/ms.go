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

// Trying an implmentation that should require less memory as less
// elements of the original array are copied
func Mergesort_less_copy_inplace(array *[]int, start int, end int) {
	if end-start > 1 {
		pivot := int(start + (end-start)/2)

		Mergesort_less_copy_inplace(array, start, pivot)
		Mergesort_less_copy_inplace(array, pivot, end)

		merge_inplace(array, start, pivot, end)
	}
}

func merge_inplace(array *[]int, start, pivot, end int) {
	save := make([]int, pivot-start)
	//for l := 0; l < pivot-start; l++ {
	//	save[l] = (*array)[start+l]
	//}
	copy(save, (*array)[start:pivot])
	i := start
	j := 0
	k := pivot

	for j < len(save) && k < end {
		if save[j] <= (*array)[k] {
			(*array)[i] = save[j]
			j++
		} else {
			(*array)[i] = (*array)[k]
			k++
		}
		i++
	}

	for j < len(save) {
		(*array)[i] = save[j]
		j++
		i++
	}

	for k < end {
		(*array)[i] = (*array)[k]
		k++
		i++
	}
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
