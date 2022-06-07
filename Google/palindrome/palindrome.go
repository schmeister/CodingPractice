package palindrome

import "fmt"

func isPalindrome(str string) bool {

	chars := []byte(str)
	max := len(chars) - 1

	for x := 0; x < max; x++ {
		if chars[x] != chars[max-x] {
			return false
		}
	}
	return true
}

func Palindrome(){
	fmt.Println(isPalindrome("1234321"))
	fmt.Println(isPalindrome("12344321"))
	fmt.Println(isPalindrome("12345321"))
}