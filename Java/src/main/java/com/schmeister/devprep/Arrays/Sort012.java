package com.schmeister.devprep.Arrays;

import com.schmeister.devprep.util.Print;

public class Sort012 {
	public static void sort012() {
		int arr[] = { 0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1 };
		int arr2[] = sort012(arr);
		Print.print("Sort012: ", arr, arr2);
	}

	static int[] sort012(int a[]) {
		int arr_size = a.length;
		int[] b = new int[arr_size];

		int bottom = 0;
		int top = arr_size - 1;

		for (int x = 0; x < arr_size - 1; x++) {
			int val = a[x];
			if (val == 0)
				b[bottom++] = val;
			else if (val == 2)
				b[top--] = val;
		}
		for (int x = bottom; x <= top; x++) {
			b[x] = 1;
		}
		return b;
	}
}
