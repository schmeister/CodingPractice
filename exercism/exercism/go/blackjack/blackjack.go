package blackjack

import (
	"strconv"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, err := strconv.Atoi(card)
	if err != nil {
		if card == "ace" {
			value = 11
		} else if card == "king" || card == "queen" || card == "jack" || card == "ten" {
			value = 10
		} else if card == "nine" {
			value = 9
		} else if card == "eight" {
			value = 8
		} else if card == "seven" {
			value = 7
		} else if card == "six" {
			value = 6
		} else if card == "five" {
			value = 5
		} else if card == "four" {
			value = 4
		} else if card == "three" {
			value = 3
		} else if card == "two" {
			value = 2
		}
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	myTotal := ParseCard(card1) + ParseCard(card2)
	dTotal := ParseCard(dealerCard)

	if card1 == card2 && card1 == "ace" {
		return "P"
	} else if myTotal == 21 {
		if dTotal < 10 {
			return "W"
		} else {
			return "S"
		}
	} else if myTotal <= 20 && myTotal >= 17 {
		return "S"
	} else if myTotal <= 16 && myTotal >= 12 {
		if dTotal >= 7 {
			return "H"
		}
		return "S"
	}
	return "H"
}
