package random_test

import (
	"atomicgo.dev/random"
	"fmt"
)

func ExampleBool() {
	random.Seed(1337) // Set seed for deterministic output, not required

	if random.Bool() {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}

	// Output: Heads
}

func ExampleBoolSlice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.BoolSlice(5)
	fmt.Println(s)

	// Output: [true true false false true]
}
