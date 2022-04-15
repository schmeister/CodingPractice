package com.schmeister.devprep.Arrays;

import com.schmeister.devprep.util.Print;

public class EquilibriumPoint {
	public static void equilibriumPoint() {
		int arr[] = { -7, 1, 5, 2, -4, 3, 0 };
		equilibriumPoint(arr);

		int arr2[] = { 1, 3, 5, 2, 2 };
		equilibriumPoint(arr2);

		int arr3[] = { 1 };
		equilibriumPoint(arr3);
	}

	public static int equilibriumPoint(int arr[]) {
		int n = arr.length;
		int left = 0;
		int right = n - 1;
		int leftsum = 0;
		int rightsum = 0;

		while (left < right) {
			if (leftsum <= rightsum)
				leftsum += arr[left++];
			else
				rightsum += arr[right--];
		}
		Print.print("Equilibrium", arr, left);

		/* return -1 if no equilibrium index is found */
		return -1;
	}
}
