package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	fmt.Println("nums before removing duplicates:  ", nums)

	k := removeDuplicates(nums)

	fmt.Println("nums after removing duplicates:   ", nums[0:k]) //[0 0 1 1 2 3 3]
}

func removeDuplicates(nums []int) int {
	var k int

	m := make(map[int]int)

	for i := range nums {
		val, ok := m[nums[i]]

		if !ok {
			m[nums[i]] = 1
			nums[k] = nums[i]
			k++
		} else {
			if val < 2 {
				m[nums[i]] = 2
				nums[k] = nums[i]
				k++
			}
		}
	}

	return k
}

// Notes:
// 1. The method removes duplicates in place such that each unique element appears at most twice. i.e a number can only appear twice in the array.
// 2. I used an hashmap to keep track of each number in nums and mapped it to the number of times it appeared.
// 3. If the hashmap returns false, we add the current number to it. i.e m[nums[i]] = 1 . One because its the first time the item appeared. We keep the current item in nums.
// 4. If the hashmap returns true, we check if the value is less than 2. if true, we update its value to 2 since this is the second time the item appears in the array. We keep the current item in nums.
// 5. If hashmap returns true and the value is 2 or more, we ignore that item.
