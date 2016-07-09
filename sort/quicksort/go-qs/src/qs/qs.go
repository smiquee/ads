// Quicksort implementations

package qs

// Classical implementation with additional arrays/slices
func Quicksort(array *[]int) []int {
	if len(*array) > 1 {
		less := make([]int, 0)
		equal := make([]int, 0)
		greater := make([]int, 0)

		pivot := (*array)[len(*array)/2]

		for x := 0; x < len(*array); x++ {
			if (*array)[x] < pivot {
				less = append(less, (*array)[x])
				continue
			}
			if (*array)[x] == pivot {
				equal = append(equal, (*array)[x])
				continue
			}
			greater = append(greater, (*array)[x])
		}
		less = Quicksort(&less)
		greater = Quicksort(&greater)

		ret := make([]int, 0)
		ret = append(ret, less...)
		ret = append(ret, equal...)
		ret = append(ret, greater...)
		return ret
	}
	return *array
}

func Quicksort_inplace(array *[]int, start int, end int) {
	if end-start > 1 {
		pivot := (*array)[start+((end-start)/2)]

		left := start
		right := end

		for left <= right {
			for (left <= end) && ((*array)[left] < pivot) {
				left++
			}
			for (right >= start) && ((*array)[right] > pivot) {
				right--
			}
			if left <= right {
				tmp := (*array)[left]
				(*array)[left] = (*array)[right]
				(*array)[right] = tmp
				left++
				right--
			}
		}
		Quicksort_inplace(array, left, end)
		Quicksort_inplace(array, start, right)
	}
}
