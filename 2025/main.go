package main

import (
	"flag"
	"fmt"
	"os"

	"advent-of-code/day1"
)

func main() {
	day := flag.Int("day", 0, "The Advent of Code day to Run")
	part := flag.Int("part", 0, "The part of the day's solution to run (1 or 2)")
	inputPath := flag.String("input", "", "Path to the input file for the day")

	flag.Parse()

	if *day == 0 {
		fmt.Println("Please specify a day using  -day=<number>")
		flag.Usage()
		os.Exit(1)
	}
	if *part == 0 {
		fmt.Println("Please specifiy a part using -part=<number>")
		os.Exit(1)
	}
	if *part != 1 && *part != 2 {
		fmt.Println("Part must be 1 or 2")
		os.Exit(1)
	}

	var input string
	if *inputPath != "" {
		input = *inputPath
	} else {
		defaultInputFile := fmt.Sprintf("day%d/input.txt", *day)
		input = defaultInputFile
		fmt.Printf("Using default input file: %s\n", defaultInputFile)
	}

	switch *day {
	case 1:
		day1.RunDay1(input, *part)
	// case 2:
	// 	day2.RunDay2(input, *part)
	// Add more cases for each day
	default:
		fmt.Printf("Day %d not implemented yet.\n", *day)
	}
}
