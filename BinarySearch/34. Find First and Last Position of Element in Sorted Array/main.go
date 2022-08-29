package _4__Find_First_and_Last_Position_of_Element_in_Sorted_Array

func searchRange(nums []int, target int) []int {
	front := binarySearch(nums, target, true)
	if front == -1 {
		return []int{-1, -1}
	}
	end := binarySearch(nums, target, false)
	return []int{front, end}
}

func binarySearch(nums []int, target int, isFront bool) int {
	left := 0
	right := len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid - 1
		}else{
			if isFront {
				if mid == left || nums[mid-1] != nums[mid] {
					return mid
				}else {
					right = mid - 1
				}
			}else{
				if mid == right || nums[mid+1] != nums[mid]{
					return mid
				}else {
					left = mid + 1
				}
			}
		}
	}

	return -1
}
