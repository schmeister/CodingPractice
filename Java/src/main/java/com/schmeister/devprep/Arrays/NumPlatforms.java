package com.schmeister.devprep.Arrays;

import java.util.Arrays;

import com.schmeister.devprep.util.Print;

public class NumPlatforms {
	public static void numPlatforms() {
		int arr[] = { 900, 940, 950, 1100, 1500, 1800 };
		int dep[] = { 910, 1200, 1120, 1130, 1900, 2000 };
		Print.print("Platforms", findPlatform(arr, dep));
	}

	static int findPlatform(int arr[], int dep[]) {
		int arrL = arr.length;
		int depL = dep.length;

		Arrays.sort(arr);
		Arrays.sort(dep);

		int plat_needed = 0, result = 1;
		int i = 0, j = 0;

		while (i < arrL && j < depL) {

			// Arrival
			if (arr[i] <= dep[j]) {
				plat_needed++;

//				System.out.println(arr[i] + "\tarr\t" + i + "\t" + plat_needed);

				i++;
			}
			// Departure
			else if (arr[i] > dep[j]) {
				plat_needed--;

//				System.out.println(dep[j] + "\tdep\t" + j + "\t" + plat_needed);

				j++;
			}

			if (plat_needed > result)
				result = plat_needed;
		}
		return result;
	}
}
