package com.schmeister.devprep.Strings;

import com.schmeister.devprep.util.Print;

public class Strstr {
	public static int strstr(String s2, String s1) {
		int counter = 0; // pointing s2
		int i = 0;
		for (; i < s1.length(); i++) {
			if (counter == s2.length())
				break;
			if (s2.charAt(counter) == s1.charAt(i)) {
				counter++;
			} else {
				// Special case where character preceding the i'th character is duplicate
				if (counter > 0) {
					i -= counter;
				}
				counter = 0;
			}
		}
		return counter < s2.length() ? -1 : i - counter;
	}

	public static void strstr() {
		String s1 = "geeksfffffoorrfoorforgeeks";
		String s2 = "geeksforgeeks";
		String find1 = "fr";
		String find2 = "for";
		// System.out.println(s2.indexOf("for"));
		Print.print("Strstr", strstr(find1, s1));
		Print.print("Strstr", strstr(find2, s2));
	}
}
