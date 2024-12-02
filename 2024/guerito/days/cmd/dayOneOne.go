/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// dayOneOneCmd represents the dayOneOne command
var dayOneOneCmd = &cobra.Command{
	Use:   "dayOneOne",
	Short: "Solve Day One Question one",
	Long:  `No need for a long description`,
	Run: func(cmd *cobra.Command, args []string) {

		// filePath := os.Args[1]
		readFile, err := os.Open("inputs/oneOne.txt")

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		var listOne []int
		var listTwo []int
		var distance int

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
		sort.Ints(listOne)
		sort.Ints(listTwo)

		for i, d1 := range listOne {
			// fmt.Println(id, listTwo[i], id - listTwo[i])
			// distance += id - listTwo[i]
			if d1 > listTwo[i] {
				distance += d1 - listTwo[i]
			} else {
				distance += listTwo[i] - d1
			}
		}

		fmt.Println("The total distance is", distance)

	},
}

func init() {
	rootCmd.AddCommand(dayOneOneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayOneOneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayOneOneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
