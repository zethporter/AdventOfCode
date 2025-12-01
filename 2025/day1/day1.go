// day1/day1.go
package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
)

// RunDay1 is the entry point for Day 1's solution.
func RunDay1(input string, part int) {
	fmt.Printf("Running Day 1, Part %d\n", part)
	if part == 1 {
		solvePart1(input)
	} else {
		solvePart2(input)
	}
}

func solvePart1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal("Failed to open file", "file", input)
	}

	currPos := 50
	zeroCnt := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		dir, amt, err := parseDirectionAndValue(line)
		if err == nil {

			switch dir {
			case 'R':
				currPos = roundOverAdd(currPos, amt)
				zeroCnt = isZeroAdd(currPos, zeroCnt)
			case 'L':
				currPos = roundOverSubtract(currPos, amt)
				zeroCnt = isZeroAdd(currPos, zeroCnt)
			default:
				log.Debug("Invalid Rune?")
			}
		} else {
			log.Error("Error: %v", err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Day 1, Part 1: Error Scanning input: %v", "error", err)
	}

	log.Info("Day 1, Part 1 solution:", "Solution", zeroCnt) // Replace with actual logic
}

func solvePart2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal("Failed to open file", "file", input)
	}

	currPos := 50
	zeroCnt := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		dir, amt, err := parseDirectionAndValue(line)
		if err == nil {

			switch dir {
			case 'R':
				newPos, zeroAdder := roundOverAddCountRollover(currPos, amt)
				currPos = newPos
				zeroCnt += zeroAdder
			case 'L':
				newPos, zeroAdder := roundOverSubtractCountRollover(currPos, amt)
				currPos = newPos
				zeroCnt += zeroAdder
			default:
				log.Debug("Invalid Rune?")
			}
		} else {
			log.Error("Error: %v", err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Day 1, Part 2: Error Scanning input: %v", "error", err)
	}

	log.Info("Day 1, Part 2 solution:", "Solution", zeroCnt) // Replace with actual logic
}

func parseDirectionAndValue(s string) (rune, int, error) {
	if len(s) < 2 {
		return 0, 0, fmt.Errorf("invalid format: string too short '%s'", s)
	}

	// Extract the first character as a rune (for single character).
	direction := rune(s[0])

	if direction != 'R' && direction != 'L' {
		return 0, 0, fmt.Errorf("'%v' is not a valid direction", direction)
	}

	// Extract the rest of the string which should be the number.
	numberStr := s[1:]

	// Parse the number string to an integer.
	value, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse number from '%s': %w", numberStr, err)
	}

	return direction, value, nil
}

func roundOverAdd(current int, add int) int {
	result := (current + add) % (100)
	return result
}

func roundOverSubtract(current int, subtract int) int {
	result := (current - subtract) % (100)
	if result < 0 {
		result += (100)
	}
	return result
}

func isZeroAdd(current int, zeroCnt int) int {
	if current == 0 {
		return zeroCnt + 1
	}
	return zeroCnt
}

func roundOverAddCountRollover(current int, add int) (int, int) {
	combo := current + add
	currPos := combo % (200)
	zeroCnt := 0
	if combo > 100 {
		zeroCnt += (combo - currPos) / 100
	}
	if currPos == 0 {
		zeroCnt += 1
	}
	return currPos, zeroCnt
}

func roundOverSubtractCountRollover(current int, add int) (int, int) {
	combo := current - add
	currPos := combo % (100)
	zeroCnt := 0
	if combo < -100 {
		zeroCnt += (combo - currPos) / 100
	}
	if currPos == 0 {
		zeroCnt += 1
	}
	if currPos < 0 {
		currPos += (100)
	}
	return currPos, zeroCnt
}
