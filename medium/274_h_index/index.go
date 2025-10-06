package main 

import (
	"fmt"
	"sort"
)

func main(){
	citations := []int{3,0,6,1,5}

	fmt.Println("H index for ",citations," is: ")
	fmt.Println(hIndex(citations)) //3
}

func hIndex(citations []int) int {
    sort.Ints(citations)
    var hIndex int
    var tempH int

    for i:=0; i<len(citations);i++{
        tempH = len(citations) - i
        if citations[i] >= tempH && tempH > hIndex{
            hIndex = tempH
            break
        }
    }

    return hIndex
}

// Notes:
// 1. H index is the max value of h for which the person has atleast h papers with atleast h citations.
// 2. tempH holds the number papers for each citations[i] that have citations[i] or more citations.
// 3. hIndex keeps track of the maximum value of tempH that was found.
// 4. we break out of the loop when we find the first item that is greater or equal to tempH becuase every item to its right will have a lower tempH value. 
