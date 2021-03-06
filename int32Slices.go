// Code generated by yasupGen; DO NOT EDIT.

package yasup

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

var zeroValueInt32 int32

//Int32Insert will append elem at the position i. Might return ErrIndexOutOfBounds.
func Int32Insert(sl *[]int32, elem int32, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//Int32Delete delete the element at the position i. Might return ErrIndexOutOfBounds.
func Int32Delete(sl *[]int32, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//Int32Contains will return true if elem is present in the slice and false otherwise.
func Int32Contains(sl []int32, elem int32) bool {
	for i := range sl {
		if sl[i] == elem {
			return true
		}
	}
	return false
}

//Int32Index returns the index of the first instance of elem, or -1 if elem is not present.
func Int32Index(sl []int32, elem int32) int {
	for i := range sl {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Int32LastIndex returns the index of the last instance of elem in the slice, or -1 if elem is not present.
func Int32LastIndex(sl []int32, elem int32) int {
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Int32Count will return an int representing the amount of times that elem is present in the slice.
func Int32Count(sl []int32, elem int32) int {
	var n int
	for i := range sl {
		if sl[i] == elem {
			n++
		}
	}
	return n
}

//Int32Push is equivalent to Int32Insert with index len(*sl).
func Int32Push(sl *[]int32, elem int32) {
	Int32Insert(sl, elem, len(*sl))
}

//Int32FrontPush is equivalent to Int32Insert with index 0.
func Int32FrontPush(sl *[]int32, elem int32) {
	Int32Insert(sl, elem, 0)
}

//Int32Pop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func Int32Pop(sl *[]int32) (int32, error) {
	if len(*sl) == 0 {
		return zeroValueInt32, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	Int32Delete(sl, last)
	return ret, nil
}

//Int32Pop is equivalent to getting and removing the first element of the slice. Might return ErrEmptySlice.
func Int32FrontPop(sl *[]int32) (int32, error) {
	if len(*sl) == 0 {
		return zeroValueInt32, ErrEmptySlice
	}
	ret := (*sl)[0]
	Int32Delete(sl, 0)
	return ret, nil
}

//Int32Replace modifies the slice with the first n non-overlapping instances of old replaced by new. If n equals -1, there is no limit on the number of replacements.
func Int32Replace(sl []int32, old, new int32, n int) (replacements int) {
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

//Int32ReplaceAll is equivalent to Int32Replace with n = -1.
func Int32ReplaceAll(sl []int32, old, new int32) (replacements int) {
	return Int32Replace(sl, old, new, -1)
}

//Int32Equals compares two int32 slices. Returns true if their elements are equal.
func Int32Equals(a, b []int32) bool {
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

//Int32FastShuffle will randomly swap the int32 elements of a slice using math/rand (fast but not cryptographycally secure).
func Int32FastShuffle(sp []int32) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Int32SecureShuffle will randomly swap the int32 elements of a slice using crypto/rand (resource intensive but cryptographically secure).
func Int32SecureShuffle(sp []int32) error {
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
