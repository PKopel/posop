package lib

import (
	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

func add[T number](a, b T) T {
	return a + b
}

func mul[T number](a, b T) T {
	return a * b
}

// assumptions:
// 1. bases are from range [2,256] (uint16)
// 2. all numbers are from range [0,255] (uint8)
// 3. all numbers are represented as arrays with the most meaningful 'digit' first
// 4. variables declared by 'var' are initialized to 0

// result: < 0 if n < m, == 0 if n == m, > 0 if n > m
func Compare(base uint16, n []uint8, m []uint8) int {
	lenN := len(n)
	lenM := len(m)
	if lenN > lenM {
		return 1
	}
	if lenN < lenM {
		return -1
	}
	i := 0
	for n[i] == m[i] && i < lenN-1 {
		i++
	}
	if n[i] > m[i] {
		return 1
	}
	if n[i] < m[i] {
		return -1
	}

	return 0
}

type binOp = func(uint16, uint16) uint16
type carryOp = func(uint8, uint8, uint8) (uint8, uint8)

func opOnPair(base uint16, op binOp) carryOp {
	return func(n uint8, m uint8, c uint8) (uint8, uint8) {
		n16 := uint16(n)
		m16 := uint16(m)
		c16 := uint16(c)
		var r uint16
		var nc uint16

		r = op(n16, m16) + c16
		if r >= base {
			nc = r / base
			r = r % base
		}

		return uint8(r), uint8(nc)
	}
}

func Add(base uint16, n []uint8, m []uint8) []uint8 {
	lenN := len(n)
	lenM := len(m)
	addPair := opOnPair(base, add[uint16])
	// l - longer, s - shorter
	var lenL int
	var lenS int
	var s []uint8
	var l []uint8

	if lenN >= lenM {
		lenL = lenN
		l = n
		lenS = lenM
		s = m
	} else {
		lenL = lenM
		l = m
		lenS = lenN
		s = n
	}

	// creating slice for result
	// initial length is lenL, with option to extend by 1
	r := make([]uint8, lenL, lenL+1)
	var c uint8

	for lenS > 0 {
		r[lenL-1], c = addPair(l[lenL-1], s[lenS-1], c)
		lenL--
		lenS--
	}

	for lenL > 0 {
		r[lenL-1], c = addPair(l[lenL-1], 0, c)
		lenL--
	}

	if c != 0 {
		r = append(r, 0)
		copy(r[1:], r)
		r[0] = c
	}

	return r
}

func multiplyByOneDigit(base uint16, n []uint8, m uint8) []uint8 {
	lenN := len(n)
	mulPair := opOnPair(base, mul[uint16])

	r := make([]uint8, lenN, lenN+1)
	var c uint8

	for lenN > 0 {
		r[lenN-1], c = mulPair(n[lenN-1], m, c)
		lenN--
	}

	if c != 0 {
		r = append(r, 0)
		copy(r[1:], r)
		r[0] = c
	}

	return r
}

// returns n * m
func Multiply(base uint16, n []uint8, m []uint8) []uint8 {
	lenM := len(m)

	parts := make([][]uint8, lenM)

	for i := 0; lenM > i; i++ {
		parts[i] = multiplyByOneDigit(base, n, m[lenM-i-1])
		zeroes := make([]uint8, i)
		parts[i] = append(parts[i], zeroes...)
	}

	s := make([]uint8, 0)
	for _, p := range parts {
		s = Add(base, p, s)
	}

	return s
}

func Even(base uint16, n []uint8) bool {
	lenN := len(n)
	if base%2 == 0 || lenN == 1 {
		// if base is even or number is one 'digit',
		// last 'digit' is enough
		return n[lenN-1]%2 == 0
	} else {
		// if base is odd and number is more than
		// one 'digit', there must be even number
		// of odd 'digits' in an even number
		count := 0
		for _, d := range n {
			count = count + int(d%2)
		}
		return count%2 == 0
	}
}
