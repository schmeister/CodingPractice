package com.schmeister.devprep.Arrays;

import com.schmeister.devprep.LinkedLists.AddAsNumbers;
import com.schmeister.devprep.LinkedLists.BalancedBrackets;
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
import com.schmeister.devprep.Numbers.Prime;
import com.schmeister.devprep.StackAndQueues.NextGreatestElement;
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
import com.schmeister.devprep.Trees.Nary;

import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;

/**
 * Unit test for simple App.
 */
public class AppTest 
    extends TestCase
{
    /**
     * Create the test case
     *
     * @param testName name of the test case
     */
    public AppTest( String testName )
    {
        super( testName );
    }

    /**
     * @return the suite of tests being tested
     */
    public static Test suite()
    {
        return new TestSuite( AppTest.class );
    }

    /**
     * Rigourous Test :-)
     */
    public void testApp()
    {
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
		BalancedBrackets.balancedBrackets();
		
		// Stack and Queue
		NextGreatestElement.nextGreatestElement();
		
		// Numbers
		Prime.prime();
		
		// Trees
		Nary.nary();

//        assertTrue( true );
    }
}
