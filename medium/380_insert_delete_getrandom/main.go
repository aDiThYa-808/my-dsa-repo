package main 

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
    array []int
    hmap map[int]int
}


func Constructor() RandomizedSet {
    array := make([]int,0)
    hmap := make(map[int]int) 

    return RandomizedSet{array,hmap}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,ok := this.hmap[val]; !ok{
        this.array = append(this.array,val)
        this.hmap[val] = len(this.array) - 1
        return true
    }

    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if index,ok := this.hmap[val]; ok{
        last := this.array[len(this.array)-1]
        this.array[index] = last
        this.hmap[last] = index
        delete(this.hmap,val)
        this.array = this.array[:len(this.array)-1]
        return true
    }

    return false
}


func (this *RandomizedSet) GetRandom() int {
    return this.array[rand.Intn(len(this.array))]
}

func main(){
	randomizedSet := Constructor()

	fmt.Println(randomizedSet.Insert(1)) // true
	fmt.Println(randomizedSet.Remove(2)) // false
	fmt.Println(randomizedSet.Insert(2)) //true
	fmt.Println(randomizedSet.GetRandom()) // either 1 or 2
}

// Notes:
// 1. This is a simple problem where you perfom insert,delete and getRandom on a set with all 3 functions having a time complexity of O(1)
// 2. The set is stored in the array and the hashmap keeps track of each items index.
