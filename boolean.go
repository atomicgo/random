package random

// Bool returns a random boolean.
func Bool() bool {
	return r.Intn(2) == 0
}

// BoolSlice returns a slice of random booleans.
// If count is less than 0, BoolSlice will panic.
func BoolSlice(count int) []bool {
	if count < 0 {
		panic("Random bool slice count must not be less than 0")
	}

	s := make([]bool, count)
	for i := range s {
		s[i] = Bool()
	}

	return s
}
