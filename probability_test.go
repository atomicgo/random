package random_test

import (
	"fmt"

	"atomicgo.dev/random"
)

func ExampleProbability() {
	random.Seed(1337) // Set seed for deterministic output, not required

	// We are using a pronged coin with a 80% chance of landing on Heads.
	if random.Probability(0.8) {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}

	// Output: Heads
}
