/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Day1Cmd represents the Day1 command
var Day1Cmd = &cobra.Command{
	Use:   "Day1",
	Short: "Day 1: Sonar Sweep",
	Long: `Day 1: Sonar Sweep
	Link: https://adventofcode.com/2021/day/1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Day 1 - Part A")
		a()
		fmt.Println("Day 1 - Part B")
		b()
	},
}

func init() {
	rootCmd.AddCommand(Day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Day1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func a() {
	input := convertInput()

	increases := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}
	fmt.Printf("Num. Increases: %d\n", increases)
}

func b() {
	input := convertInput()

	increases := 0
	for i := 3; i < len(input); i++ {
		w1 := input[i-1] + input[i-2] + input[i-3] // prev window
		w2 := input[i] + input[i-1] + input[i-2]   // current window
		if w2 > w1 {
			increases++
		}
	}
	fmt.Printf("Num. Window Increases: %d\n", increases)
}

func convertInput() []int {
	data, err := os.ReadFile("./input/Day1.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}

	sa := strings.Split(string(data), "\n")
	var input []int
	for i := range sa {
		res, err := strconv.Atoi(sa[i])
		if err != nil {
			fmt.Errorf(err.Error())
		}
		input = append(input, res)
	}

	return input
}
