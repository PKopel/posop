package lib

import (
	"bytes"
	"testing"
)

func TestRebase(t *testing.T) {

	cases := map[string]struct {
		oldBase uint16
		n       []uint8
		newBase uint16
		e       []uint8
	}{
		"base 10 to base 2": {
			oldBase: uint16(10),
			n:       []uint8{1, 2, 7},
			newBase: uint16(2),
			e:       []uint8{1, 1, 1, 1, 1, 1, 1},
		},
		"base 2 to base 10": {
			oldBase: uint16(2),
			n:       []uint8{1, 0, 1, 1},
			newBase: uint16(10),
			e:       []uint8{1, 1},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Rebase(tc.oldBase, tc.newBase, tc.n)
			if !bytes.Equal(r, tc.e) {
				t.Errorf("%v in base %v should be %v in %v, is %v", tc.n, tc.oldBase, tc.e, tc.newBase, r)
			}
		})
	}
}

func TestTruncate(t *testing.T) {

	cases := map[string]struct {
		n []uint8
		e []uint8
	}{
		"no leading 0": {
			n: []uint8{1, 2, 7},
			e: []uint8{1, 2, 7},
		},
		"leading 0": {
			n: []uint8{0, 1, 1},
			e: []uint8{1, 1},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Truncate(tc.n)
			if !bytes.Equal(r, tc.e) {
				t.Errorf("truncated %v should be %v, is %v", tc.n, tc.e, r)
			}
		})
	}
}
