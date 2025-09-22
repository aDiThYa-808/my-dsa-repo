package main

import "fmt"

func main() {
	fmt.Println("abcabcbb: ", lengthOfLongestSubstring("abcabcbb")) //3
	fmt.Println("pwwkew:   ", lengthOfLongestSubstring("pwwkew"))   //3
	fmt.Println("abba:     ", lengthOfLongestSubstring("abba"))     //2
}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	hash_map := make(map[uint8]int)
	var start int
	var currLen int
	var maxLen int

	for i := 0; i < len(s); i++ {
		val, ok := hash_map[s[i]]

		if !ok {
			hash_map[s[i]] = i
			currLen = i - start + 1
		} else {
			if start <= val {
				start = val + 1
			}
			hash_map[s[val]] = i
			currLen = i - start + 1
		}

		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

// Notes:
// 1. s[i] returns ASCII code of the value (ex: a = 97). Therefore the value type in the hashmap is uint8.
// 2. hash_map[s[i]] returns a index value along with a false if item not found and vice versa. if !ok item not found in the hashmap.
// 3. The solution used sliding window technique.
