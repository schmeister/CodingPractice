package testDome

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(r rune) {
	buf := make([]byte, 1)
	_ = utf8.EncodeRune(buf, r)
	_, err := strconv.Atoi(string(buf))
	if err == nil{
		n.input = n.input + string(buf)
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}

func Rune() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('7')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
