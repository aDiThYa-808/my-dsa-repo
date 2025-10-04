package main

import "fmt"

func main(){
	nums:= []int{7,0,9,6,9,6,1,7,9,0,1,2,9,0,3}

	fmt.Println("Minimum jumps to reach the last index in ",nums," is: ",jump(nums));
}

func jump(nums []int) int {
    var jumps int;
    var farthest int;
    var rangeEnd int;
    var i int;

    for i < len(nums)-1{
        farthest = max(farthest,i+nums[i])

        if i == rangeEnd{
            rangeEnd = farthest
            jumps++
        }

        if rangeEnd >= len(nums)-1{
            break;
        }

        i++
    }

    return jumps;
}

func max(a int , b int) int{
    if a>b{
        return a;
    }else{
        return b;
    }
}

// Notes:
// 1. rangeEnd holds the last index to which we can jump from the last jump point.
// 2. we use farthest to keep track of the farthest index we can jump to, from an index within rangeEnd.
// 3. when we reach rangeEnd, we update it to the farthest index from the range creating a new range.
// 4. we increment jump only when we are at the rangeEnd.
// 5. we break out of the loop when we reach the end to prevent iterations that are not needed.
