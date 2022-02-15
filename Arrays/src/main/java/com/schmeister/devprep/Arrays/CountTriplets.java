package com.schmeister.devprep.Arrays;

import java.util.Arrays;

public class CountTriplets {

	public static void triplets()
    {
        int arr[] = { 5, 32, 1, 7, 10, 50, 19, 21, 2, 18, 2 };
        int n = arr.length;
        findTriplet(arr, n);
    }
	
	private static void findTriplet(int arr[], int n) {
		// sort the array
		Arrays.sort(arr);

		// for every element in arr
		// check if a pair exist(in array) whose
		// sum is equal to arr element
		for (int i = n - 1; i >= 0; i--) {
			int j = 0;
			int k = i - 1;
			while (j < k) {
				if (arr[i] == arr[j] + arr[k]) {

					// pair found
					System.out.println("numbers are " + arr[i] + " " + arr[j] + " " + arr[k]);

					return;
				} else if (arr[i] > arr[j] + arr[k])
					j += 1;
				else
					k -= 1;
			}
		}
	}
}
