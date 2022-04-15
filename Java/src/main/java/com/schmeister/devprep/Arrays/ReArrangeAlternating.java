package com.schmeister.devprep.Arrays;

import java.util.Arrays;

import com.schmeister.devprep.App;
import com.schmeister.devprep.util.Print;

public class ReArrangeAlternating {
	// Driver code
	public static void rearrangeAlternating() {
		int arr[] = new int[] { 1, 2, 3, 4, 5, 6 };
		int temp[] = rearrangeAlternating(arr);

		Print.print("Alternating: ", arr, temp);
	}

	static int[] rearrangeAlternating(int[] arr) {
		int n = arr.length;
		// Auxiliary array to hold modified array
		int temp[] = arr.clone();

		// Indexes of smallest and largest elements
		// from remaining array.
		int small = 0, large = n - 1;

		// To indicate whether we need to copy remaining
		// largest or remaining smallest at next position
		boolean flag = true;

		// Store result in temp[]
		for (int i = 0; i < n; i++) {
			if (flag)
				temp[i] = arr[large--];
			else
				temp[i] = arr[small++];

			flag = !flag;
		}
		return temp;
	}
}
