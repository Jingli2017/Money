package _38__Product_of_Array_Except_Self_

func productExceptSelf(nums []int) []int {
	n := len(nums)

	left := make([]int, n)
	left[0] = 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := 1
	for i := n-2; i >= 0; i-- {
		right *= nums[i+1]
		left[i] *= right
	}

	return left
}
