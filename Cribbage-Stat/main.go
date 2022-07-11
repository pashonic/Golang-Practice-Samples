package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"
)

func getPercentString(numerator uint64, denominator uint64) string {
	percent := (float64(numerator) / float64(denominator)) * 100
	return fmt.Sprintf("%.2f%%", percent)
}

type sampleResult struct {
	total15Score    uint64
	totalRunScore   uint64
	totalPairScore  uint64
	totalFlushScore uint64
	totalNobScore   uint64
}

func dealerSimulation(c chan sampleResult) {

	// Create a deck of cards
	deck := cribbage.CreateDeck()

	// Deal 6 cards
	dealtHand := deck.Deal(6)

	// Get the best 4 cards that result in highest score
	selectedHand := dealtHand.GetBestCards(4)

	// Deal cut card
	var cutCard cribbage.Card = deck.Deal(1)[0]

	// Add cut card to deck
	dealtHand = append(selectedHand, cutCard)

	// Add 2 cards from deck
	finalHand := append(dealtHand, deck.Deal(2)...)

	sampleResult := sampleResult{
		total15Score:    uint64(finalHand.GetFifteenScore()),
		totalRunScore:   uint64(finalHand.GetRunScore()),
		totalPairScore:  uint64(finalHand.GetPairScore()),
		totalFlushScore: uint64(finalHand.GetFlushScore()),
		totalNobScore:   uint64(finalHand.GetNobScore(&cutCard)),
	}

	c <- sampleResult
}

func main() {

	var total15Score uint64 = 0
	var totalRunScore uint64 = 0
	var totalPairScore uint64 = 0
	var totalFlushScore uint64 = 0
	var totalNobScore uint64 = 0
	var sampleCount int = 10000000
	var sampleThreads int = 20
	var sampleCycles int = sampleCount / sampleThreads

	// Run simulated cribbage hand.
	start := time.Now()
	finalSampleCount := 0
	for cycle := 0; cycle < sampleCycles; cycle++ {
		ch := make(chan sampleResult)

		for i := 0; i < sampleThreads; i++ {
			finalSampleCount++
			go dealerSimulation(ch)
		}

		for i := 0; i < sampleThreads; i++ {
			sampleResult := <-ch
			total15Score += sampleResult.total15Score
			totalRunScore += sampleResult.totalRunScore
			totalPairScore += sampleResult.totalPairScore
			totalFlushScore += sampleResult.totalFlushScore
			totalNobScore += sampleResult.totalNobScore
		}
	}
	duration := time.Since(start)

	// Add up totals
	totalScore := total15Score + totalRunScore + totalPairScore + totalFlushScore + totalNobScore

	// Print result
	varMap := []struct {
		name  string
		value uint64
	}{
		{"15 Combo", total15Score},
		{"Run", totalRunScore},
		{"Pairs", totalPairScore},
		{"Flush", totalFlushScore},
		{"Nob Score", totalNobScore},
	}

	spacer := fmt.Sprintf("+%s+\n", strings.Repeat("-", 68))
	fmt.Println("Threads: ", sampleThreads)
	fmt.Println("Duration: ", duration.Seconds())
	fmt.Println("Sample Count: ", finalSampleCount)
	fmt.Printf(spacer)
	fmt.Printf("| %-20s | %-20s | %-20s |\n", "Score Type", "Points", "Percentage")
	fmt.Printf(spacer)
	for _, varItem := range varMap {
		u64String := fmt.Sprintf("%v", varItem.value)
		fmt.Printf("| %-20s | %-20s | %-20s |\n", varItem.name, u64String, getPercentString(varItem.value, totalScore))
		fmt.Printf(spacer)
	}
}
