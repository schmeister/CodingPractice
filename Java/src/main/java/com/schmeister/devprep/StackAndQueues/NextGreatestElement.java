package com.schmeister.devprep.StackAndQueues;

import java.util.Arrays;
import java.util.Stack;

import com.schmeister.devprep.util.Print;

public class NextGreatestElement {

	static int arr[] = { 11, 10, 9, 12, 13, 21, 3 };

	/*
	 * prints element and NGE pair for all elements of arr[] of size n
	 */
	public static String printNGE() {
		Stack<Integer> s = new Stack<>();
		int nge[] = new int[arr.length];

		// iterate for rest of the elements
		for (int i = arr.length - 1; i >= 0; i--) {
			/*
			 * if stack is not empty, then pop an element from stack. If the popped element
			 * is smaller than next, then a) print the pair b) keep popping while elements
			 * are smaller and stack is not empty
			 */
			if (!s.empty()) {
				while (!s.empty() && s.peek() <= arr[i]) {
					s.pop();
				}
			}
			nge[i] = s.empty() ? -1 : s.peek();
			s.push(arr[i]);
		}
		StringBuilder sb = new StringBuilder();
		for (int i = 0; i < arr.length; i++) {
			sb.append(String.format("%-4d", nge[i]));
		}
		return sb.toString();
	}

	static public void nextGreatestElement() {
		// NextGreaterElement nge = new
		// NextGreaterElement();
		Print.print("nextGreatestElement", Arrays.toString(arr));
		Print.print("nextGreatestElement", printNGE());
	}
}
