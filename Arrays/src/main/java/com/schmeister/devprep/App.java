package com.schmeister.devprep;

import java.util.ArrayList;
import java.util.Arrays;

import com.schmeister.devprep.Arrays.*;
import com.schmeister.devprep.Strings.Atoi;
import com.schmeister.devprep.Strings.FormPalindrome;
import com.schmeister.devprep.Strings.IsAnagram;
import com.schmeister.devprep.Strings.IsPalindrome;
import com.schmeister.devprep.Strings.IsRotated;
import com.schmeister.devprep.Strings.LongestPalindrome;
import com.schmeister.devprep.Strings.LongestPrefix;
import com.schmeister.devprep.Strings.LongestUniqueSubStr;
import com.schmeister.devprep.Strings.Permutations;
import com.schmeister.devprep.Strings.RecursiveRemoveDuplicates;
import com.schmeister.devprep.Strings.RemoveDuplicates;
import com.schmeister.devprep.Strings.ReverseWords;
import com.schmeister.devprep.Strings.RomanToNumber;
import com.schmeister.devprep.Strings.Strstr;

/**
 * Hello world!
 *
 */
public class App {
	public static void main(String[] args) {
		// Arrays
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
		MakeBiggestNumber.makeBiggestNumber();
		
		// Strings
		ReverseWords.reverseWords();
		Permutations.main();
		LongestPalindrome.longest();
		IsPalindrome.isPalindrome();
		RecursiveRemoveDuplicates.recursiveRemoveDups();
		IsRotated.isRotated();
		RomanToNumber.romanToNumber();
		IsAnagram.isAnagram();
		RemoveDuplicates.removeDuplicates();
		FormPalindrome.formPalindrome();
		LongestUniqueSubStr.longestUnique();
		Atoi.atoi();
		Strstr.strstr();
		LongestPrefix.longestPrefix();
	}
}
