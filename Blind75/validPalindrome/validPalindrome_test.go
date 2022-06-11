package validPalindrome

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T)  {
	fmt.Println(ValidPalindrome(""))
	fmt.Println(ValidPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(ValidPalindrome("race a car"))
}
