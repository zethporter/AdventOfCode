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
	"unicode"

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

		var total int

		for fileScanner.Scan() {
			str := fileScanner.Text()
			int1 := 'a'
			var int2 rune

			for _, character := range str {
				if unicode.IsDigit(character) {
					if int1 == 'a' {
						int1 = character
					}
					int2 = character
				}
			}
			var sb strings.Builder
			sb.WriteRune(int1)
			sb.WriteRune(int2)

			temp, err := strconv.Atoi(sb.String())

			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				total = temp + total
			}
		}

		fmt.Println(total)
		readFile.Close()
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
