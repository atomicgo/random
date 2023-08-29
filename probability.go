package random

// Probability returns true with a probability of p.
// 1 = 100%; 0.5 = 50%; 0 = 0%
// If p is less than 0 or greater than 1, Probability panics.
func Probability(p float64) bool {
	if p < 0 || p > 1 {
		panic("Random probability must not be less than 0 or greater than 1")
	}

	return Float64(0, 1) <= p
}
