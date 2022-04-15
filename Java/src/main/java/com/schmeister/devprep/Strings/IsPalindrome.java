package com.schmeister.devprep.Strings;

import com.schmeister.devprep.App;
import com.schmeister.devprep.util.Print;

public class IsPalindrome {
	static boolean isPalindrome(String str) {
		int i = 0;
		int j = str.length() - 1;

		while (i < j) {

			if (str.charAt(i) != str.charAt(j))
				return false;

			i++;
			j--;
		}
		return true;
	}

	public static void isPalindrome() {
		String str = "geeks";
		Print.print("IsPalindrome", str, isPalindrome(str));
		str = "geeksskeeg";
		Print.print("IsPalindrome", str, isPalindrome(str));
		str = "geekskeeg";
		Print.print("IsPalindrome", str, isPalindrome(str));
	}
}
