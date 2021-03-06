// Code generated by yasupGen; DO NOT EDIT.

package yasup_test

import (
	yasup "github.com/j4rv/yasup"
	"testing"
)

func Test_Int8Insert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []int8
		insertAt int
	}
	base := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tcs := []testCase{
		{"First", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
		{"Middle", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base) / 2},
		{"Last", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base)},
		{"Empty slice", []int8{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		yasup.Int8Insert(&tc.slice, -128, tc.insertAt)
		if tc.slice[tc.insertAt] != (-128) {
			t.Error(tc)
		}
	}
}

func Test_Int8FastShuffle(t *testing.T) {
	shuffles := [][]int8{}
	for i := 0; i < 8; i++ {
		or := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		yasup.Int8FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.Int8Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.Int8FastShuffle(nil)
}

func Test_Int8SecureShuffle(t *testing.T) {
	shuffles := [][]int8{}
	for i := 0; i < 8; i++ {
		or := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		yasup.Int8SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.Int8Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.Int8SecureShuffle(nil)
}

func Test_Int8Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int8
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []int8{}, true},
		{"Right nil, left empty", []int8{}, nil, true},
		{"Left nil, right not empty", nil, []int8{-128}, false},
		{"Right nil, left not empty", []int8{-128}, nil, false},

		{"Equals empty", []int8{}, []int8{}, true},
		{"Equals 1", []int8{-128}, []int8{-128}, true},
		{"Equals 2", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -128}, false},
		{"Not equals 2", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{}, false},
		{"Not equals 3", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{-128, -128}, false},
	}
	for _, tc := range tcs {
		got := yasup.Int8Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc)
		}
	}
}
