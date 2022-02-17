package com.schmeister.devprep.Arrays;

import java.util.Arrays;

import com.schmeister.devprep.App;
import com.schmeister.devprep.Print;

public class ReverseInGroups {

	public static void reverse()
    {
         
        int arr[] = {1, 2, 3, 4, 5, 6, 7, 8};
        int k = 6;
     
        reverse(arr, k);
     
        Print.print("Reverse", arr);
    }
	public static void reverse(int[] arr, int size) {
		int len = arr.length;
		for (int i = 0; i < len; i += size) {
			int left = i;

			// to handle case when k is not multiple
			// of n
			int right = Math.min(i + size - 1, len - 1);
			int temp;

			// reverse the sub-array [left, right]
			while (left < right) {
				temp = arr[left];
				arr[left] = arr[right];
				arr[right] = temp;
				left += 1;
				right -= 1;
			}
		}
	}
}
