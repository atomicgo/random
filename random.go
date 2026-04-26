package random

import (
	"math/rand"
	"sync"
)

var (
	source *lockedSource //nolint:gochecknoglobals // Global random source
	r      *rand.Rand    //nolint:gochecknoglobals,varnamelen // Global random generator
)

func init() {
	source = &lockedSource{src: rand.NewSource(rand.Int63())} //nolint:gosec // Unsecure random is fine for this package
	r = rand.New(source)                                      //nolint:gosec // Unsecure random is fine for this package
}

// Seed sets the seed of the random generator.
// By default, the seed is already randomized on package initialization.
// You only need to call this function if you need a specific seed.
func Seed(seed int64) {
	source.Seed(seed)
}

type lockedSource struct {
	mu  sync.Mutex
	src rand.Source
}

func (s *lockedSource) Int63() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.src.Int63()
}

func (s *lockedSource) Seed(seed int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.src.Seed(seed)
}

func (s *lockedSource) Uint64() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.src.(rand.Source64).Uint64() //nolint:forcetypeassert // rand.NewSource returns a Source64.
}
