package com.schmeister.devprep.Arrays;

public class SubArrayWithGivenSum {

	public static void getSum() {
		int n = 5;
		int s = 12;
		int a[] = { 1, 2, 3, 7, 5 };
		System.out.println(subArraySum(a, n, s));

		n = 10;
		s = 15;
		int a2[] = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
		System.out.println(subArraySum(a2, n, s));
	}

	private static int subArraySum(int arr[], int n, int sum) {
		int curr_sum = arr[0];
		int start = 0;

		// Pick a starting point
		for (int i = 1; i <= n; i++) {
			// If curr_sum exceeds the sum,
			// then remove the starting elements
			while (curr_sum > sum && start < i - 1) {
				curr_sum = curr_sum - arr[start];
				start++;
			}

			// If curr_sum becomes equal to sum,
			// then return true
			if (curr_sum == sum) {
				int p = i - 1;
				System.out.println("Sum found between indexes " + start + " and " + p);
				return 1;
			}

			// Add this element to curr_sum
			if (i < n)
				curr_sum = curr_sum + arr[i];
		}

		System.out.println("No subarray found");
		return 0;
	}
}
