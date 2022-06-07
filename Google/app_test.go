package main

import (
	"schmeister/google/arbitraryPointer"
	"schmeister/google/balancedParends"
	"schmeister/google/circularGraph"
	"schmeister/google/duplicateSubTree"
	"schmeister/google/findBoundaries"
	"schmeister/google/jumpingNumbers"
	"schmeister/google/kClosestNumbers"
	"schmeister/google/kthLargest"
	"schmeister/google/lruCache"
	"schmeister/google/palindrome"
	"schmeister/google/pathsForASum"
	"schmeister/google/subArraySum"
	"schmeister/google/wordBreak"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMain(t *testing.T) {
	kthLargest.KthLargestElement()
	arbitraryPointer.ArbitraryPointer()

	lruCache.LRUCache()

	circularGraph.CircularGraph()

	findBoundaries.FindIndex()

	palindrome.Palindrome()

	kClosestNumbers.FindClosestElement()

	pathsForASum.PathsForASum()

	subArraySum.SubArraySum()

	balancedParends.Balanced()

	jumpingNumbers.ListJumpingNumbers(1005)

	wordBreak.WordBreak()

	duplicateSubTree.DuplicateSubTree()
}
