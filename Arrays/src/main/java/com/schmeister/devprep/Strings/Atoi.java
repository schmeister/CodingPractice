package com.schmeister.devprep.Strings;

import com.schmeister.devprep.Print;

public class Atoi {
	static int myAtoi(String str) {
		// Initialize result
		int res = 0;

		for (int i = 0; i < str.length(); ++i)
			res = (res * 10) + (str.charAt(i) - '0');

		// return result.
		return res;
	}

	// Driver code
	public static void atoi() {
		String str = "89789";

		// Function call
		int val = myAtoi(str);
		Print.print("Atoi", val);
	}
}