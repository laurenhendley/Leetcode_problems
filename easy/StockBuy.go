func maxProfit(prices []int) int {
    max_profit := 0
    min_price := prices[0]

    for i := 0; i<=len(prices)-1;i++ {
        if prices[i] < min_price {
            min_price = prices[i]
        } else if prices[i] - min_price > max_profit {
            max_profit = prices[i] - min_price
        }
    }

    return max_profit
}
