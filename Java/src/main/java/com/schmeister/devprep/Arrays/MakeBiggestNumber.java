package com.schmeister.devprep.Arrays;

import java.util.Collections;
import java.util.Comparator;
import java.util.Iterator;
import java.util.Vector;

public class MakeBiggestNumber {

	// The main function that prints the
	// arrangement with the largest value.
	// The function accepts a vector of strings
	static void printLargest(Vector<String> arr) {

		Collections.sort(arr, new Comparator<String>() {
			// A comparison function which is used by
			// sort() in printLargest()
			public int compare(String X, String Y) {

				// first append Y at the end of X
				String XY = X + Y;

				// then append X at the end of Y
				String YX = Y + X;

				// Now see which of the two
				// formed numbers
				// is greater
				return XY.compareTo(YX) > 0 ? -1 : 1;
			}
		});

		Iterator<String> it = arr.iterator();

		while (it.hasNext())
			System.out.print(it.next());
		System.out.println();
	}

	// Driver code
	public static void makeBiggestNumber() {

		Vector<String> arr;
		arr = new Vector<String>();

		// output should be 6054854654
		arr.add("60");
		arr.add("54");
		arr.add("546");
		arr.add("548");
		printLargest(arr);
	}
}