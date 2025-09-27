package main;

import "fmt";

func main(){
	var prices []int = []int{7,1,5,3,6,4};
	profit := maxProfit(prices);
	fmt.Println("Maximum profit for ",prices," is ",profit); //5
}

func maxProfit(prices []int) int {
    var profit int;
    var buy int = prices[0];

    for i:=1;i<len(prices);i++{
        isBetterProfit := prices[i] - buy > profit;
        if buy > prices[i]{
            buy = prices[i];
        }else if isBetterProfit{
            profit = prices[i] - buy
        }
    }
    return profit;
}

// Notes:
// 1. We store the first value in buy and change it when we find a smaller value in prices array.
// 2. if the prices[i] is greater than buy value, we check if the profit is greater than the previous profit. if yes ,we update profit.
