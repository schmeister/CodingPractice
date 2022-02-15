package com.schmeister.devprep.Arrays;

public class Kadanes {

	public static void kadanes() {
		int arr[] = { 1, 2, 3, -2, 5 };
		System.out.println(kadanes(arr));
		System.out.println(kandaneForMaxSubArrayForNegativ(arr));
		
		int arr2[] = { -1, -2, -3, -2, -5 };
		System.out.println(kadanes(arr2));
		System.out.println(kandaneForMaxSubArrayForNegativ(arr2));		
	}

	public static int kadanes(int[] arr) {
		int maxEndHere = 0;
		int maxSoFar = 0;
		for (int i = 0; i < arr.length; i++) {
			maxEndHere += arr[i];
			if (maxEndHere < 0) {
				maxEndHere = 0;
			}
			if (maxSoFar < maxEndHere) {
				maxSoFar = maxEndHere;
			}
		}
		return maxSoFar;
	}

	public static int kandaneForMaxSubArrayForNegativ(int[] arr) {
		int maxEndHere = arr[0];
		int maxSoFar = arr[0];
		for (int i = 1; i < arr.length; i++) {
			maxEndHere = Math.max(arr[i], maxEndHere + arr[i]);
			maxSoFar = Math.max(maxSoFar, maxEndHere);
		}
		return maxSoFar;
	}
}
