package com.schmeister.devprep.Numbers;

import com.schmeister.devprep.util.Print;

public class Prime {
	static int i = 2;

	public static boolean isPrime(int n) {

		// Corner cases
		if (n == 0 || n == 1) {
			return false;
		}

		// Checking Prime
		if (n == i)
			return true;

		// Base cases
		if (n % i == 0) {
			return false;
		}
		i++;
		return isPrime(n);
	}

	// Driver Code
	public static void prime() {
		int prev = 0;
		for (int x = 1; x < 30000; x++) {
			i = 2;
			if (isPrime(x)) {
				Print.print("Prime", x, x-prev);
				prev = x;
			}
		}
	}
}
