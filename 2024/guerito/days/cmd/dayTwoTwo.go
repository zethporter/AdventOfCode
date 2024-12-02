/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func dayTwoTwoHandler(stringArr []string, hasRemoved bool) int {
	fmt.Println(stringArr, hasRemoved)
	var decreasing bool
	var isSafe bool
	max := len(stringArr) - 1
	for i, s := range stringArr {
		if (i == max) {
			panic("Reached the end without finding if safe")
		}
		tempCurrent, err := strconv.Atoi(s)
		tempNext, errNext := strconv.Atoi(stringArr[i+1])
		if err != nil || errNext != nil {
			panic("Error parsing numbers")
		}
		if (i == 0) {
			decreasing = tempCurrent > tempNext
		}
		isSafe = stillSafe(tempCurrent, tempNext, decreasing)
		if (!isSafe && !hasRemoved) {
			var newArr = append(stringArr[:i], stringArr[i+1:]...)
			return dayTwoTwoHandler(newArr, true)
		}
		if (!isSafe && hasRemoved) {
			return 0
		}
		if (i + 1 == max) {
			return 1
		}
	}
	panic("Shouldn't have reached this point")
}

// dayOneTwoCmd represents the dayOneTwo command
var dayTwoTwo = &cobra.Command{
	Use:   "dayTwoTwo",
	Short: "Day 2 Question 2",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		readFile, err := os.Open("inputs/two.txt")

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		// var listOne []int
		// var listTwo []int
		safe := 0

		for fileScanner.Scan() {
			str := fileScanner.Text()
			stringArr := strings.Fields(str) 
			safeAdder := dayTwoTwoHandler(stringArr, false)
			if (safeAdder > 1) {
				fmt.Println(safeAdder)
			}
			safe = safeAdder + safe
		}
		readFile.Close()

		fmt.Println("The total safety score is", safe)
	},
}

func init() {
	rootCmd.AddCommand(dayTwoTwo)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayOneTwoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayOneTwoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
