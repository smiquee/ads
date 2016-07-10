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
