package validAnagram

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T){
	str1 := "anagram"
	str2 := "naagram"
	str3 := "naagramm"
	str4 := "mnaagram"

	fmt.Println(ValidAnagram(str1,str2))
	fmt.Println(ValidAnagram(str1,str3))
	fmt.Println(ValidAnagram(str4,str3))
}