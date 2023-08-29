package random_test

import (
	"atomicgo.dev/random"
	"fmt"
	"math"
	"testing"
)

func ExampleInt() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max = 0, math.MaxInt
	res := random.Int(min, max)

	fmt.Println(res)
	// Output: 5799089487994996006
}

func ExampleIntSlice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.IntSlice(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func TestIntSliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.IntSliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleIntSliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.IntSliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}
func ExampleUint() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max uint = 0, math.MaxUint
	res := random.Uint(min, max)

	fmt.Println(res)
	// Output: 15022461524849771814
}

func ExampleUintSlice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.UintSlice(1, 100, 10)

	fmt.Println(s)
	// Output: [15 15 53 72 9 66 79 64 53 34]
}

func TestUintSliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.UintSliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleUintSliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.UintSliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [15 53 72 9 66 79 64 34 57 17]
}

func ExampleUint8() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max uint8 = 0, math.MaxUint8
	res := random.Uint8(min, max)

	fmt.Println(res)
	// Output: 84
}

func ExampleUint8Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint8Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [77 58 99 3 33 97 89 43 68 100]
}

func TestUint8SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Uint8SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleUint8SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint8SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [77 58 99 3 33 97 89 43 68 100]
}

func ExampleUint16() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max uint16 = 0, math.MaxUint16
	res := random.Uint16(min, max)

	fmt.Println(res)
	// Output: 596
}

func ExampleUint16Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint16Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [77 58 99 3 33 97 89 43 68 100]
}

func TestUint16SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Uint16SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleUint16SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint16SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [77 58 99 3 33 97 89 43 68 100]
}

func ExampleUint32() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max uint32 = 0, math.MaxUint32
	res := random.Uint32(min, max)

	fmt.Println(res)
	// Output: 2700411476
}

func ExampleUint32Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint32Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [77 58 99 3 33 97 89 43 68 100]
}

func TestUint32SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Uint32SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleUint32SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint32SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [77 58 99 3 33 97 89 43 68 100]
}

func ExampleUint64() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max uint64 = 0, math.MaxUint64
	res := random.Uint64(min, max)

	fmt.Println(res)
	// Output: 15022461524849771814
}

func ExampleUint64Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint64Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [15 15 53 72 9 66 79 64 53 34]
}

func TestUint64SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Uint64SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleUint64SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Uint64SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [15 53 72 9 66 79 64 34 57 17]
}

func ExampleInt8() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max int8 = 0, math.MaxInt8
	res := random.Int8(min, max)

	fmt.Println(res)
	// Output: 42
}

func ExampleInt8Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int8Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func TestInt8SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Int8SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleInt8SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int8SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func ExampleInt16() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max int16 = 0, math.MaxInt16
	res := random.Int16(min, max)

	fmt.Println(res)
	// Output: 298
}

func ExampleInt16Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int16Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func TestInt16SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Int16SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleInt16SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int16SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func ExampleInt32() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max int32 = 0, math.MaxInt32
	res := random.Int32(min, max)

	fmt.Println(res)
	// Output: 1350205738
}

func ExampleInt32Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int32Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func TestInt32SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Int32SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleInt32SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int32SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [39 79 100 52 17 49 95 22 84 50]
}

func ExampleInt64() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max int64 = 0, math.MaxInt64
	res := random.Int64(min, max)

	fmt.Println(res)
	// Output: 5799089487994996006
}

func ExampleInt64Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int64Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [7 7 53 64 9 58 71 56 45 34]
}

func TestInt64SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Int64SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}

	for _, i := range s {
		if i == 0 {
			t.Errorf("Expected non-zero integer, got %d", i)
		}
	}
}

func ExampleInt64SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Int64SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [7 53 64 9 58 71 56 45 34 57]
}

func ExampleFloat32() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max float32 = 0, math.MaxFloat32
	res := random.Float32(min, max)

	fmt.Println(res)
	// Output: 0.6287385
}

func ExampleFloat32Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Float32Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [63.245113 34.879265 42.326313 96.07255 88.959915 66.83283 57.638523 84.8611 92.85987 97.56791]
}

func TestFloat32SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Float32SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}
}

func ExampleFloat32SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Float32SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [63.245113 34.879265 42.326313 96.07255 88.959915 66.83283 57.638523 84.8611 92.85987 97.56791]
}

func ExampleFloat64() {
	random.Seed(1337) // Set seed for deterministic output, not required

	var min, max float64 = 0, math.MaxFloat64
	res := random.Float64(min, max)

	fmt.Println(res)
	// Output: 0.6287385421322026
}

func ExampleFloat64Slice() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Float64Slice(1, 100, 10)

	fmt.Println(s)
	// Output: [63.24511567108806 34.87926481692567 42.326312224687584 96.0725445866481 88.95991853003598 66.83282968657517 57.63852232585161 84.86109852327596 92.85987365256264 97.56790999047796]
}

func TestFloat64SliceUnique(t *testing.T) {
	random.Seed(1337)

	s := random.Float64SliceUnique(1, 100, 10)

	if len(s) != 10 {
		t.Errorf("Expected slice length of 10, got %d", len(s))
	}

	if !isUnique(s) {
		t.Errorf("Expected unique slice, got %v", s)
	}
}

func ExampleFloat64SliceUnique() {
	random.Seed(1337) // Set seed for deterministic output, not required

	s := random.Float64SliceUnique(1, 100, 10)

	fmt.Println(s)
	// Output: [63.24511567108806 34.87926481692567 42.326312224687584 96.0725445866481 88.95991853003598 66.83282968657517 57.63852232585161 84.86109852327596 92.85987365256264 97.56790999047796]
}
