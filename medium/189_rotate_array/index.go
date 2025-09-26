package main;

import "fmt";

func main(){
	nums := []int{-1,-100,3,99};
	nums1:= []int{1,2,3,4,5,6,7}
	nums2 := []int{1,2};
	k := 2;
	k1:=3;
	k2:=7;
	fmt.Println("Before rotation: ",nums," ",nums1," ",nums2);
	rotate(nums,k);
	rotate(nums1,k1);
	rotate(nums2,k2);
	fmt.Println("After rotation: ",nums," ",nums1," ",nums2); // [3 99 -1 -100] [5 6 7 1 2 3 4] [2 1]
}

func rotate(nums []int, k int)  {

    k = k % len(nums); // normalize k

    if k == 0{
        return;
    }

    var end int = len(nums)-1;

    for start:= range nums{
        var temp int = nums[end];
        if end > start{
            nums[end] = nums[start];
            nums[start] = temp;
            end--;
        }
    }

    end = k-1;

    for start:=0 ; start < k ; start++{
        var temp int = nums[end];
        if end > start{
            nums[end] = nums[start];
            nums[start] = temp;
            end--;
        }
    }

    end = len(nums)-1;

    for start:=k;start<len(nums);start++{
        var temp int = nums[end];
        if end>start{
            nums[end] = nums[start];
            nums[start] = temp;
            end--;
        }
    }
    
}


// Notes: 
// 1. k is normalized to ignore redundant rotations. 
// 2. For example , say len(nums) is 2 and k is 7. 
//    Rotating nums 5 times is going to give back the same array, so we ignore the first 5 rotations and only do the remaining 2. 
//    7 % 5 = 2.
// 3. We first reverse the nums array.
// 4. second, we reverse the first k items in the reversed nums array in-place.
// 5. finally, we reverse the remaining items in the reversed nums array in-place.

