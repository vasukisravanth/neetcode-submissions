func maxProfit(prices []int) int {
	// minprice:=prices[0]
	// maxprofit:=0
	// for i:=0;i<len(prices)-1;i++{
	// 	if prices[i]<minprice{
	// 		minprice=prices[i]
	// 	}
	// 	profit:=prices[i]-minprice
	// 	if profit>maxprofit{
	// 		maxprofit=profit
	// 	}

	// }

	// return maxprofit
		minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit

}
