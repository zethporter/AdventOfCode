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

// dayOneTwoCmd represents the dayOneTwo command
var dayOneTwoCmd = &cobra.Command{
	Use:   "dayOneTwo",
	Short: "Day 1 question 2",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		readFile, err := os.Open("inputs/oneOne.txt")

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		var listOne []int
		var listTwo []int
		var similarityScore int

		for fileScanner.Scan() {
			str := fileScanner.Text()
			stringArr := strings.Fields(str)
			tempOne, errOne := strconv.Atoi(stringArr[0])
			tempTwo, errTwo := strconv.Atoi(stringArr[1])
			if errOne != nil || errTwo != nil {
				fmt.Println(errOne, errTwo)
			}
			if errOne == nil && errTwo == nil {
				listOne = append(listOne, tempOne)
				listTwo = append(listTwo, tempTwo)
			}
		}
		readFile.Close()

		for _, d1 := range listOne {
			// fmt.Println(id, listTwo[i], id - listTwo[i])
			// distance += id - listTwo[i]
			tempCount := 0
			for _, d2 := range listTwo {
				if d1 == d2 {
					tempCount++
				}
			}
			similarityScore += tempCount * d1
		}

		fmt.Println("The total similarity score is", similarityScore)
	},
}

func init() {
	rootCmd.AddCommand(dayOneTwoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayOneTwoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayOneTwoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
