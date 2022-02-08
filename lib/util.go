package lib

import (
	"fmt"
	"math"
)

func PrintNum(base uint16, n []uint8) {
	fmt.Printf("%v base %d", n, base)
}

func Rebase(oldBase uint16, newBase uint16, n []uint8) {
	lenN := len(n)
	lenR := int(math.Log(float64(oldBase))/math.Log(float64(newBase))) * lenN
	r := make([]uint8, lenR)

}
