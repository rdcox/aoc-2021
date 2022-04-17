/*
Copyright © 2022 NAME HERE ryancox101@gmail.com

*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Day3Cmd represents the Day3 command
var Day3Cmd = &cobra.Command{
	Use:   "Day3",
	Short: "Day 3: Binary Diagnostic",
	Long: `
	Day 3: Binary Diagnostic
	Link: https://adventofcode.com/2021/day/3
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Day 3 - Part A")
		Day3a()
		fmt.Println("Day 3 - Part B")
		Day3b()
	},
}

func init() {
	rootCmd.AddCommand(Day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Day3a() {
	// read everything from the input file
	data, err := os.ReadFile("./input/Day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	// make a list of 1's for
	sa := strings.Split(string(data), "\n")
	counts := make([]int, len(sa[0]))
	for i := range sa {
		for j := range sa[i] {
			// TODO: bad conversion
			res, err := strconv.Atoi(string(sa[i][j]))
			if err != nil {
				log.Fatal(err)
			}
			counts[j] += res
		}
	}

	var gamBuf bytes.Buffer
	var epBuf bytes.Buffer
	for i := range counts {
		if counts[i] > len(sa)/2 {
			gamBuf.WriteString("1")
			epBuf.WriteString("0")
		} else {
			gamBuf.WriteString("0")
			epBuf.WriteString("1")
		}
	}

	gamma, err := strconv.ParseInt(gamBuf.String(), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(epBuf.String(), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Product of Gamma & Epsilon: %d\n", gamma*epsilon)
}

func Day3b() {

}
