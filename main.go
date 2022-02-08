package main

import (
	"fmt"

	"github.com/PKopel/posop/lib"
)

func compare() {
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

func add() {
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
	r := lib.Add(base, n, m)
	rs := lib.NumToString(base, r)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	fmt.Printf("%s + %s = %s\n", ns, ms, rs)
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

func main() {
	fmt.Print("Chose operation:\n1. compare\n2. add\n3. check even\nNumber: ")
	var op string
	fmt.Scanln(&op)
	switch {
	case op == "1":
		compare()
	case op == "2":
		add()
	case op == "3":
		even()
	}
}
