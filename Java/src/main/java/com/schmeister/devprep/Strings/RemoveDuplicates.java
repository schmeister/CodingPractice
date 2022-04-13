package com.schmeister.devprep.Strings;

import java.util.HashSet;
import java.util.LinkedHashSet;

import com.schmeister.devprep.Print;

public class RemoveDuplicates {
	static String removeDuplicate(String str) {
		
		char[] ch = str.toCharArray();
		int n = ch.length;
		HashSet<Character> s = new LinkedHashSet<>(n - 1);
		for (char x : ch)
			s.add(x);

		StringBuilder sb = new StringBuilder();
		for (char x : s)
			sb.append(x);

		return sb.toString();
	}

	public static void removeDuplicates() {
		String str = "geeksforgeeks";

		Print.print("RemoveDuplicates", removeDuplicate(str));
	}
}
