package main

import "fmt"

func main(){
	var nums1 []int = []int{2,7,11,15};
	var nums2 []int = []int{3,2,4};
	var nums3 []int = []int{3,3};

	fmt.Println("nums: [2,7,11,15] target: 9 : ",twoSum(nums1,9)); //[1 0]
	fmt.Println("nums: [3,2,4]     target: 6 : ",twoSum(nums2,6)); //[2 1]
	fmt.Println("nums: [3,3]       target: 6 : ",twoSum(nums3,6))  //[0 1]
}

func twoSum(nums []int, target int) []int {
    res:= make([]int,0,2);
    hash_map := make(map[int]int);
    for i:=0;i<len(nums);i++{
        y := target - nums[i];
        val,ok := hash_map[y];
        if ok{
            res = append(res,i);
            res = append(res,val);
            return res;
        }
        hash_map[nums[i]]=i
    }
    return res;
}

// Notes:
// 1. nums[i] + y = target, therefore y = target - nums[i]. If y is present in the hashmap, we append indexes of nums[i] and y to result slice. else we add the current item and its index to the hashmap.

