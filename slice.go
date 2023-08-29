package random

// Entry returns a random element from a slice.
func Entry[T any](s []T) T {
	return s[Int(0, len(s)-1)]
}

// Entries returns a slice of n random elements from a slice.
// If n is less than 0, Entries will panic.
func Entries[T any](s []T, n int) []T {
	if n < 0 {
		panic("Random entries count must not be less than 0")
	}

	p := make([]T, n)
	for i := range p {
		p[i] = Entry(s)
	}

	return p
}

// EntriesUnique returns a slice of n unique random elements from a slice.
// If n is less than 0, EntriesUnique will panic.
// If n is greater than the length of s, EntriesUnique will panic.
func EntriesUnique[T any](s []T, n int) []T {
	if n < 0 {
		panic("Random entries unique count must not be less than 0")
	}

	if n > len(s) {
		panic("Random entries unique count must not be greater than slice length")
	}

	randomIndexes := IntSliceUnique(0, len(s)-1, n)

	p := make([]T, n)
	for i, v := range randomIndexes {
		p[i] = s[v]
	}

	return p
}

// Shuffle returns a shuffled copy of a slice.
func Shuffle[T any](s []T) []T {
	r.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
