/*
Copyright © 2021 rdcox <ryancox101@gmail.com>

May be reproduced without limit, with attribution.

“Anyone found copying and distributing this code without permission will be considered
a mighty good friend of ours, because we don’t give a durn.”
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc-2021",
	Short: "Solutions to Advent of Code 2021",
	Long: `
	Advent of Code is an annual set of programming puzzles produced by Eric Wastl. 
	2021 is my first year participating in AOC and I have created this cobra application 
	as a convenient vehicle for organizing my solutions in golang.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aoc-2021.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
