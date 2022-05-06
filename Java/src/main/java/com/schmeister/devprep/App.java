package com.schmeister.devprep;

import java.util.ArrayList;

class Class {
	Person[] students;
	Person teacher;
	Subject subject;
}

class Subject {
	String name;
	String description;
}

class Address {
	String street;
	String zip;
	String apt;
	String city;
	String state;
	String phone;
}

class Semester {
	String semesterID;
	Subject[] classes;
}

class Person {
	String name;
	String dob;
	Address[] addresses;
	Semester[] semesters;
}

class School {
	Class[] classes;
	Person[] teachers;
	Person[] students;
}

public class App {
	public static void main(String[] args) {
//		System.out.println(isPalindrome("12321"));
//		System.out.println(isPalindrome("12322"));
//		System.out.println(isPalindrome("124421"));

		System.out.println(longestPalindrome("abaxyzzyx"));

	}

	public static String longestPalindrome(String str) {
		int len = str.length();

		String sub = "";
		int subLen = sub.length();
		for (int x = 0; x < len; x++) {
			for (int y = x + 1; y < len; y++) {
				String s = str.substring(x, y + 1);
				boolean isPal = isPalindrome(s);
				if (isPal && s.length() > subLen) {
					sub = s;
					subLen = s.length();
				}
			}
		}
		return sub;
	}

	public static boolean isPalindrome(String str) {
		char[] chars = str.toCharArray();

		int len = chars.length;
		int middle = len / 2;

		for (int x = 0; x < middle; x++) {
			char c1 = chars[x];
			char c2 = chars[len - x - 1];

			if (c1 != c2)
				return false;
		}
		return true;
	}
}
