package com.schmeister.devprep.Strings;

import java.util.Arrays;

import com.schmeister.devprep.App;
import com.schmeister.devprep.util.Print;

public class ReverseWords {

	public static void reverseWords() {
		reverseWords("I like this program very much");
		reverseWords("I like this program very much again");
	}

	public static String[] reverseWords(String str) {
		String[] words = str.split(" ");
		int len = words.length;

		for (int x = 0; x < len / 2; x++) {
			String temp = words[x];
			words[x] = words[len - x - 1];
			words[len - x - 1] = temp;
		}
		Print.print("ReverseWords", words);
		
		return words;
	}
}
