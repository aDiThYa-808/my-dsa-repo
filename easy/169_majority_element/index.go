package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 1, 2, 1}
	fmt.Println("Majority element in ", nums, " is :", majorityElement(nums)) //1
}

func majorityElement(nums []int) int {
	var count int
	var candidate int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// Notes:
// 1. Uses Boyer Moore voting algorithm.
// 2. For [1 2 3 1 1 2 1], majorityElement does the following:
// See 1 → count=0 → candidate=1, count=1
// See 2 → 2 != 1 → count=0
// See 3 → count=0 → candidate=3, count=1
// See 1 → 1 != 3 → count=0 → candidate=1, count=1
// See 1 → 1 == 1 → count=2
// See 2 → 2 != 1 → count=1
// See 1 → 1 == 1 → count=2
