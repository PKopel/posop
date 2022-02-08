package lib

import (
	"bytes"
	"testing"
)

func TestCompare(t *testing.T) {
	base := uint16(10)
	n := []uint8{1, 2, 3}
	m := []uint8{2, 3, 4}
	r := Compare(base, n, m)
	if r >= 0 {
		t.Errorf("%v should be less than %v", n, m)
	}
}

func TestAdd(t *testing.T) {
	base := uint16(10)
	n := []uint8{1, 2, 3}
	m := []uint8{2, 3, 4}
	e := []uint8{3, 5, 7}
	r := Add(base, n, m)
	if !bytes.Equal(r, e) {
		t.Errorf("%v + %v should be %v, is %v", n, m, e, r)
	}
}

func TestEven(t *testing.T) {
	base := uint16(10)
	n := []uint8{1, 2, 3}
	m := []uint8{2, 3, 4}
	r := Even(base, n)
	if r {
		t.Errorf("%v should be odd", n)
	}
	r = Even(base, m)
	if !r {
		t.Errorf("%v should be even", m)
	}
}
