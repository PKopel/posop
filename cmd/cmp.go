/*
Copyright © 2022 Paweł Kopel pawel.kopel2@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/PKopel/posop/lib"
	"github.com/spf13/cobra"
)

// cmpCmd represents the cmp command
var cmpCmd = &cobra.Command{
	Use:   "cmp",
	Short: "compare two numbers",
	Long: `Returns: 
-1 if first number is smaller
 0 if numbers are equal
 1 if first number is greater`,
	Run: func(cmd *cobra.Command, args []string) {
		binaryOp(compare, args...)
	},
}

func compare(n, m []uint8) {
	r := lib.Compare(base, n, m)
	ns := lib.NumToString(base, n)
	ms := lib.NumToString(base, m)
	if verbose {
		switch {
		case r < 0:
			fmt.Printf("%s < %s\n", ns, ms)
		case r == 0:
			fmt.Printf("%s == %s\n", ns, ms)
		case r > 0:
			fmt.Printf("%s > %s\n", ns, ms)
		}
	} else {
		fmt.Printf("%d\n", r)
	}
}

func init() {
	rootCmd.AddCommand(cmpCmd)
}
