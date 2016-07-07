#!/usr/bin/python


def qsort(array):
    """
    Classical quicksort using additional arrays.
    """
    less = list()
    equal = list()
    greater = list()

    if len(array) > 1:
        pivot = array[len(array) / 2]

        for x in array:
            if x < pivot:
                less.append(x)
                continue
            if x == pivot:
                equal.append(x)
                continue
            greater.append(x)
        less = qsort(less)
        greater = qsort(greater)

        return less + equal + greater
    return array


def qsort_inplace(array, start=0, end=-1):
    """
    Quicksort that sort the array in place, with no additional
    arrays.
    """
    if end == -1:
        end = len(array) - 1
    if end - start > 1:
        pivot = array[start + ((end - start) / 2)]
        left = start
        right = end

        while left <= right:
            while array[left] < pivot and left < end:
                left += 1
            while array[right] > pivot and right > start:
                right -= 1
            if left <= right:
                tmp = array[right]
                array[right] = array[left]
                array[left] = tmp
                left += 1
                right -= 1
        qsort_inplace(array, left, end)
        qsort_inplace(array, start, right)


def main(nb):
    """
    Main function.
    """
    import time
    import numpy.random as nprnd
    import psutil
    import gc
    import os

    gc.disable()
    process = psutil.Process(os.getpid())
    array = list(nprnd.randint(nb, size=nb))
    gc.collect()
    start = time.time()
    qsort(array)
    stop = time.time()
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print "qsort:         %fs   (%f)" % (stop - start, mem)

    gc.collect()

    start = time.time()
    qsort_inplace(array)
    stop = time.time()
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print "qsort_inplace: %fs   (%f)" % (stop - start, mem)


if __name__ == "__main__":
    """
    If called as a program.
    """
    import sys
    if len(sys.argv) == 2:
        main(int(sys.argv[1]))
    else:
        main(100)
