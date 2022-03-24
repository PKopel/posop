/*
Copyright © 2022 Paweł Kopel pawel.kopel2@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/PKopel/posop/lib"
	"github.com/spf13/cobra"
)

const message = `Chose operation:
1. compare
2. add
3. multiply
4. check even
Number: `

// rootCmd represents the base command when called without any subcommands
var (
	base    uint16
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "posop",
		Short: "Simple tool for interacting with positional systems.",
		Long: `posop is a CLI tool written in Go that allows for easy
computations with numbers in positional systems with bases from 2
to 256.`,
		Run: func(cmd *cobra.Command, args []string) {
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
				even(args)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint16VarP(&base, "base", "b", uint16(10), "Base of positional system to use")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Turn on verbose logging")
}

func binaryOp(op func([]uint8, []uint8), args ...string) {
	if len(args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}
	nums := lib.Slice[string](args)
	n, err := lib.Number(nums.Get(0, ""))
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	m, err := lib.Number(nums.Get(1, ""))
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	op(n, m)
}
