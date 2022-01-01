/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Day2Cmd represents the Day2 command
var Day2Cmd = &cobra.Command{
	Use:   "Day2",
	Short: "Day 2: Dive!",
	Long: `
	Day 2: Dive!
	Link: https://adventofcode.com/2021/day/2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Day 2 - Part A")
		Day2a()
		fmt.Println("Day 2 - Part B")
		Day2b()
	},
}

func init() {
	rootCmd.AddCommand(Day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Day2a() {
	data, err := os.ReadFile("./input/Day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	hzPos := 0
	depPos := 0
	// TODO: Unnecessary to compile 2 regexes
	dirPat, _ := regexp.Compile(`[a-z]+`)
	magPat, _ := regexp.Compile(`\d+`)
	sa := strings.Split(string(data), "\n")
	for i := range sa {
		direction := dirPat.FindString(sa[i])
		magnitude, err := strconv.Atoi(magPat.FindString(sa[i]))

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			hzPos += magnitude
		case "down":
			depPos += magnitude
		case "up":
			depPos -= magnitude
		}
	}

	fmt.Printf("Product of depth & horizontal pos: %d\n", hzPos*depPos)
}

func Day2b() {
	// TODO: copy/pasted this whole block from part A... gross
	data, err := os.ReadFile("./input/Day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	hzPos := 0
	depPos := 0
	aim := 0
	// TODO: Unnecessary to compile 2 regexes
	dirPat, _ := regexp.Compile(`[a-z]+`)
	magPat, _ := regexp.Compile(`\d+`)
	sa := strings.Split(string(data), "\n")
	for i := range sa {
		direction := dirPat.FindString(sa[i])
		magnitude, err := strconv.Atoi(magPat.FindString(sa[i]))

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			hzPos += magnitude
			depPos += aim * magnitude
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}

	fmt.Printf("Product of depth & horizontal pos w/aim: %d\n", hzPos*depPos)
}
