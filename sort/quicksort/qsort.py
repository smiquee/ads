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


def main(input=None):
    """
    Main function.
    """
    import time
    import psutil
    import gc
    import os

    if input is None:
        import numpy.random as nprnd
        array = list(nprnd.randint(10000, size=10000))
    else:
        array = list()
        try:
            with open(input, 'r') as _f:
                for elt in _f.readlines():
                    array.append(int(elt))
        except (IOError, ValueError) as err:
            os.sys.stderr.write("%s\n" % err)
            os.sys.exit(1)

    # print(array)
    gc.disable()
    process = psutil.Process(os.getpid())
    gc.collect()
    start = time.time()
    # narray = qsort(array)
    qsort(array)
    stop = time.time()
    # print(narray)
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print("qsort:         %fs   (%f)" % (stop - start, mem))

    gc.collect()

    start = time.time()
    qsort_inplace(array)
    stop = time.time()
    # print(array)
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print("qsort_inplace: %fs   (%f)" % (stop - start, mem))

    start = time.time()
    array.sort()
    stop = time.time()
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print("python_sort:   %fs   (%f)" % (stop - start, mem))


if __name__ == "__main__":
    """
    If called as a program.
    """
    import sys
    if len(sys.argv) == 2:
        main(sys.argv[1])
    else:
        main()
