package random

// StringSet constants are sets of characters that can be used in random strings.
const (
	StringSetLowercase = "abcdefghijklmnopqrstuvwxyz"
	StringSetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	StringSetNumbers   = "0123456789"
	StringSetSymbols   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	StringSetAll       = StringSetLowercase + StringSetUppercase + StringSetNumbers + StringSetSymbols
)

// String returns a random string of length l, using the characters in the set.
// If l is less than 0, String panics.
// If the set is empty, or l is 0, String returns an empty string.
func String(l int, set string) string {
	if l < 0 {
		panic("Random string length must not be less than 0")
	}

	if l == 0 || set == "" {
		return ""
	}

	b := make([]byte, l)
	for i := range b {
		b[i] = set[r.Intn(len(set))]
	}

	return string(b)
}

// StringSlice returns a slice of random strings of length l, using the characters in the set.
// If l is less than 0, StringSlice panics.
// If the set is empty, or l is 0, StringSlice returns an empty slice.
// If count is less than 0, StringSlice will panic.
func StringSlice(l int, set string, count int) []string {
	if count < 0 {
		panic("Random string slice count must not be less than 0")
	}

	s := make([]string, count)
	for i := range s {
		s[i] = String(l, set)
	}

	return s
}
