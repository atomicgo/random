package random_test

import (
	"os"
	"testing"

	"atomicgo.dev/random"
)

func TestMain(m *testing.M) {
	random.Seed(1337)

	os.Exit(m.Run())
}

func ExampleSeed() {
	random.Seed(1337)
}

func isUnique[T comparable](s []T) bool {
	m := make(map[T]struct{})
	for _, v := range s {
		if _, ok := m[v]; ok {
			return false
		}

		m[v] = struct{}{}
	}

	return true
}

func TestIsUniqueHelper(t *testing.T) {
	t.Parallel()

	t.Run("unique", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3, 4, 5, 6}
		if !isUnique(s) {
			t.Errorf("Expected %v to be unique", s)
		}
	})

	t.Run("not unique", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3, 4, 5, 6, 1}
		if isUnique(s) {
			t.Errorf("Expected %v to not be unique", s)
		}
	})
}
