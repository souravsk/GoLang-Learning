package godsa

import (
	"fmt"
)

func findPivot(arr []int) int {
	var start int = 0
	var end int = len(arr) - 1

	for start >= end {
		var mid int = start + (end-start)/2

		if mid < end && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] < arr[mid+1] {
			return mid - 1
		}
		if arr[mid] <= arr[start] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func RotatedBinarySearch() {
	var arr = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Print(findPivot(arr))
}
