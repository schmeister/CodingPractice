package com.schmeister.devprep.Arrays;

import com.schmeister.devprep.util.Print;

public class Kadanes {

	public static void kadanes() {
		int arr[] = { 1, 2, 3, -2, 5 };
		Print.print("Kadanes",arr,kadanes(arr));
		Print.print("Kadanes",arr,kandaneForMaxSubArrayForNegativ(arr));
		
		int arr2[] = { -1, -2, -3, -2, -5 };
		Print.print("Kadanes",arr,kadanes(arr2));
		Print.print("Kadanes",arr,kandaneForMaxSubArrayForNegativ(arr2));
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
