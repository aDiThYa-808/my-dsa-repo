package main 

import "fmt"

func main(){
	nums := []int{2,3,1,1,4};

	fmt.Println("Can we reach the last index in ",nums," ? : ",canJump(nums)); //true
}

func canJump(nums []int) bool {
    var jumpsLeft int = nums[0]
    var i int;
    for i < len(nums)-1{
        if jumpsLeft < nums[i]{
            jumpsLeft = nums[i]
        }
        jumpsLeft--;
        i++;
        if jumpsLeft < 0{
            return false
        }
    }
    return true;
}

// Notes:
// 1. Everytime we find an item that has a greater value than jumpsLeft, we replace jumpsLeft with that value.
// 2.if jumpsLeft becomes negative before reaching the last index, we return false because we cannot jump any further.
// 3. Treat jumps like fuel points, indexes like pit stops and jumpsLeft as your fuel tank.
// 4. If the fuel points at one stop is greater than the fuel points in you fuel tank, we update your fuel tank to the new higher value.
// 5. Everytime you move one index ahead, your fuel point decreases.
// 6. If your fuel points become negative, we cannot move any further and we return false.
