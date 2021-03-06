// Code generated by yasupGen; DO NOT EDIT.

package yasup

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

var zeroValueFloat32 float32

//Float32Insert will append elem at the position i. Might return ErrIndexOutOfBounds.
func Float32Insert(sl *[]float32, elem float32, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//Float32Delete delete the element at the position i. Might return ErrIndexOutOfBounds.
func Float32Delete(sl *[]float32, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//Float32Contains will return true if elem is present in the slice and false otherwise.
func Float32Contains(sl []float32, elem float32) bool {
	for i := range sl {
		if sl[i] == elem {
			return true
		}
	}
	return false
}

//Float32Index returns the index of the first instance of elem, or -1 if elem is not present.
func Float32Index(sl []float32, elem float32) int {
	for i := range sl {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Float32LastIndex returns the index of the last instance of elem in the slice, or -1 if elem is not present.
func Float32LastIndex(sl []float32, elem float32) int {
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Float32Count will return an int representing the amount of times that elem is present in the slice.
func Float32Count(sl []float32, elem float32) int {
	var n int
	for i := range sl {
		if sl[i] == elem {
			n++
		}
	}
	return n
}

//Float32Push is equivalent to Float32Insert with index len(*sl).
func Float32Push(sl *[]float32, elem float32) {
	Float32Insert(sl, elem, len(*sl))
}

//Float32FrontPush is equivalent to Float32Insert with index 0.
func Float32FrontPush(sl *[]float32, elem float32) {
	Float32Insert(sl, elem, 0)
}

//Float32Pop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func Float32Pop(sl *[]float32) (float32, error) {
	if len(*sl) == 0 {
		return zeroValueFloat32, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	Float32Delete(sl, last)
	return ret, nil
}

//Float32Pop is equivalent to getting and removing the first element of the slice. Might return ErrEmptySlice.
func Float32FrontPop(sl *[]float32) (float32, error) {
	if len(*sl) == 0 {
		return zeroValueFloat32, ErrEmptySlice
	}
	ret := (*sl)[0]
	Float32Delete(sl, 0)
	return ret, nil
}

//Float32Replace modifies the slice with the first n non-overlapping instances of old replaced by new. If n equals -1, there is no limit on the number of replacements.
func Float32Replace(sl []float32, old, new float32, n int) (replacements int) {
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

//Float32ReplaceAll is equivalent to Float32Replace with n = -1.
func Float32ReplaceAll(sl []float32, old, new float32) (replacements int) {
	return Float32Replace(sl, old, new, -1)
}

//Float32Equals compares two float32 slices. Returns true if their elements are equal.
func Float32Equals(a, b []float32) bool {
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

//Float32FastShuffle will randomly swap the float32 elements of a slice using math/rand (fast but not cryptographycally secure).
func Float32FastShuffle(sp []float32) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Float32SecureShuffle will randomly swap the float32 elements of a slice using crypto/rand (resource intensive but cryptographically secure).
func Float32SecureShuffle(sp []float32) error {
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
