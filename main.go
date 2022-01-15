package main

import (
	"fmt"

	"github.com/mitchallen/coin"
)

func demoFlip() {

	// Demo Flip()

	m := make(map[bool]int)

	m[true] = 0
	m[false] = 0

	for i := 0; i < 100; i++ {
		m[coin.Flip()]++
	}

	fmt.Print("[Flip]: ")
	fmt.Println(m)
}

func demoWeighted() {

	// Demo Weighted(w)

	m := make(map[bool]int)

	m[true] = 0
	m[false] = 0

	var testWeight float32 = 0.75

	for i := 0; i < 100; i++ {
		m[coin.Weighted(testWeight)]++
	}

	fmt.Printf("[Weight(%.2f)]: ", testWeight)
	fmt.Println(m)
}

func main() {
	fmt.Println()
	demoFlip()
	demoWeighted()
	fmt.Println()
}
