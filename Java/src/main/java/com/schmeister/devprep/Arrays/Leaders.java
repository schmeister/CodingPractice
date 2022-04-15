package com.schmeister.devprep.Arrays;

import java.util.ArrayList;
import java.util.Arrays;

import com.schmeister.devprep.App;
import com.schmeister.devprep.util.Print;

public class Leaders {

	public static void leaders() {
		int arr[] = new int[] { 16, 17, 4, 3, 5, 2 };
		leaders(arr);
	}

	public static void leaders(int[] arr) {
		ArrayList<Integer> al = new ArrayList<Integer>();
		int len = arr.length;
		int max = Integer.MIN_VALUE;
		for (int x = len - 1; x >= 0; x--) {
			if (arr[x] > max) {
				al.add(arr[x]);
				max = arr[x];
			}
		}

		Print.print("Leaders", arr,  al);
	}
}
