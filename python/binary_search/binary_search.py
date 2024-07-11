def binary_search(arr, target):
    l = 0
    r = len(arr) - 1

    while l <= r:
        mid = l + (r-l)//2
        sample = arr[mid]

        if sample == target:
            return mid
        if sample < target:
            l = mid + 1
        elif sample > target:
            r = mid - 1

    return -1
