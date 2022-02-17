package com.schmeister.devprep.Arrays;

import java.util.Arrays;

import com.schmeister.devprep.App;
import com.schmeister.devprep.Print;

public class MissingNumber {
	// Driver code
	public static void getMissingNo() {
		int a[] = { 1, 2, 4, 5, 6 };
		int n = a.length + 1;
		int miss = getMissingNo(a, n);

		Print.print("MissingNumber", a, miss);
	}

	static int getMissingNo(int a[], int n) {
		int n_elements_sum = n * (n + 1) / 2;
		int sum = 0;

		for (int i = 0; i < n - 1; i++)
			sum += a[i];

		return n_elements_sum - sum;
	}
}