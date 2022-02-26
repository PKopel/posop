package main

import (
	"fmt"

	"github.com/PKopel/posop/lib"
)

func binaryOp(op func(uint16, []uint8, []uint8)) {
	base, err := lib.Base()
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	n, err := lib.Number()
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	m, err := lib.Number()
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	op(base, n, m)
}

func compare(base uint16, n []uint8, m []uint8) {
	r := lib.Compare(base, n, m)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	switch {
	case r < 0:
		fmt.Printf("%s < %s\n", ns, ms)
	case r == 0:
		fmt.Printf("%s == %s\n", ns, ms)
	case r > 0:
		fmt.Printf("%s > %s\n", ns, ms)
	}
}

func add(base uint16, n []uint8, m []uint8) {
	r := lib.Add(base, n, m)
	rs := lib.NumToString(base, r)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	fmt.Printf("%s + %s = %s\n", ns, ms, rs)
}

func multiply(base uint16, n []uint8, m []uint8) {
	r := lib.Multiply(base, n, m)
	rs := lib.NumToString(base, r)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	fmt.Printf("%s * %s = %s\n", ns, ms, rs)
}

func even() {
	base, err := lib.Base()
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	n, err := lib.Number()
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	r := lib.Even(base, n)
	ns := lib.NumToString(base, n)
	if r {
		fmt.Printf("%s is even\n", ns)
	} else {
		fmt.Printf("%s is odd\n", ns)
	}
}

const message = `Chose operation:
1. compare
2. add
3. multiply
4. check even
Number: `

func main() {
	fmt.Print(message)
	var op string
	fmt.Scanln(&op)
	switch {
	case op == "1":
		binaryOp(compare)
	case op == "2":
		binaryOp(add)
	case op == "3":
		binaryOp(multiply)
	case op == "4":
		even()
	}
}
