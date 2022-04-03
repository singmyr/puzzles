package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(threshold int, ints ...int) int {
	min := -1
	minIndex := -1
	for i, v := range ints {
		if v >= threshold && min == -1 {
			min = v
			minIndex = i
		} else if v >= threshold && v < min {
			min = v
			minIndex = i
		}
	}

	return minIndex
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var testCases int
	checkedTestCases := false
	testCase := 0
	for scanner.Scan() {
		if !checkedTestCases {
			testCases, _ = strconv.Atoi(scanner.Text())
			checkedTestCases = true
		} else {
			testCase++
			if testCase > testCases {
				break
			}

			numDice, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()

			diceData := strings.Split(scanner.Text(), " ")

			cur := 1
			chain := 0

			if len(diceData) >= numDice {
				dice := make([]int, numDice)
				for i := 0; i < numDice; i++ {
					dice[i], _ = strconv.Atoi(diceData[i])
				}

				for {
					die := min(cur, dice...)
					if die >= 0 {
						chain++
						cur++
						dice = remove(dice, die)
					} else {
						break
					}
				}
			}

			fmt.Printf("Case #%v: %v\n", testCase, chain)
		}
	}
}
