package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println("nums before removing duplicates:  ", nums)

	k := removeDuplicates(nums)

	fmt.Println("nums after removing duplicates:   ", nums[0:k]) //[0 1 2 3 4]

}

func removeDuplicates(nums []int) int {
	var k int
	hash_map := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := hash_map[nums[i]]

		if !ok {
			hash_map[nums[i]] = i
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// Notes:
