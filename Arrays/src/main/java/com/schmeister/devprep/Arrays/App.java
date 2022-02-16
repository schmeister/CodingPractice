package com.schmeister.devprep.Arrays;

import java.util.ArrayList;
import java.util.Arrays;

/**
 * Hello world!
 *
 */
public class App {
	public static void main(String[] args) {
		SubArrayWithGivenSum.getSum();
		CountTriplets.triplets();

		Kadanes.kadanes();

		MissingNumber.getMissingNo();
		MergeSortedArrays.mergeSortedArrays();

		ReArrangeAlternating.rearrangeAlternating();
		Sort012.sort012();

		EquilibriumPoint.equilibriumPoint();

		Leaders.leaders();

		NumPlatforms.numPlatforms();
		
		ReverseInGroups.reverse();
		RainWater.rainWater();
		StockBuySell.stockBuySell();
		ZigZag.zigZag();
	}

	public static void print(String name, int a1, int a2, int a3) {
		System.out.printf("%-18s: %-3d = %3d + %d\n", name, a1, a2, a3);
	}

	public static void print(String name, int[] arr1, int[] arr2, int[] arr3) {
		System.out.printf("%-18s: %-18s - %19s = %s\n", name, Arrays.toString(arr1), Arrays.toString(arr2),
				Arrays.toString(arr3));
	}

	public static void print(String name, int[] arr1, int[] arr2) {
		System.out.printf("%-18s: %-40s - %s\n", name, Arrays.toString(arr1), Arrays.toString(arr2));
	}

	public static void print(String name, int[] arr1, ArrayList<Integer> al) {
		System.out.printf("%-18s: %-40s - %s\n", name, Arrays.toString(arr1), al);
	}

	public static void print(String name, int[] arr1, int val) {
		System.out.printf("%-18s: %-40s - %s\n", name, Arrays.toString(arr1), val);
	}

	public static void print(String name, int val) {
		System.out.printf("%-18s: %d\n", name, val);
	}

	public static void print(String name, int[] arr) {
		System.out.printf("%-18s: %s\n", name, Arrays.toString(arr));
	}
}
