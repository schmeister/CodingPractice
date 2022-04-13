package com.schmeister.devprep;

import java.util.ArrayList;
import java.util.Arrays;

public class Print {

	public static void print(String name, boolean a1) {
		System.out.printf("%-18s: %b\n", name, a1);
	}

	public static void print(String name, String str, boolean a1) {
		System.out.printf("%-18s: %-20s %b\n", name, str, a1);
	}

	public static void print(String name, int i, boolean a1) {
		System.out.printf("%-18s: %-20d %b\n", name, i, a1);
	}

	public static void print(String name, int i, int i2) {
		System.out.printf("%-18s: %-20d %d\n", name, i, i2);
	}

	public static void print(String name, String str, String str2, boolean a1) {
		System.out.printf("%-18s: %-10s %-10s %b\n", name, str, str2, a1);
	}

	public static void print(String name, String str, int a1) {
		System.out.printf("%-18s: %-20s %d\n", name, str, a1);
	}

	public static void print(String name, String str, String str2) {
		System.out.printf("%-18s: %-20s %s\n", name, str, str2);
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

	public static void print(String name, String val) {
		System.out.printf("%-18s: %s\n", name, val);
	}

	public static void print(String name, int[] arr) {
		System.out.printf("%-18s: %s\n", name, Arrays.toString(arr));
	}

	public static void print(String name, String[] arr) {
		System.out.printf("%-18s: %s\n", name, Arrays.toString(arr));
	}

	public static void print(String name, Object[] arr) {
		System.out.printf("%-18s: %s\n", name, Arrays.toString(arr));
	}

	public static void print(String name, boolean[][] arr) {
		System.out.printf("%-18s\n", name);
		print(arr);
	}

	public static void print(boolean table[][]) {
		for (int i = 0; i < table.length; i++) {
			for (int j = 0; j < table[i].length; j++) {
				if (table[i][j])
					System.out.printf("%6s", table[i][j]);
				else
					System.out.printf("%6s", "");
			}
			System.out.println();
		}
		System.out.println();
	}

	public static void print(int table[][]) {
		for (int i = 0; i < table.length; i++) {
			for (int j = 0; j < table[i].length; j++) {
				System.out.printf("%6s", table[i][j]);
			}
			System.out.println();
		}
		System.out.println();
	}
}
