package main

import "text/template"

/*
  TODO new:
  FrontPush(slice, val)
  FrontPop(slice) (val, error)
  Contains(slice, val) bool
  Count(slice, val) int
  Index(slice, val) int
  LastIndex(slice, val) int
  Replace(slice, old, new, n) replacements int
  ReplaceAll(slice, old, new) replacements int
*/

var srcTemplate = template.Must(template.New("").Parse(`
// Code generated by yasupGen; DO NOT EDIT.

package {{.FilePackageSimple}}

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
	{{ if .TypePackage }}"{{ .TypePackage }}"{{ end }}
)

//{{.TypeCased}}Insert will append elem at the position i. Might return ErrIndexOutOfBounds.
func {{.TypeCased}}Insert(sl *[]{{.Type}}, elem {{.Type}}, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//{{.TypeCased}}Delete delete the element at the position i. Might return ErrIndexOutOfBounds.
func {{.TypeCased}}Delete(sl *[]{{.Type}}, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//{{.TypeCased}}Push is equivalent to {{.TypeCased}}Insert with index len(*sl)
func {{.TypeCased}}Push(sl *[]{{.Type}}, elem {{.Type}}) {
	{{.TypeCased}}Insert(sl, elem, len(*sl))
}

//{{.TypeCased}}Pop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func {{.TypeCased}}Pop(sl *[]{{.Type}}) ({{.Type}}, error) {
	if len(*sl) == 0 {
		var zeroVal {{.Type}}
		return zeroVal, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	{{.TypeCased}}Delete(sl, last)
	return ret, nil
}

//{{.TypeCased}}Equals compares two {{.Type}} slices. Returns true if their elements are equal.
func {{.TypeCased}}Equals(a, b []{{.Type}}) bool {
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

//{{.TypeCased}}FastShuffle will randomly swap the {{.Type}} elements of a slice using math/rand (fast but not cryptographycally secure).
func {{.TypeCased}}FastShuffle(sp []{{.Type}}) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//{{.TypeCased}}SecureShuffle will randomly swap the {{.Type}} elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func {{.TypeCased}}SecureShuffle(sp []{{.Type}}) error {
	var i int64
	size := int64(len(sp))-1
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
`))