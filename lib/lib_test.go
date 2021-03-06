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
		"less than equal length": {
			base: uint16(10),
			n:    []uint8{1, 2, 3},
			m:    []uint8{2, 3, 4},
			e:    -1,
		},
		"greater than equal length": {
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
		"less than different length": {
			base: uint16(10),
			n:    []uint8{2, 3},
			m:    []uint8{2, 3, 4},
			e:    -1,
		},
		"greater than differnet length": {
			base: uint16(5),
			n:    []uint8{4, 2, 3},
			m:    []uint8{3, 4},
			e:    1,
		},
		"leading zeroes": {
			base: uint16(5),
			n:    []uint8{4, 2, 3},
			m:    []uint8{0, 0, 3, 4},
			e:    1,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Compare(tc.base, tc.n, tc.m)
			if r != tc.e {
				switch {
				case tc.e == -1:
					t.Errorf("%v should be less than %v", tc.n, tc.m)
				case tc.e == 1:
					t.Errorf("%v should be greater than %v", tc.n, tc.m)
				case tc.e == 0:
					t.Errorf("%v should be equal %v", tc.n, tc.m)
				}
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
		"base 11": {
			base: uint16(11),
			n:    []uint8{10, 0, 1},
			m:    []uint8{1, 1},
			e:    []uint8{10, 1, 2},
		},
		"base 23": {
			base: uint16(23),
			n:    []uint8{11, 22, 3},
			m:    []uint8{11, 22, 3},
			e:    []uint8{1, 0, 21, 6},
		},
		"base 222": {
			base: uint16(222),
			n:    []uint8{100, 123},
			m:    []uint8{1, 123, 14},
			e:    []uint8{2, 1, 137},
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

func TestMultiply(t *testing.T) {

	cases := map[string]struct {
		base uint16
		n    []uint8
		m    []uint8
		e    []uint8
	}{
		"base 10": {
			base: uint16(10),
			n:    []uint8{6, 2},
			m:    []uint8{4},
			e:    []uint8{2, 4, 8},
		},
		"base 2": {
			base: uint16(2),
			n:    []uint8{1, 0, 1},
			m:    []uint8{1, 0},
			e:    []uint8{1, 0, 1, 0},
		},
		"base 222": {
			base: uint16(222),
			n:    []uint8{100, 123},
			m:    []uint8{1},
			e:    []uint8{100, 123},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := Multiply(tc.base, tc.n, tc.m)
			if !bytes.Equal(r, tc.e) {
				t.Errorf("%v * %v should be %v, is %v", tc.n, tc.m, tc.e, r)
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
