package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Number(s ...string) ([]uint8, error) {
	var numS string
	var err error
	if len(s) == 0 || len(s[0]) == 0 {
		fmt.Println("Enter number (in format <num> [space <num>], e.g. '1 2 3 12' for 0x123c):")
		in := bufio.NewReader(os.Stdin)
		numS, err = in.ReadString('\n')
		if err != nil {
			return nil, err
		}
	} else {
		numS = s[0]
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
