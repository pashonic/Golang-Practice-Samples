package main

import "fmt"

type card struct {
	Type  string
	Suit  string
	Value int
}

type cardSet []card

func createDeck() (deck cardSet) {

	deck = append(deck, card{Type: "Eight", Suit: "Heart", Value: 5})
	deck = append(deck, card{Type: "Seven", Suit: "Heart", Value: 5})
	deck = append(deck, card{Type: "Eight", Suit: "Heart", Value: 5})
	deck = append(deck, card{Type: "Seven", Suit: "Heart", Value: 5})
	deck = append(deck, card{Type: "Seven", Suit: "Heart", Value: 10})
	return deck
}

func get15Score(deck cardSet) int {
	valueSet := []int{}
	for _, v := range deck {
		valueSet = append(valueSet, v.Value)
	}
	return subsetSum(valueSet, 15) * 2
}

func subsetSum(array []int, target int) int {
	total := 0

	for i := 0; i < len(array)-1; i++ {
		total += subsetSumHelp(array, target, []int{array[i]}, i+1, array[i])
	}
	return total
}

func subsetSumHelp(array []int, target int, pass []int, start int, sum int) int {
	if sum == target {
		return 1
	}

	if sum > target {
		return 0
	}

	total := 0
	for i := start; i < len(array); i++ {
		total += subsetSumHelp(array, target, append(pass, array[i]), i+1, sum+array[i])
	}
	return total
}

func main() {
	deck := createDeck()

	fmt.Println(get15Score(deck))
}
