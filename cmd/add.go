/*
Copyright © 2022 Paweł Kopel pawel.kopel2@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/PKopel/posop/lib"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add two numbers",
	Run: func(cmd *cobra.Command, args []string) {
		binaryOp(add, args...)
	},
}

func add(n, m []uint8) {
	r := lib.Add(base, n, m)
	rs := lib.NumToString(base, r)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	if verbose {
		fmt.Printf("%s + %s = %s\n", ns, ms, rs)
	} else {
		fmt.Printf("%s\n", rs)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
