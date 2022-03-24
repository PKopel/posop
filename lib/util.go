package lib

import (
	"math"
)

func rebasedLen(newBase uint16, oldBase uint16, oldLen int) int {
	// +1 for confidence
	scale := math.Log(float64(oldBase))/math.Log(float64(newBase)) + 1
	return int(math.Ceil(scale)) * oldLen
}

func Rebase(oldBase uint16, newBase uint16, n []uint8) []uint8 {
	lenN := len(n)
	lenR := rebasedLen(newBase, oldBase, lenN)
	r := make([]uint8, lenR)

	// will only work for not too big numbers
	var num int64
	for _, d := range n {
		num = num*int64(oldBase) + int64(d)
	}

	for num != 0 {
		r[lenR-1] = uint8(num % int64(newBase))
		num = num / int64(newBase)
		lenR--
	}

	i := 0
	for r[i] == 0 {
		i++
	}

	return r[i:]
}

type Slice[T any] []T

func (arr *Slice[T]) Get(i int, def T) T {
	if len(*arr) > i {
		return (*arr)[i]
	} else {
		return def
	}
}
