// Code generated by yasupGen; DO NOT EDIT.

package yasup

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

var zeroValueFloat64 float64

//Float64Insert will append elem at the position i. Might return ErrIndexOutOfBounds.
func Float64Insert(sl *[]float64, elem float64, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//Float64Delete delete the element at the position i. Might return ErrIndexOutOfBounds.
func Float64Delete(sl *[]float64, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//Float64Contains will return true if elem is present in the slice and false otherwise.
func Float64Contains(sl []float64, elem float64) bool {
	for i := range sl {
		if sl[i] == elem {
			return true
		}
	}
	return false
}

//Float64Index returns the index of the first instance of elem, or -1 if elem is not present.
func Float64Index(sl []float64, elem float64) int {
	for i := range sl {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Float64LastIndex returns the index of the last instance of elem in the slice, or -1 if elem is not present.
func Float64LastIndex(sl []float64, elem float64) int {
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Float64Count will return an int representing the amount of times that elem is present in the slice.
func Float64Count(sl []float64, elem float64) int {
	var n int
	for i := range sl {
		if sl[i] == elem {
			n++
		}
	}
	return n
}

//Float64Push is equivalent to Float64Insert with index len(*sl).
func Float64Push(sl *[]float64, elem float64) {
	Float64Insert(sl, elem, len(*sl))
}

//Float64FrontPush is equivalent to Float64Insert with index 0.
func Float64FrontPush(sl *[]float64, elem float64) {
	Float64Insert(sl, elem, 0)
}

//Float64Pop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func Float64Pop(sl *[]float64) (float64, error) {
	if len(*sl) == 0 {
		return zeroValueFloat64, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	Float64Delete(sl, last)
	return ret, nil
}

//Float64Pop is equivalent to getting and removing the first element of the slice. Might return ErrEmptySlice.
func Float64FrontPop(sl *[]float64) (float64, error) {
	if len(*sl) == 0 {
		return zeroValueFloat64, ErrEmptySlice
	}
	ret := (*sl)[0]
	Float64Delete(sl, 0)
	return ret, nil
}

//Float64Replace modifies the slice with the first n non-overlapping instances of old replaced by new. If n equals -1, there is no limit on the number of replacements.
func Float64Replace(sl []float64, old, new float64, n int) (replacements int) {
	left := n
	for i := range sl {
		if left == 0 {
			break // no replacements left
		}
		if sl[i] == old {
			sl[i] = new
			left--
		}
	}
	return n - left
}

//Float64ReplaceAll is equivalent to Float64Replace with n = -1.
func Float64ReplaceAll(sl []float64, old, new float64) (replacements int) {
	return Float64Replace(sl, old, new, -1)
}

//Float64Equals compares two float64 slices. Returns true if their elements are equal.
func Float64Equals(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//Float64FastShuffle will randomly swap the float64 elements of a slice using math/rand (fast but not cryptographycally secure).
func Float64FastShuffle(sp []float64) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Float64SecureShuffle will randomly swap the float64 elements of a slice using crypto/rand (resource intensive but cryptographically secure).
func Float64SecureShuffle(sp []float64) error {
	var i int64
	size := int64(len(sp)) - 1
	for i = 0; i < size+1; i++ {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(size))
		if err != nil {
			return err
		}
		randI := bigRandI.Int64()
		sp[size-i], sp[randI] = sp[randI], sp[size-i]
	}
	return nil
}
