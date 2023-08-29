package random

import "math"

// Int returns a random integer between min and max, inclusive.
// If min > max, Int will panic.
func Int(min, max int) int {
	if min > max {
		panic("Random int minimum cannot be greater than maximum")
	}

	// Handle the edge case where min == 0 and max == math.MaxInt.
	if min == 0 && max == math.MaxInt {
		return r.Int()
	}

	return r.Intn(max-min+1) + min
}

// IntSlice returns a slice of random integers between min and max, inclusive.
// If min > max, IntSlice will panic.
// If count is less or euqal to 0, IntSliceUnique will return an empty slice.
func IntSlice(min, max, count int) []int {
	if count <= 0 {
		return []int{}
	}

	s := make([]int, count)
	for i := range s {
		s[i] = Int(min, max)
	}

	return s
}

// IntSliceUnique returns a slice of unique random integers between min and max, inclusive.
// If min > max, IntSliceUnique will panic.
// If count is equal or less than 0, IntSliceUnique will return an empty slice.
// If count is greater than the number of integers between min and max, IntSliceUnique will panic.
func IntSliceUnique(min, max, count int) []int {
	if count <= 0 {
		return []int{}
	}

	if count > max-min+1 {
		panic("Random int slice unique count cannot be greater than the number of integers between min and max")
	}

	s, used := make([]int, count), make(map[int]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int(min, max)
		for used[randInt] {
			randInt = Int(min, max)
		}

		used[randInt] = true

		s[i] = randInt
	}

	return s
}

// Int8 returns a random int8 between min and max, inclusive.
// If min > max, Int8 panics.
func Int8(min, max int8) int8 {
	if min > max {
		panic("Random int8 minimum cannot be greater than maximum")
	}

	return int8(Int32(int32(min), int32(max)))
}

// Int8Slice returns a slice of random int8s between min and max, inclusive.
// If min > max, Int8Slice will panic.
// If count is less or euqal to 0, Int8SliceUnique will return an empty slice.
func Int8Slice(min, max, count int8) []int8 {
	if count <= 0 {
		return []int8{}
	}

	s := make([]int8, count)
	for i := range s {
		s[i] = Int8(min, max)
	}

	return s
}

// Int8SliceUnique returns a slice of unique random int8s between min and max, inclusive.
// If min > max, Int8SliceUnique will panic.
// If count is equal or less than 0, Int8SliceUnique will return an empty slice.
// If count is greater than the number of int8s between min and max, Int8SliceUnique will panic.
func Int8SliceUnique(min, max int8, count int) []int8 {
	if count <= 0 {
		return []int8{}
	}

	if count > int(max-min)+1 {
		panic("Random int8 slice unique count cannot be greater than the number of int8s between min and max")
	}

	s, used := make([]int8, count), make(map[int8]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int8(min, max)
		for used[randInt] {
			randInt = Int8(min, max)
		}

		used[randInt] = true

		s[i] = randInt
	}

	return s
}

// Int16 returns a random int16 between min and max, inclusive.
// If min > max, Int16 panics.
func Int16(min, max int16) int16 {
	if min > max {
		panic("Random int16 minimum cannot be greater than maximum")
	}

	return int16(Int32(int32(min), int32(max)))
}

// Int16Slice returns a slice of random int16s between min and max, inclusive.
// If min > max, Int16Slice will panic.
// If count is less or euqal to 0, Int16SliceUnique will return an empty slice.
func Int16Slice(min, max, count int16) []int16 {
	if count <= 0 {
		return []int16{}
	}

	s := make([]int16, count)
	for i := range s {
		s[i] = Int16(min, max)
	}

	return s
}

