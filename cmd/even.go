/*
Copyright © 2022 Paweł Kopel pawel.kopel2@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/PKopel/posop/lib"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "check if a number is even",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			os.Exit(1)
		}
		even(args)
	},
}

func even(args []string) {
	nums := lib.Slice[string](args)
	n, err := lib.Number(nums.Get(0, ""))
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	r := lib.Even(base, n)
	ns := lib.NumToString(base, n)
	if verbose {
		if r {
			fmt.Printf("%s is even\n", ns)
		} else {
			fmt.Printf("%s is odd\n", ns)
		}
	} else {
		fmt.Printf("%t\n", r)
	}
}

func init() {
	rootCmd.AddCommand(evenCmd)
}
