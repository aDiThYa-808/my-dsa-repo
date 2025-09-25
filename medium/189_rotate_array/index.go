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

    mainReversedArr:= make([]int,0,len(nums));

    for i:=len(nums)-1;i>=0;i--{
        mainReversedArr = append(mainReversedArr,nums[i]); 
    }

    firstArr := make([]int,0,k); 
    secondArr := make([]int,len(nums)-k);

    for i:=0;i<k;i++{
        firstArr = append(firstArr,mainReversedArr[i]);
    }

    for i:=k;i<len(mainReversedArr);i++{
        secondArr = append(secondArr,mainReversedArr[i]);     
    }

    finalReversedArr := make([]int,0,len(nums))

    for i:=len(firstArr)-1;i>=0;i--{
        finalReversedArr = append(finalReversedArr,firstArr[i]);
    }
    for i:=len(secondArr)-1;i>=0;i--{
        finalReversedArr = append(finalReversedArr,secondArr[i]);
    }

    copy(nums,finalReversedArr);
}

// Notes: 
// 1. k is normalized to ignore redundant rotations. 
// 2. For example , say len(nums) is 2 and k is 7. 
//    Rotating nums 5 times is going to give back the same array, so we ignore the first 5 rotations and only do the remaining 2. 
//    7 % 5 = 2.
// 3. We reverse the nums array and split it into 2 new slices. one containing first k items and the other containing the remaining items.
// 4. We reverse both the arrays and append them to finalReversedArr to get the result. we copy it to nums.

