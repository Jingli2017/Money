package ___Median_of_Two_Sorted_Arrays

import "math"

/*
	Input: nums1 = [1,2], nums2 = [3,4]
	Output: 2.50000
	Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/
func findMedianSortedArrays(X []int, Y []int) float64 {
	// binary search on the small array
	if len(X) > len(Y) {
		return findMedianSortedArrays(Y, X)
	}
	left, right := 0, len(X)
	// partition
	for left <= right {
		// how many element on left
		parx := (left + right) / 2
		// parx + pary = len(X) + len(Y) + 1
		// +1 to handle odd number of total element, left has one more element
		pary := (len(X)+len(Y)+1)/2 - parx

		var leftMaxX int
		if parx == 0 {
			leftMaxX = math.MinInt32
		} else {
			leftMaxX = X[parx-1]
		}

		var rightMinX int
		if parx == len(X) {
			rightMinX = math.MaxInt32
		} else {
			rightMinX = X[parx]
		}

		var leftMaxY int
		if pary == 0 {
			leftMaxY = math.MinInt32
		} else {
			leftMaxY = Y[pary-1]
		}

		var rightMinY int
		if pary == len(Y) {
			rightMinY = math.MaxInt32
		} else {
			rightMinY = Y[pary]
		}

		if leftMaxX <= rightMinY && leftMaxY <= rightMinX {
			if (len(X)+len(Y))%2 == 0 {
				return (float64(max(leftMaxX, leftMaxY)) + float64(min(rightMinX, rightMinY))) / 2.0
			} else {
				return float64(max(leftMaxX, leftMaxY))
			}
		} else if leftMaxX > rightMinY {
			right = parx - 1
		} else {
			left = parx + 1
		}
	}

	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
