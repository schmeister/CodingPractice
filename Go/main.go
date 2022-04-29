package main

import (
	"fmt"

	"github.com/schmeister/codingPrep/competency/testDome"
	"github.com/schmeister/codingPrep/competency/tutorialEdge"
	"github.com/schmeister/codingPrep/interfaces"
)

func main() {
	interfaces.Vowels()
	interfaces.Geometry()

	fmt.Println(testDome.QuadFindRoots(2, 10, 8))
	testDome.Rune()
	fmt.Println(testDome.UniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))

	fmt.Println("Hello World")


	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := tutorialEdge.GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
	tutorialEdge.Interface()
	tutorialEdge.Sorting()
}
