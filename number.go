package random

import "math"

// Int returns a random integer between minimum and maximum, inclusive.
// If minimum > maximum, Int will panic.
func Int(minimum, maximum int) int {
	if minimum > maximum {
		panic("Random int minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum == 0 and maximum == math.MaxInt.
	if minimum == 0 && maximum == math.MaxInt {
		return r.Int()
	}

	return r.Intn(maximum-minimum+1) + minimum
}

// IntSlice returns a slice of random integers between minimum and maximum, inclusive.
// If minimum > maximum, IntSlice will panic.
// If count is less or euqal to 0, IntSliceUnique will return an empty slice.
func IntSlice(minimum, maximum, count int) []int {
	if count <= 0 {
		return []int{}
	}

	s := make([]int, count)
	for i := range s {
		s[i] = Int(minimum, maximum)
	}

	return s
}

// IntSliceUnique returns a slice of unique random integers between minimum and maximum, inclusive.
// If minimum > maximum, IntSliceUnique will panic.
// If count is equal or less than 0, IntSliceUnique will return an empty slice.
// If count is greater than the number of integers between minimum and maximum, IntSliceUnique will panic.
func IntSliceUnique(minimum, maximum, count int) []int {
	if count <= 0 {
		return []int{}
	}

	if count > maximum-minimum+1 {
		panic("Random int slice unique count cannot be greater than the number of integers between minimum and maximum")
	}

	s, used := make([]int, count), make(map[int]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int(minimum, maximum)
		for used[randInt] {
			randInt = Int(minimum, maximum)
		}

		used[randInt] = true

		s[i] = randInt
	}

	return s
}

// Int8 returns a random int8 between minimum and maximum, inclusive.
// If minimum > maximum, Int8 panics.
func Int8(minimum, maximum int8) int8 {
	if minimum > maximum {
		panic("Random int8 minimum cannot be greater than maximum")
	}

	return int8(Int32(int32(minimum), int32(maximum))) //nolint:gosec // Result is bounded by int8 inputs.
}

// Int8Slice returns a slice of random int8s between minimum and maximum, inclusive.
// If minimum > maximum, Int8Slice will panic.
// If count is less or euqal to 0, Int8SliceUnique will return an empty slice.
func Int8Slice(minimum, maximum, count int8) []int8 {
	if count <= 0 {
		return []int8{}
	}

	s := make([]int8, count)
	for i := range s {
		s[i] = Int8(minimum, maximum)
	}

	return s
}

// Int8SliceUnique returns a slice of unique random int8s between minimum and maximum, inclusive.
// If minimum > maximum, Int8SliceUnique will panic.
// If count is equal or less than 0, Int8SliceUnique will return an empty slice.
// If count is greater than the number of int8s between minimum and maximum, Int8SliceUnique will panic.
func Int8SliceUnique(minimum, maximum int8, count int) []int8 {
	if count <= 0 {
		return []int8{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random int8 slice unique count cannot be greater than the number of int8s between minimum and maximum")
	}

	s, used := make([]int8, count), make(map[int8]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int8(minimum, maximum)
		for used[randInt] {
			randInt = Int8(minimum, maximum)
		}

		used[randInt] = true

		s[i] = randInt
	}

	return s
}

// Int16 returns a random int16 between minimum and maximum, inclusive.
// If minimum > maximum, Int16 panics.
func Int16(minimum, maximum int16) int16 {
	if minimum > maximum {
		panic("Random int16 minimum cannot be greater than maximum")
	}

	return int16(Int32(int32(minimum), int32(maximum))) //nolint:gosec // Result is bounded by int16 inputs.
}

// Int16Slice returns a slice of random int16s between minimum and maximum, inclusive.
// If minimum > maximum, Int16Slice will panic.
// If count is less or euqal to 0, Int16SliceUnique will return an empty slice.
func Int16Slice(minimum, maximum, count int16) []int16 {
	if count <= 0 {
		return []int16{}
	}

	s := make([]int16, count)
	for i := range s {
		s[i] = Int16(minimum, maximum)
	}

	return s
}

// Int16SliceUnique returns a slice of unique random int16s between minimum and maximum, inclusive.
// If minimum > maximum, Int16SliceUnique will panic.
// If count is equal or less than 0, Int16SliceUnique will return an empty slice.
// If count is greater than the number of int16s between minimum and maximum, Int16SliceUnique will panic.
func Int16SliceUnique(minimum, maximum int16, count int) []int16 {
	if count <= 0 {
		return []int16{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random int16 slice unique count cannot be greater than the number of int16s between minimum and maximum")
	}

	s, used := make([]int16, count), make(map[int16]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int16(minimum, maximum)
		for used[randInt] {
			randInt = Int16(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Int32 returns a random int32 between minimum and maximum, inclusive.
// If minimum > maximum, Int32 panics.
func Int32(minimum, maximum int32) int32 {
	if minimum > maximum {
		panic("Random int32 minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum == 0 and maximum == math.MaxInt32.
	if minimum == 0 && maximum == math.MaxInt32 {
		return r.Int31()
	}

	return r.Int31n(maximum-minimum+1) + minimum
}

// Int32Slice returns a slice of random int32s between minimum and maximum, inclusive.
// If minimum > maximum, Int32Slice will panic.
// If count is less or euqal to 0, Int32SliceUnique will return an empty slice.
func Int32Slice(minimum, maximum, count int32) []int32 {
	if count <= 0 {
		return []int32{}
	}

	s := make([]int32, count)
	for i := range s {
		s[i] = Int32(minimum, maximum)
	}

	return s
}

// Int32SliceUnique returns a slice of unique random int32s between minimum and maximum, inclusive.
// If minimum > maximum, Int32SliceUnique will panic.
// If count is equal or less than 0, Int32SliceUnique will return an empty slice.
// If count is greater than the number of int32s between minimum and maximum, Int32SliceUnique will panic.
func Int32SliceUnique(minimum, maximum int32, count int) []int32 {
	if count <= 0 {
		return []int32{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random int32 slice unique count cannot be greater than the number of int32s between minimum and maximum")
	}

	s, used := make([]int32, count), make(map[int32]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int32(minimum, maximum)
		for used[randInt] {
			randInt = Int32(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Int64 returns a random int64 between minimum and maximum, inclusive.
// If minimum > maximum, Int64 panics.
func Int64(minimum, maximum int64) int64 {
	if minimum > maximum {
		panic("Random int64 minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum == 0 and maximum == math.MaxInt64.
	if minimum == 0 && maximum == math.MaxInt64 {
		return r.Int63()
	}

	return r.Int63n(maximum-minimum+1) + minimum
}

// Int64Slice returns a slice of random int64s between minimum and maximum, inclusive.
// If minimum > maximum, Int64Slice will panic.
// If count is less or euqal to 0, Int64SliceUnique will return an empty slice.
func Int64Slice(minimum, maximum, count int64) []int64 {
	if count <= 0 {
		return []int64{}
	}

	s := make([]int64, count)
	for i := range s {
		s[i] = Int64(minimum, maximum)
	}

	return s
}

// Int64SliceUnique returns a slice of unique random int64s between minimum and maximum, inclusive.
// If minimum > maximum, Int64SliceUnique will panic.
// If count is equal or less than 0, Int64SliceUnique will return an empty slice.
// If count is greater than the number of int64s between minimum and maximum, Int64SliceUnique will panic.
func Int64SliceUnique(minimum, maximum int64, count int) []int64 {
	if count <= 0 {
		return []int64{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random int64 slice unique count cannot be greater than the number of int64s between minimum and maximum")
	}

	s, used := make([]int64, count), make(map[int64]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int64(minimum, maximum)
		for used[randInt] {
			randInt = Int64(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint returns a random uint between minimum and maximum, inclusive.
// If minimum > maximum, Uint will panic.
func Uint(minimum, maximum uint) uint {
	if minimum > maximum {
		panic("Random uint minimum cannot be greater than maximum")
	}

	return uint(Uint64(uint64(minimum), uint64(maximum)))
}

// UintSlice returns a slice of random uints between minimum and maximum, inclusive.
// If minimum > maximum, UintSlice will panic.
// If count is less or euqal to 0, UintSliceUnique will return an empty slice.
func UintSlice(minimum, maximum, count uint) []uint {
	if count <= 0 {
		return []uint{}
	}

	s := make([]uint, count)
	for i := range s {
		s[i] = Uint(minimum, maximum)
	}

	return s
}

// UintSliceUnique returns a slice of unique random uints between minimum and maximum, inclusive.
// If minimum > maximum, UintSliceUnique will panic.
// If count is equal or less than 0, UintSliceUnique will return an empty slice.
// If count is greater than the number of uints between minimum and maximum, UintSliceUnique will panic.
func UintSliceUnique(minimum, maximum uint, count int) []uint {
	if count <= 0 {
		return []uint{}
	}

	if maximum-minimum < uint(count-1) {
		panic("Random uint slice unique count cannot be greater than the number of uints between minimum and maximum")
	}

	s, used := make([]uint, count), make(map[uint]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint(minimum, maximum)
		for used[randInt] {
			randInt = Uint(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint8 returns a random uint8 between minimum and maximum, inclusive.
// If minimum > maximum, Uint8 panics.
func Uint8(minimum, maximum uint8) uint8 {
	if minimum > maximum {
		panic("Random uint8 minimum cannot be greater than maximum")
	}

	return uint8(Uint32(uint32(minimum), uint32(maximum))) //nolint:gosec // Result is bounded by uint8 inputs.
}

// Uint8Slice returns a slice of random uint8s between minimum and maximum, inclusive.
// If minimum > maximum, Uint8Slice will panic.
// If count is less or euqal to 0, Uint8SliceUnique will return an empty slice.
func Uint8Slice(minimum, maximum, count uint8) []uint8 {
	if count <= 0 {
		return []uint8{}
	}

	s := make([]uint8, count)
	for i := range s {
		s[i] = Uint8(minimum, maximum)
	}

	return s
}

// Uint8SliceUnique returns a slice of unique random uint8s between minimum and maximum, inclusive.
// If minimum > maximum, Uint8SliceUnique will panic.
// If count is equal or less than 0, Uint8SliceUnique will return an empty slice.
// If count is greater than the number of uint8s between minimum and maximum, Uint8SliceUnique will panic.
func Uint8SliceUnique(minimum, maximum uint8, count int) []uint8 {
	if count <= 0 {
		return []uint8{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random uint8 slice unique count cannot be greater than the number of uint8s between minimum and maximum")
	}

	s, used := make([]uint8, count), make(map[uint8]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint8(minimum, maximum)
		for used[randInt] {
			randInt = Uint8(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint16 returns a random uint16 between minimum and maximum, inclusive.
// If minimum > maximum, Uint16 panics.
func Uint16(minimum, maximum uint16) uint16 {
	if minimum > maximum {
		panic("Random uint16 minimum cannot be greater than maximum")
	}

	return uint16(Uint32(uint32(minimum), uint32(maximum))) //nolint:gosec // Result is bounded by uint16 inputs.
}

// Uint16Slice returns a slice of random uint16s between minimum and maximum, inclusive.
// If minimum > maximum, Uint16Slice will panic.
// If count is less or euqal to 0, Uint16SliceUnique will return an empty slice.
func Uint16Slice(minimum, maximum, count uint16) []uint16 {
	if count <= 0 {
		return []uint16{}
	}

	s := make([]uint16, count)
	for i := range s {
		s[i] = Uint16(minimum, maximum)
	}

	return s
}

// Uint16SliceUnique returns a slice of unique random uint16s between minimum and maximum, inclusive.
// If minimum > maximum, Uint16SliceUnique will panic.
// If count is equal or less than 0, Uint16SliceUnique will return an empty slice.
// If count is greater than the number of uint16s between minimum and maximum, Uint16SliceUnique will panic.
func Uint16SliceUnique(minimum, maximum uint16, count int) []uint16 {
	if count <= 0 {
		return []uint16{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random uint16 slice unique count cannot be greater than the number of uint16s between minimum and maximum")
	}

	s, used := make([]uint16, count), make(map[uint16]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint16(minimum, maximum)
		for used[randInt] {
			randInt = Uint16(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint32 returns a random uint32 between minimum and maximum, inclusive.
// If minimum > maximum, Uint32 panics.
func Uint32(minimum, maximum uint32) uint32 {
	if minimum > maximum {
		panic("Random uint32 minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum == 0 and maximum == math.MaxUint32.
	if minimum == 0 && maximum == math.MaxUint32 {
		return r.Uint32()
	}

	return r.Uint32()%(maximum-minimum+1) + minimum
}

// Uint32Slice returns a slice of random uint32s between minimum and maximum, inclusive.
// If minimum > maximum, Uint32Slice will panic.
// If count is less or euqal to 0, Uint32SliceUnique will return an empty slice.
func Uint32Slice(minimum, maximum, count uint32) []uint32 {
	if count <= 0 {
		return []uint32{}
	}

	s := make([]uint32, count)
	for i := range s {
		s[i] = Uint32(minimum, maximum)
	}

	return s
}

// Uint32SliceUnique returns a slice of unique random uint32s between minimum and maximum, inclusive.
// If minimum > maximum, Uint32SliceUnique will panic.
// If count is equal or less than 0, Uint32SliceUnique will return an empty slice.
// If count is greater than the number of uint32s between minimum and maximum, Uint32SliceUnique will panic.
func Uint32SliceUnique(minimum, maximum uint32, count int) []uint32 {
	if count <= 0 {
		return []uint32{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random uint32 slice unique count cannot be greater than the number of uint32s between minimum and maximum")
	}

	s, used := make([]uint32, count), make(map[uint32]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint32(minimum, maximum)
		for used[randInt] {
			randInt = Uint32(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint64 returns a random uint64 between minimum and maximum, inclusive.
// If minimum > maximum, Uint64 panics.
func Uint64(minimum, maximum uint64) uint64 {
	if minimum > maximum {
		panic("Random uint64 minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum is 0 and maximum is math.MaxUint64.
	if minimum == 0 && maximum == math.MaxUint64 {
		return r.Uint64()
	}

	return r.Uint64()%(maximum-minimum+1) + minimum
}

// Uint64Slice returns a slice of random uint64s between minimum and maximum, inclusive.
// If minimum > maximum, Uint64Slice will panic.
// If count is less or euqal to 0, Uint64SliceUnique will return an empty slice.
func Uint64Slice(minimum, maximum, count uint64) []uint64 {
	if count <= 0 {
		return []uint64{}
	}

	s := make([]uint64, count)
	for i := range s {
		s[i] = Uint64(minimum, maximum)
	}

	return s
}

// Uint64SliceUnique returns a slice of unique random uint64s between minimum and maximum, inclusive.
// If minimum > maximum, Uint64SliceUnique will panic.
// If count is equal or less than 0, Uint64SliceUnique will return an empty slice.
// If count is greater than the number of uint64s between minimum and maximum, Uint64SliceUnique will panic.
func Uint64SliceUnique(minimum, maximum uint64, count int) []uint64 {
	if count <= 0 {
		return []uint64{}
	}

	if maximum-minimum < uint64(count-1) {
		panic("Random uint64 slice unique count cannot be greater than the number of uint64s between minimum and maximum")
	}

	s, used := make([]uint64, count), make(map[uint64]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint64(minimum, maximum)
		for used[randInt] {
			randInt = Uint64(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Float32 returns a random float32 between minimum and maximum, inclusive.
// If minimum > maximum, Float32 panics.
func Float32(minimum, maximum float32) float32 {
	if minimum > maximum {
		panic("Random float32 minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum is 0 and maximum is math.MaxFloat32.
	if minimum == 0 && maximum == math.MaxFloat32 {
		return r.Float32()
	}

	return minimum + r.Float32()*(maximum-minimum)
}

// Float32Slice returns a slice of random float32s between minimum and maximum, inclusive.
// If minimum > maximum, Float32Slice will panic.
// If count is less or euqal to 0, Float32SliceUnique will return an empty slice.
func Float32Slice(minimum, maximum float32, count int) []float32 {
	if count <= 0 {
		return []float32{}
	}

	s := make([]float32, count)
	for i := range s {
		s[i] = Float32(minimum, maximum)
	}

	return s
}

// Float32SliceUnique returns a slice of unique random float32s between minimum and maximum, inclusive.
// If minimum > maximum, Float32SliceUnique will panic.
// If count is equal or less than 0, Float32SliceUnique will return an empty slice.
// If count is greater than the number of float32s between minimum and maximum, Float32SliceUnique will panic.
func Float32SliceUnique(minimum, maximum float32, count int) []float32 {
	if count <= 0 {
		return []float32{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random float32 slice unique count cannot be greater than the number of float32s between minimum and maximum")
	}

	s, used := make([]float32, count), make(map[float32]bool, count)

	for i := 0; i < count; i++ {
		randInt := Float32(minimum, maximum)
		for used[randInt] {
			randInt = Float32(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Float64 returns a random float64 between minimum and maximum, inclusive.
// If minimum > maximum, Float64 panics.
func Float64(minimum, maximum float64) float64 {
	if minimum > maximum {
		panic("Random float64 minimum cannot be greater than maximum")
	}

	// Handle the edge case where minimum is 0 and maximum is math.MaxFloat64.
	if minimum == 0 && maximum == math.MaxFloat64 {
		return r.Float64()
	}

	return minimum + r.Float64()*(maximum-minimum)
}

// Float64Slice returns a slice of random float64s between minimum and maximum, inclusive.
// If minimum > maximum, Float64Slice will panic.
// If count is less or euqal to 0, Float64SliceUnique will return an empty slice.
func Float64Slice(minimum, maximum float64, count int) []float64 {
	if count <= 0 {
		return []float64{}
	}

	s := make([]float64, count)
	for i := range s {
		s[i] = Float64(minimum, maximum)
	}

	return s
}

// Float64SliceUnique returns a slice of unique random float64s between minimum and maximum, inclusive.
// If minimum > maximum, Float64SliceUnique will panic.
// If count is equal or less than 0, Float64SliceUnique will return an empty slice.
// If count is greater than the number of float64s between minimum and maximum, Float64SliceUnique will panic.
func Float64SliceUnique(minimum, maximum float64, count int) []float64 {
	if count <= 0 {
		return []float64{}
	}

	if count > int(maximum-minimum)+1 {
		panic("Random float64 slice unique count cannot be greater than the number of float64s between minimum and maximum")
	}

	s, used := make([]float64, count), make(map[float64]bool, count)

	for i := 0; i < count; i++ {
		randInt := Float64(minimum, maximum)
		for used[randInt] {
			randInt = Float64(minimum, maximum)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}
