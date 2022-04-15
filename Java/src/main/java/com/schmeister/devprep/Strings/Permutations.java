package com.schmeister.devprep.Strings;

import java.util.Vector;

import com.schmeister.devprep.App;
import com.schmeister.devprep.util.Print;

public class Permutations {

	static Vector<String> permutations = new Vector<String>();

	static void permute(String s, String answer) {
		if (s.length() == 0) {
			permutations.add(answer);
			return;
		}

		for (int i = 0; i < s.length(); i++) {
			char ch = s.charAt(i);
			String left_substr = s.substring(0, i);
			String right_substr = s.substring(i + 1);
			String rest = left_substr + right_substr;
			permute(rest, answer + ch);
		}
	}

	// Driver code
	public static void main() {
		String s = "abc";
		String answer = "";

		permute(s, answer);
		Print.print("Permutations", permutations.toArray());
	}
}
