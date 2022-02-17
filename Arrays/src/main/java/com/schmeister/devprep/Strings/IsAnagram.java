package com.schmeister.devprep.Strings;

import java.util.Arrays;

import com.schmeister.devprep.Print;

public class IsAnagram {
	/*
	 * function to check whether two strings are anagram of each other
	 */
	static boolean isAnagram(String str1, String str2) {
		char[] c1 = str1.toCharArray();
		char[] c2 = str2.toCharArray();

		// Get lengths of both strings
		int n1 = c1.length;
		int n2 = c2.length;

		// If length of both strings is not same,
		// then they cannot be anagram
		if (n1 != n2)
			return false;

		// Sort both strings
		Arrays.sort(c1);
		Arrays.sort(c2);

		// Compare sorted strings
		for (int i = 0; i < n1; i++)
			if (c1[i] != c2[i])
				return false;

		return true;
	}

	/* Driver Code */
	public static void isAnagram() {
		String str1 = "test";
		String str2 = "tses";
		String str3 = "tset";

		Print.print("IsAnagram", isAnagram(str1, str2));
		Print.print("IsAnagram", isAnagram(str1, str3));
	}
}
