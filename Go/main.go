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

	// Linked Lists
	linkedLists.ReverseInGroups()

	// Mathematical
	mathematical.PrintThePattern(3)
	mathematical.MultiplicationTable(5)
	mathematical.SeriesAP(2,3,4)
	mathematical.SeriesAP(1,2,10)

	mathematical.SeriesGP(2,2,4)
	mathematical.SeriesGP(4,3,3)

	mathematical.ClosestNumber(13,4)
	mathematical.ClosestNumber(-15,6)

}
