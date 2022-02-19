package com.schmeister.devprep;

import java.util.ArrayList;
import java.util.Arrays;

import com.schmeister.devprep.Arrays.*;
import com.schmeister.devprep.LinkedLists.AddAsNumbers;
import com.schmeister.devprep.LinkedLists.DetectLoop;
import com.schmeister.devprep.LinkedLists.FindMiddle;
import com.schmeister.devprep.LinkedLists.IntersectingList;
import com.schmeister.devprep.LinkedLists.Merge2Sorted;
import com.schmeister.devprep.LinkedLists.NthFromEnd;
import com.schmeister.devprep.LinkedLists.PairwiseSwap;
import com.schmeister.devprep.LinkedLists.Palindrome;
import com.schmeister.devprep.LinkedLists.Queue;
import com.schmeister.devprep.LinkedLists.RemoveLoop;
import com.schmeister.devprep.LinkedLists.RemoveNode;
import com.schmeister.devprep.LinkedLists.Reverse;
import com.schmeister.devprep.LinkedLists.ReverseLLInGroups;
import com.schmeister.devprep.LinkedLists.Rotate;
import com.schmeister.devprep.LinkedLists.Sort;
import com.schmeister.devprep.LinkedLists.Stack;
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

		// Linked Lists
		FindMiddle.findMiddle();
		Reverse.reverse();
		Rotate.rotate();
		ReverseLLInGroups.reverseInGroups();
		IntersectingList.intersectingList();
		DetectLoop.detectLoop();
		RemoveLoop.removeLoop();
		NthFromEnd.nthFromEnd();
		Merge2Sorted.merge2Sorted();
		PairwiseSwap.pairwiseSwap();
		AddAsNumbers.addAsNumbers();
		Palindrome.palindrome();
		Queue.queue();
		Stack.stack();
		Sort.sortLinkedList();
		RemoveNode.removenode();
	}
}
