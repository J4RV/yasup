// Code generated by yasupGen; DO NOT EDIT.

package yasup

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

var zeroValueComplex128 complex128

//Complex128Insert will append elem at the position i. Might return ErrIndexOutOfBounds.
func Complex128Insert(sl *[]complex128, elem complex128, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//Complex128Delete delete the element at the position i. Might return ErrIndexOutOfBounds.
func Complex128Delete(sl *[]complex128, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//Complex128Contains will return true if elem is present in the slice and false otherwise.
func Complex128Contains(sl []complex128, elem complex128) bool {
	for i := range sl {
		if sl[i] == elem {
			return true
		}
	}
	return false
}

//Complex128Index returns the index of the first instance of elem, or -1 if elem is not present.
func Complex128Index(sl []complex128, elem complex128) int {
	for i := range sl {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Complex128LastIndex returns the index of the last instance of elem in the slice, or -1 if elem is not present.
func Complex128LastIndex(sl []complex128, elem complex128) int {
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Complex128Count will return an int representing the amount of times that elem is present in the slice.
func Complex128Count(sl []complex128, elem complex128) int {
	var n int
	for i := range sl {
		if sl[i] == elem {
			n++
		}
	}
	return n
}

//Complex128Push is equivalent to Complex128Insert with index len(*sl).
func Complex128Push(sl *[]complex128, elem complex128) {
	Complex128Insert(sl, elem, len(*sl))
}

//Complex128FrontPush is equivalent to Complex128Insert with index 0.
func Complex128FrontPush(sl *[]complex128, elem complex128) {
	Complex128Insert(sl, elem, 0)
}

//Complex128Pop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func Complex128Pop(sl *[]complex128) (complex128, error) {
	if len(*sl) == 0 {
		return zeroValueComplex128, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	Complex128Delete(sl, last)
	return ret, nil
}

//Complex128Pop is equivalent to getting and removing the first element of the slice. Might return ErrEmptySlice.
func Complex128FrontPop(sl *[]complex128) (complex128, error) {
	if len(*sl) == 0 {
		return zeroValueComplex128, ErrEmptySlice
	}
	ret := (*sl)[0]
	Complex128Delete(sl, 0)
	return ret, nil
}

//Complex128Replace modifies the slice with the first n non-overlapping instances of old replaced by new. If n equals -1, there is no limit on the number of replacements.
func Complex128Replace(sl []complex128, old, new complex128, n int) (replacements int) {
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

//Complex128ReplaceAll is equivalent to Complex128Replace with n = -1.
func Complex128ReplaceAll(sl []complex128, old, new complex128) (replacements int) {
	return Complex128Replace(sl, old, new, -1)
}

//Complex128Equals compares two complex128 slices. Returns true if their elements are equal.
func Complex128Equals(a, b []complex128) bool {
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

//Complex128FastShuffle will randomly swap the complex128 elements of a slice using math/rand (fast but not cryptographycally secure).
func Complex128FastShuffle(sp []complex128) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Complex128SecureShuffle will randomly swap the complex128 elements of a slice using crypto/rand (resource intensive but cryptographically secure).
func Complex128SecureShuffle(sp []complex128) error {
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
