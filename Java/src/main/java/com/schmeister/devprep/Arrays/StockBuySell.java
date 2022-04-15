package com.schmeister.devprep.Arrays;

import com.schmeister.devprep.util.Print;

public class StockBuySell {
	public static void stockBuySell()
    {
        int price[] = { 100, 180, 260, 310, 40, 535, 695 };
 
        // function call
        Print.print("StockBuySell", maxProfit(price));
    }
	
	static int maxProfit(int prices[]) {
        int size = prices.length;
		int maxProfit = 0;

		for (int i = 1; i < size; i++)
			// Only add value when monotonic in the upward direction.
			if (prices[i] > prices[i - 1])
				maxProfit += prices[i] - prices[i - 1];
		return maxProfit;
	}
}
