package random_test

import (
	"fmt"

	"fmt"

	"atomicgo.dev/random"
)

func ExampleEntry() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := []int{1, 2, 3, 4, 5}
	fmt.Print(random.Entry(s))

	// Output: 4
}

func ExampleEntries() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := []int{1, 2, 3, 4, 5}
	fmt.Print(random.Entries(s, 3))

	// Output: [4 4 5]
}

func ExampleEntriesUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := []int{1, 2, 3, 4, 5}
	fmt.Print(random.EntriesUnique(s, 3))

	// Output: [4 5 2]
}

func ExampleShuffle() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := []int{1, 2, 3, 4, 5}
	fmt.Print(random.Shuffle(s))

	// Output: [1 3 5 2 4]
}
