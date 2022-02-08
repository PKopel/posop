package lib

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Base() (uint16, error) {
	var baseS string
	fmt.Print("Enter base: ")
	fmt.Scanln(&baseS)

	if b, err := strconv.Atoi(baseS); err != nil {
		return 0, fmt.Errorf("wrong format of base: %s, should be decimal number", baseS)
	} else {
		return uint16(b), nil
	}
}

func Number() ([]uint8, error) {
	fmt.Println("Enter number (in format <num> [space <num>], e.g. '1 2 3 12' for 0x123c):")
	in := bufio.NewReader(os.Stdin)
	numS, err := in.ReadString('\n')
	if err != nil {
		return nil, err
	}

	digitsS := strings.Fields(numS)
	num := make([]uint8, len(digitsS))
	for i, digitS := range digitsS {
		if d, err := strconv.Atoi(digitS); err != nil {
			return nil, fmt.Errorf("wrong format: %s, should be decimal number", digitS)
		} else {
			num[i] = uint8(d)
		}
	}

	return num, nil
}

func NumToString(base uint16, n []uint8) string {
	return fmt.Sprintf("%v base %d", n, base)
}

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
