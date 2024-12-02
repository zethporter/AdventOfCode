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

func stillSafe(currentNum int, nextNum int, decreasing bool) bool {
	if (currentNum == nextNum) {
		return false
	}
	if (decreasing) {
		if (currentNum < nextNum) {
			return false
		}
		if (currentNum - nextNum > 3) {
			return false
		}
	} else {
		if (nextNum < currentNum ) {
			return false
		}
		if (nextNum - currentNum > 3) {
			return false
		}
	}

	return true
}

// dayOneTwoCmd represents the dayOneTwo command
var dayTwoOneCmd = &cobra.Command{
	Use:   "dayTwoOne",
	Short: "Day 2 Question 1",
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
				if (!isSafe) {
					break
				}
				if (i + 1 == max) {
					safe++
					break
				}
			}
		}
		readFile.Close()

		fmt.Println("The total safety score is", safe)
	},
}

func init() {
	rootCmd.AddCommand(dayTwoOneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayOneTwoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayOneTwoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
