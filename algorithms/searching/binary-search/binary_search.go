package binarysearch

func search(s []int, v int) int {
	lo, hi := 0, len(s)-1

	for lo <= hi {
		mid := ((hi - lo) >> 1) + lo

		if s[mid] == v {
			return mid
		}

		if s[mid] < v {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
