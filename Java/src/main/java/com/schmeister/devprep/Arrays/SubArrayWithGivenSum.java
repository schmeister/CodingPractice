package com.schmeister.devprep.Arrays;

import java.util.Arrays;

import com.schmeister.devprep.App;
import com.schmeister.devprep.util.Print;

public class SubArrayWithGivenSum {

	public static void getSum() {
		int s = 12;
		int a[] = { 1, 2, 3, 7, 5 };
		subArraySum(a, s);

		s = 16;
		int a2[] = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
		subArraySum(a2, s);
	}

	private static int subArraySum(int arr[], int sum) {
		int n = arr.length;
		int curr_sum = arr[0];
		int start = 0;

		// Sliding window:
		// Add higher element in array if low
		// Subtract lower element in array if high
		for (int i = 1; i <= n; i++) {
			while (curr_sum > sum && start < i - 1) {
				curr_sum = curr_sum - arr[start];
				start++;
			}
			if (curr_sum == sum) {
				int p = i - 1;
				Print.print("Sum: ", Arrays.copyOfRange(arr, start, p + 1), sum);
				return 1;
			}

			if (i < n)
				curr_sum = curr_sum + arr[i];
		}

		System.out.println("No subarray found");
		return 0;
	}
}
