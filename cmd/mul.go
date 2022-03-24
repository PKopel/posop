/*
Copyright © 2022 Paweł Kopel pawel.kopel2@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/PKopel/posop/lib"
	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "multiply first number by the second one",
	Run: func(cmd *cobra.Command, args []string) {
		binaryOp(multiply, args...)
	},
}

func multiply(n, m []uint8) {
	r := lib.Multiply(base, n, m)
	rs := lib.NumToString(base, r)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	if verbose {
		fmt.Printf("%s * %s = %s\n", ns, ms, rs)
	} else {
		fmt.Printf("%s\n", rs)
	}
}

func init() {
	rootCmd.AddCommand(mulCmd)
}
