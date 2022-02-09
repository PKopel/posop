package lib

import (
	"bytes"
	"testing"
)

func TestCompare(t *testing.T) {

	cases := map[string]struct {
		base uint16
		n    []uint8
		m    []uint8
		e    int
	}{
		"less than": {
			base: uint16(10),
			n:    []uint8{1, 2, 3},
			m:    []uint8{2, 3, 4},
			e:    -1,
		},
		"greater than": {
			base: uint16(5),
			n:    []uint8{4, 2, 3},
			m:    []uint8{2, 3, 4},
			e:    1,
		},
		"equal": {
			base: uint16(23),
			n:    []uint8{11, 22, 3},
			m:    []uint8{11, 22, 3},
			e:    0,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Compare(tc.base, tc.n, tc.m)
			if r != tc.e {
				t.Errorf("%v should be %s %v", tc.n, name, tc.m)
			}
		})
	}
}

func TestAdd(t *testing.T) {

	cases := map[string]struct {
		base uint16
		n    []uint8
		m    []uint8
		e    []uint8
	}{
		"base 10": {
			base: uint16(10),
			n:    []uint8{1, 2, 3},
			m:    []uint8{2, 3, 4},
			e:    []uint8{3, 5, 7},
		},
		"base 2": {
			base: uint16(2),
			n:    []uint8{1, 0, 1},
			m:    []uint8{1, 1, 1},
			e:    []uint8{1, 1, 0, 0},
		},
		"base 23": {
			base: uint16(23),
			n:    []uint8{11, 22, 3},
			m:    []uint8{11, 22, 3},
			e:    []uint8{1, 0, 21, 6},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Add(tc.base, tc.n, tc.m)
			if !bytes.Equal(r, tc.e) {
				t.Errorf("%v + %v should be %v, is %v", tc.n, tc.m, tc.e, r)
			}
		})
	}
}

func TestEven(t *testing.T) {

	cases := map[string]struct {
		base uint16
		n    []uint8
		e    bool
	}{
		"even number even base": {
			base: uint16(10),
			n:    []uint8{1, 2, 4},
			e:    true,
		},
		"odd number even base": {
			base: uint16(4),
			n:    []uint8{3, 2, 3},
			e:    false,
		},
		"even number odd base": {
			base: uint16(23),
			n:    []uint8{11, 22, 3},
			e:    true,
		},
		"odd number odd base": {
			base: uint16(11),
			n:    []uint8{10, 3},
			e:    false,
		},
		"short odd number": {
			base: uint16(5),
			n:    []uint8{3},
			e:    false,
		},
		"short even number": {
			base: uint16(23),
			n:    []uint8{22},
			e:    true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Even(tc.base, tc.n)
			if r != tc.e {
				if tc.e {
					t.Errorf("%v in base %d should be even", tc.n, tc.base)
				} else {
					t.Errorf("%v in base %d should be odd", tc.n, tc.base)
				}
			}
		})
	}
}
