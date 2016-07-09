#!/usr/bin/python


def merge_sort(array):
    """
    Classical merge sort algorithm
    """
    if len(array) <= 1:
        return array

    pivot = len(array) / 2

    left = array[:pivot]
    right = array[pivot:]

    left = merge_sort(left)
    right = merge_sort(right)

    return _merge(left, right)


def _merge(left, right):
    ret = list()
    lleft = len(left)
    lright = len(right)

    l = 0
    r = 0

    while l < lleft and r < lright:
        if left[l] <= right[r]:
            ret.append(left[l])
            l += 1
        else:
            ret.append(right[r])
            r += 1

    while l < lleft:
        ret.append(left[l])
        l += 1
    while r < lright:
        ret.append(right[r])
        r += 1

    return ret


# The following merge function is clearly not efficient!
def _merge2(left, right):

    ret = list()

    while len(left) > 0 and len(right) > 0:
        if left[0] <= right[0]:
            ret.append(left[0])
            left = left[1:]
        else:
            ret.append(right[0])
            right = right[1:]

    while len(left) > 0:
        ret.append(left[0])
        left = left[1:]
    while len(right) > 0:
        ret.append(right[0])
        right = right[1:]

    return ret


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
    # narray = merge_sort(array)
    merge_sort(array)
    stop = time.time()
    # print(narray)
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print("merge_sort:         %fs   (%f)" % (stop - start, mem))

    start = time.time()
    array.sort()
    stop = time.time()
    mem = process.get_memory_info()[0] / float(2 ** 20)
    print("python_sort:        %fs   (%f)" % (stop - start, mem))


if __name__ == "__main__":
    """
    If called as a program.
    """
    import sys
    if len(sys.argv) == 2:
        main(sys.argv[1])
    else:
        main()
