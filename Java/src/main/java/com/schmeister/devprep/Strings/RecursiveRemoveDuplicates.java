package com.schmeister.devprep.Strings;

import com.schmeister.devprep.util.Print;

public class RecursiveRemoveDuplicates {
	private static String removeDuplicates(String s, char ch) {
		
		// If length of string is 1 or 0
		if (s == null || s.length() <= 1) {
			return s;
		}

		int i = 0;
		while (i < s.length()) {
			if (i + 1 < s.length() && s.charAt(i) == s.charAt(i + 1)) {
				int j = i;
				while (j + 1 < s.length() && s.charAt(j) == s.charAt(j + 1)) {
					j++;
				}
				char lastChar = i > 0 ? s.charAt(i - 1) : ch;

				String remStr = removeDuplicates(s.substring(j + 1, s.length()), lastChar);

				s = s.substring(0, i);

				// Recursively remove all the adjacent
				// characters formed by removing the
				// adjacent characters
				while (remStr.length() > 0 && s.length() > 0 && remStr.charAt(0) == s.charAt(s.length() - 1)) {

					// Have to check whether this is the
					// repeated character that matches the
					// last char of the parent String
					while (remStr.length() > 0 && remStr.charAt(0) != ch
							&& remStr.charAt(0) == s.charAt(s.length() - 1)) {
						remStr = remStr.substring(1, remStr.length());
					}
					s = s.substring(0, s.length() - 1);
				}
				s = s + remStr;
				i = j;
			} else {
				i++;
			}
		}
		return s;
	}

	// Driver Code
	public static void recursiveRemoveDups() {

		String str1 = "mississipie";
		Print.print("RecRemDup", removeDuplicates(str1, ' '));
		String str2 = "ocvvcolop";
		Print.print("RecRemDup", removeDuplicates(str2, ' '));
	}
}
