package main

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return true
	}
	var index uint = 0
	for index < uint(len(nums)-1) {
		step := nums[index] // value on the indexed position
		if step == 0 {      // break if unable to jump
			return false
		}
		index += step // jump the index to certain position
	}
	if index == uint(len(nums)-1) {
		return true
	}
	return false
}
