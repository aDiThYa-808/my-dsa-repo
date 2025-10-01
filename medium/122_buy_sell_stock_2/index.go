package main;

import "fmt";

func main(){
	prices := []int{7,1,5,3,6,4};

	fmt.Println("Max profit for ",prices," is ",maxProfit(prices)); //7

}

func maxProfit(prices []int) int {
    var profit int;

    for i:=0;i<len(prices)-1;i++{
        isGreenLine := prices[i] < prices[i+1];

        if isGreenLine{
            profit += prices[i+1] - prices[i];
        }
    }

    return profit;
}

// Notes:
// 1. Used Greedy approach
// 2. GreenLine means the current price is less than the next price which will give us a positive difference(prices[i+1] - prices[i] = positive).
// 3. we ignore all adjacent pairs that do not have a GreenLine.
// 4. if it isGreenLine we add the difference to profit.
