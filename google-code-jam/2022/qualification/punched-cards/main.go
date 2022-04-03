package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			parts := strings.Split(scanner.Text(), " ")
			r, _ := strconv.Atoi(parts[0])
			c, _ := strconv.Atoi(parts[1])
			fmt.Printf("Case #%v:\n", testCase)
			// fmt.Printf("%v, %v", r, c)
			rows := (r * 2) + 1
			cols := c
			for i := 0; i < rows; i++ {
				startWithDots := i < 2
				even := "+-"
				odd := "|."
				for j := 0; j < cols; j++ {
					if startWithDots && j < 1 {
						fmt.Printf("..")
					} else {
						if i%2 == 0 {
							fmt.Printf(even)
						} else {
							fmt.Printf(odd)
						}
					}
				}
				if i%2 == 0 {
					fmt.Printf("+")
				} else {
					fmt.Printf("|")
				}
				fmt.Printf("\n")
			}
		}
	}
}
