package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	fmt.Println("nums before removal:        ", nums)

	k := removeElement(nums, val) // each of the first k values in nums != val.

	fmt.Println("val to be removed:          ", val)
	fmt.Println("First k values without val: ", nums[0:k]) //[0 1 3 0 4]
}

func removeElement(nums []int, val int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		} else {

		}
	}
	return j
}

// Notes:
// 1. removeElement method removes val from nums in-place.
// 2. removeElement returns j. All values from nums[0] to nums[j] are != val.
