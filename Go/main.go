package main

import (
	"github.com/schmeister/codingPrep/arrays"
	"github.com/schmeister/codingPrep/linkedLists"
	"github.com/schmeister/codingPrep/mathematical"
	"github.com/schmeister/codingPrep/trees"
)

func main() {

	// Array
	arrays.SubArraySum()

	// Tree
	trees.Nary()
	trees.BFS()

	// Linked Lists
	linkedLists.ReverseInGroups()

	// Mathematical
	mathematical.PrintThePattern(3)
	mathematical.MultiplicationTable(5)
	mathematical.SeriesAP(2, 3, 4)
	mathematical.SeriesAP(1, 2, 10)

	mathematical.SeriesGP(2, 2, 4)
	mathematical.SeriesGP(4, 3, 3)

	mathematical.ClosestNumber(13, 4)
	mathematical.ClosestNumber(-15, 6)
	mathematical.RunArmstrong(371)
	mathematical.RunArmstrong(375)
	mathematical.SumisPalindrome(56)
	mathematical.SumisPalindrome(579683547939999)
	mathematical.SumisPalindrome(579683547839999)

	mathematical.Reverse(1)
	mathematical.Reverse(221)
	mathematical.Reverse(2213)

	mathematical.KthDigit(5, 2, 2)
	mathematical.KthDigit(5, 9, 2)
	mathematical.KthDigit(5, 9, 4)

	mathematical.BinaryToDecimal("1001")
	mathematical.BinaryToDecimal("10011")

	mathematical.JumpingNumbers(20)
	mathematical.JumpingNumbers(105)

	mathematical.GCD(5, 10)
	mathematical.GCD(5, 11)
}