// Int16SliceUnique returns a slice of unique random int16s between min and max, inclusive.
// If min > max, Int16SliceUnique will panic.
// If count is equal or less than 0, Int16SliceUnique will return an empty slice.
// If count is greater than the number of int16s between min and max, Int16SliceUnique will panic.
func Int16SliceUnique(min, max int16, count int) []int16 {
	if count <= 0 {
		return []int16{}
	}

	if count > int(max-min)+1 {
		panic("Random int16 slice unique count cannot be greater than the number of int16s between min and max")
	}

	s, used := make([]int16, count), make(map[int16]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int16(min, max)
		for used[randInt] {
			randInt = Int16(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Int32 returns a random int32 between min and max, inclusive.
// If min > max, Int32 panics.
func Int32(min, max int32) int32 {
	if min > max {
		panic("Random int32 minimum cannot be greater than maximum")
	}

	// Handle the edge case where min == 0 and max == math.MaxInt32.
	if min == 0 && max == math.MaxInt32 {
		return r.Int31()
	}

	return r.Int31n(max-min+1) + min
}

// Int32Slice returns a slice of random int32s between min and max, inclusive.
// If min > max, Int32Slice will panic.
// If count is less or euqal to 0, Int32SliceUnique will return an empty slice.
func Int32Slice(min, max, count int32) []int32 {
	if count <= 0 {
		return []int32{}
	}

	s := make([]int32, count)
	for i := range s {
		s[i] = Int32(min, max)
	}

	return s
}

// Int32SliceUnique returns a slice of unique random int32s between min and max, inclusive.
// If min > max, Int32SliceUnique will panic.
// If count is equal or less than 0, Int32SliceUnique will return an empty slice.
// If count is greater than the number of int32s between min and max, Int32SliceUnique will panic.
func Int32SliceUnique(min, max int32, count int) []int32 {
	if count <= 0 {
		return []int32{}
	}

	if count > int(max-min)+1 {
		panic("Random int32 slice unique count cannot be greater than the number of int32s between min and max")
	}

	s, used := make([]int32, count), make(map[int32]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int32(min, max)
		for used[randInt] {
			randInt = Int32(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Int64 returns a random int64 between min and max, inclusive.
// If min > max, Int64 panics.
func Int64(min, max int64) int64 {
	if min > max {
		panic("Random int64 minimum cannot be greater than maximum")
	}

	// Handle the edge case where min == 0 and max == math.MaxInt64.
	if min == 0 && max == math.MaxInt64 {
		return r.Int63()
	}

	return r.Int63n(max-min+1) + min
}

// Int64Slice returns a slice of random int64s between min and max, inclusive.
// If min > max, Int64Slice will panic.
// If count is less or euqal to 0, Int64SliceUnique will return an empty slice.
func Int64Slice(min, max, count int64) []int64 {
	if count <= 0 {
		return []int64{}
	}

	s := make([]int64, count)
	for i := range s {
		s[i] = Int64(min, max)
	}

	return s
}

// Int64SliceUnique returns a slice of unique random int64s between min and max, inclusive.
// If min > max, Int64SliceUnique will panic.
// If count is equal or less than 0, Int64SliceUnique will return an empty slice.
// If count is greater than the number of int64s between min and max, Int64SliceUnique will panic.
func Int64SliceUnique(min, max int64, count int) []int64 {
	if count <= 0 {
		return []int64{}
	}

	if count > int(max-min)+1 {
		panic("Random int64 slice unique count cannot be greater than the number of int64s between min and max")
	}

	s, used := make([]int64, count), make(map[int64]bool, count)

	for i := 0; i < count; i++ {
		randInt := Int64(min, max)
		for used[randInt] {
			randInt = Int64(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint returns a random uint between min and max, inclusive.
// If min > max, Uint will panic.
func Uint(min, max uint) uint {
	if min > max {
		panic("Random uint minimum cannot be greater than maximum")
	}

	return uint(Uint64(uint64(min), uint64(max)))
}

// UintSlice returns a slice of random uints between min and max, inclusive.
// If min > max, UintSlice will panic.
// If count is less or euqal to 0, UintSliceUnique will return an empty slice.
func UintSlice(min, max, count uint) []uint {
	if count <= 0 {
		return []uint{}
	}

	s := make([]uint, count)
	for i := range s {
		s[i] = Uint(min, max)
	}

	return s
}

// UintSliceUnique returns a slice of unique random uints between min and max, inclusive.
// If min > max, UintSliceUnique will panic.
// If count is equal or less than 0, UintSliceUnique will return an empty slice.
// If count is greater than the number of uints between min and max, UintSliceUnique will panic.
func UintSliceUnique(min, max uint, count int) []uint {
	if count <= 0 {
		return []uint{}
	}

	if count > int(max-min)+1 {
		panic("Random uint slice unique count cannot be greater than the number of uints between min and max")
	}

	s, used := make([]uint, count), make(map[uint]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint(min, max)
		for used[randInt] {
			randInt = Uint(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint8 returns a random uint8 between min and max, inclusive.
// If min > max, Uint8 panics.
func Uint8(min, max uint8) uint8 {
	if min > max {
		panic("Random uint8 minimum cannot be greater than maximum")
	}

	return uint8(Uint32(uint32(min), uint32(max)))
}

// Uint8Slice returns a slice of random uint8s between min and max, inclusive.
// If min > max, Uint8Slice will panic.
// If count is less or euqal to 0, Uint8SliceUnique will return an empty slice.
func Uint8Slice(min, max, count uint8) []uint8 {
	if count <= 0 {
		return []uint8{}
	}

	s := make([]uint8, count)
	for i := range s {
		s[i] = Uint8(min, max)
	}

	return s
}

// Uint8SliceUnique returns a slice of unique random uint8s between min and max, inclusive.
// If min > max, Uint8SliceUnique will panic.
// If count is equal or less than 0, Uint8SliceUnique will return an empty slice.
// If count is greater than the number of uint8s between min and max, Uint8SliceUnique will panic.
func Uint8SliceUnique(min, max uint8, count int) []uint8 {
	if count <= 0 {
		return []uint8{}
	}

	if count > int(max-min)+1 {
		panic("Random uint8 slice unique count cannot be greater than the number of uint8s between min and max")
	}

	s, used := make([]uint8, count), make(map[uint8]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint8(min, max)
		for used[randInt] {
			randInt = Uint8(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint16 returns a random uint16 between min and max, inclusive.
// If min > max, Uint16 panics.
func Uint16(min, max uint16) uint16 {
	if min > max {
		panic("Random uint16 minimum cannot be greater than maximum")
	}

	return uint16(Uint32(uint32(min), uint32(max)))
}

// Uint16Slice returns a slice of random uint16s between min and max, inclusive.
// If min > max, Uint16Slice will panic.
// If count is less or euqal to 0, Uint16SliceUnique will return an empty slice.
func Uint16Slice(min, max, count uint16) []uint16 {
	if count <= 0 {
		return []uint16{}
	}

	s := make([]uint16, count)
	for i := range s {
		s[i] = Uint16(min, max)
	}

	return s
}

// Uint16SliceUnique returns a slice of unique random uint16s between min and max, inclusive.
// If min > max, Uint16SliceUnique will panic.
// If count is equal or less than 0, Uint16SliceUnique will return an empty slice.
// If count is greater than the number of uint16s between min and max, Uint16SliceUnique will panic.
func Uint16SliceUnique(min, max uint16, count int) []uint16 {
	if count <= 0 {
		return []uint16{}
	}

	if count > int(max-min)+1 {
		panic("Random uint16 slice unique count cannot be greater than the number of uint16s between min and max")
	}

	s, used := make([]uint16, count), make(map[uint16]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint16(min, max)
		for used[randInt] {
			randInt = Uint16(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint32 returns a random uint32 between min and max, inclusive.
// If min > max, Uint32 panics.
func Uint32(min, max uint32) uint32 {
	if min > max {
		panic("Random uint32 minimum cannot be greater than maximum")
	}

	// Handle the edge case where min == 0 and max == math.MaxUint32.
	if min == 0 && max == math.MaxUint32 {
		return r.Uint32()
	}

	return r.Uint32()%(max-min+1) + min
}

// Uint32Slice returns a slice of random uint32s between min and max, inclusive.
// If min > max, Uint32Slice will panic.
// If count is less or euqal to 0, Uint32SliceUnique will return an empty slice.
func Uint32Slice(min, max, count uint32) []uint32 {
	if count <= 0 {
		return []uint32{}
	}

	s := make([]uint32, count)
	for i := range s {
		s[i] = Uint32(min, max)
	}

	return s
}

// Uint32SliceUnique returns a slice of unique random uint32s between min and max, inclusive.
// If min > max, Uint32SliceUnique will panic.
// If count is equal or less than 0, Uint32SliceUnique will return an empty slice.
// If count is greater than the number of uint32s between min and max, Uint32SliceUnique will panic.
func Uint32SliceUnique(min, max uint32, count int) []uint32 {
	if count <= 0 {
		return []uint32{}
	}

	if count > int(max-min)+1 {
		panic("Random uint32 slice unique count cannot be greater than the number of uint32s between min and max")
	}

	s, used := make([]uint32, count), make(map[uint32]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint32(min, max)
		for used[randInt] {
			randInt = Uint32(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Uint64 returns a random uint64 between min and max, inclusive.
// If min > max, Uint64 panics.
func Uint64(min, max uint64) uint64 {
	if min > max {
		panic("Random uint64 minimum cannot be greater than maximum")
	}

	// Handle the edge case where min is 0 and max is math.MaxUint64.
	if min == 0 && max == math.MaxUint64 {
		return r.Uint64()
	}

	return r.Uint64()%(max-min+1) + min
}

// Uint64Slice returns a slice of random uint64s between min and max, inclusive.
// If min > max, Uint64Slice will panic.
// If count is less or euqal to 0, Uint64SliceUnique will return an empty slice.
func Uint64Slice(min, max, count uint64) []uint64 {
	if count <= 0 {
		return []uint64{}
	}

	s := make([]uint64, count)
	for i := range s {
		s[i] = Uint64(min, max)
	}

	return s
}

// Uint64SliceUnique returns a slice of unique random uint64s between min and max, inclusive.
// If min > max, Uint64SliceUnique will panic.
// If count is equal or less than 0, Uint64SliceUnique will return an empty slice.
// If count is greater than the number of uint64s between min and max, Uint64SliceUnique will panic.
func Uint64SliceUnique(min, max uint64, count int) []uint64 {
	if count <= 0 {
		return []uint64{}
	}

	if count > int(max-min)+1 {
		panic("Random uint64 slice unique count cannot be greater than the number of uint64s between min and max")
	}

	s, used := make([]uint64, count), make(map[uint64]bool, count)

	for i := 0; i < count; i++ {
		randInt := Uint64(min, max)
		for used[randInt] {
			randInt = Uint64(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Float32 returns a random float32 between min and max, inclusive.
// If min > max, Float32 panics.
func Float32(min, max float32) float32 {
	if min > max {
		panic("Random float32 minimum cannot be greater than maximum")
	}

	// Handle the edge case where min is 0 and max is math.MaxFloat32.
	if min == 0 && max == math.MaxFloat32 {
		return r.Float32()
	}

	return min + r.Float32()*(max-min)
}

// Float32Slice returns a slice of random float32s between min and max, inclusive.
// If min > max, Float32Slice will panic.
// If count is less or euqal to 0, Float32SliceUnique will return an empty slice.
func Float32Slice(min, max float32, count int) []float32 {
	if count <= 0 {
		return []float32{}
	}

	s := make([]float32, count)
	for i := range s {
		s[i] = Float32(min, max)
	}

	return s
}

// Float32SliceUnique returns a slice of unique random float32s between min and max, inclusive.
// If min > max, Float32SliceUnique will panic.
// If count is equal or less than 0, Float32SliceUnique will return an empty slice.
// If count is greater than the number of float32s between min and max, Float32SliceUnique will panic.
func Float32SliceUnique(min, max float32, count int) []float32 {
	if count <= 0 {
		return []float32{}
	}

	if count > int(max-min)+1 {
		panic("Random float32 slice unique count cannot be greater than the number of float32s between min and max")
	}

	s, used := make([]float32, count), make(map[float32]bool, count)

	for i := 0; i < count; i++ {
		randInt := Float32(min, max)
		for used[randInt] {
			randInt = Float32(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}

// Float64 returns a random float64 between min and max, inclusive.
// If min > max, Float64 panics.
func Float64(min, max float64) float64 {
	if min > max {
		panic("Random float64 minimum cannot be greater than maximum")
	}

	// Handle the edge case where min is 0 and max is math.MaxFloat64.
	if min == 0 && max == math.MaxFloat64 {
		return r.Float64()
	}

	return min + r.Float64()*(max-min)
}

// Float64Slice returns a slice of random float64s between min and max, inclusive.
// If min > max, Float64Slice will panic.
// If count is less or euqal to 0, Float64SliceUnique will return an empty slice.
func Float64Slice(min, max float64, count int) []float64 {
	if count <= 0 {
		return []float64{}
	}

	s := make([]float64, count)
	for i := range s {
		s[i] = Float64(min, max)
	}

	return s
}

// Float64SliceUnique returns a slice of unique random float64s between min and max, inclusive.
// If min > max, Float64SliceUnique will panic.
// If count is equal or less than 0, Float64SliceUnique will return an empty slice.
// If count is greater than the number of float64s between min and max, Float64SliceUnique will panic.
func Float64SliceUnique(min, max float64, count int) []float64 {
	if count <= 0 {
		return []float64{}
	}

	if count > int(max-min)+1 {
		panic("Random float64 slice unique count cannot be greater than the number of float64s between min and max")
	}

	s, used := make([]float64, count), make(map[float64]bool, count)

	for i := 0; i < count; i++ {
		randInt := Float64(min, max)
		for used[randInt] {
			randInt = Float64(min, max)
		}

		used[randInt] = true
		s[i] = randInt
	}

	return s
}
