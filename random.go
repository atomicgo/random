package random

import (
	"math/rand"
)

var r *rand.Rand //nolint: gochecknoglobals, varnamelen // Global random generator

func init() {
	r = rand.New(rand.NewSource(rand.Int63())) //nolint:all // Unsecure random is fine for this package
}

// Seed sets the seed of the random generator.
// By default, the seed is already randomized on package initialization.
// You only need to call this function if you need a specific seed.
func Seed(seed int64) {
	r.Seed(seed)
}
