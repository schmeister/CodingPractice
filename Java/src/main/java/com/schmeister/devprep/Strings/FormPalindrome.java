package com.schmeister.devprep.Strings;

import java.util.Arrays;

import com.schmeister.devprep.util.Print;

public class FormPalindrome {

	static int findMinInsertions(char str[], int l, int h) {

		if (l > h)
			return Integer.MAX_VALUE;
		if (l == h)
			return 0;
		if (l == h - 1)
			return (str[l] == str[h]) ? 0 : 1;

		int retVal;

		if ((str[l] == str[h]))
			retVal = findMinInsertions(str, l + 1, h - 1);
		else {
			int left = findMinInsertions(str, l, h - 1);
			int right = findMinInsertions(str, l + 1, h);
			retVal = Integer.min(left, right) + 1;
		}
		return retVal;
	}

	// A DP function to find minimum number
	// of insertions
	static int findMinInsertionsDP(char str[], int n) {
		// Create a table of size n*n. table[i][j]
		// will store minimum number of insertions
		// needed to convert str[i..j] to a palindrome.
		int table[][] = new int[n][n];
		int l, h, gap;

		// Fill the table
		for (gap = 1; gap < n; ++gap)
			for (l = 0, h = gap; h < n; ++l, ++h)
				table[l][h] = (str[l] == str[h]) ? table[l + 1][h - 1]
						: (Integer.min(table[l][h - 1], table[l + 1][h]) + 1);

		// Return minimum number of insertions
		// for str[0..n-1]
		return table[0][n - 1];
	}

	// Driver program to test above function.
	public static void formPalindrome() {
		String str = "geeks";
		Print.print("FormPal", str, findMinInsertionsDP(str.toCharArray(), str.length()));
		Print.print("FormPal", str, findMinInsertions(str.toCharArray(), 0, str.length() - 1));
	}
}
