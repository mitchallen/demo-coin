package main

import (
	"fmt"

	"github.com/mitchallen/coin"
)

func demoFlip() {

	// Demo Flip()

	a := make([]bool, 100)

	heads := 0
	tails := 0

	for i, _ := range a {
		a[i] = coin.Flip()
		if a[i] {
			heads++
		} else {
			tails++
		}
	}

	fmt.Printf("[Flip()] Heads: %d Tails: %d \n", heads, tails)
}

func demoWeighted() {

	// Demo Weighted(w)

	a := make([]bool, 100)

	heads := 0
	tails := 0

	var testWeight float32 = 0.75

	for i, _ := range a {
		a[i] = coin.Weighted(testWeight)
		if a[i] {
			heads++
		} else {
			tails++
		}
	}

	fmt.Printf("[Weight(%.2f)] Heads: %d Tails: %d \n", testWeight, heads, tails)
}

func main() {
	demoFlip()
	demoWeighted()
}
