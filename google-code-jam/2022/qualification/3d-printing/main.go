package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(ints ...int) int {
	min := -1
	for _, v := range ints {
		if min == -1 {
			min = v
		} else if v < min {
			min = v
		}
	}

	return min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var testCases int
	checkedTestCases := false
	testCase := 1
	for scanner.Scan() {
		if !checkedTestCases {
			testCases, _ = strconv.Atoi(scanner.Text())
			checkedTestCases = true
		} else {
			cArr := []int{}
			mArr := []int{}
			yArr := []int{}
			kArr := []int{}
			for i := 0; i < 3; i++ {
				parts := strings.Split(scanner.Text(), " ")
				c, _ := strconv.Atoi(parts[0])
				m, _ := strconv.Atoi(parts[1])
				y, _ := strconv.Atoi(parts[2])
				k, _ := strconv.Atoi(parts[3])

				cArr = append(cArr, c)
				mArr = append(mArr, m)
				yArr = append(yArr, y)
				kArr = append(kArr, k)

				if i < 2 {
					scanner.Scan()
				}
			}

			colors := make([]int, 4)
			colors[0] = min(cArr...)
			colors[1] = min(mArr...)
			colors[2] = min(yArr...)
			colors[3] = min(kArr...)

			left := 1000000
			result := make([]int, 4)
			for i, c := range colors {
				if left >= c {
					left -= c
					result[i] = c
				} else {
					result[i] = left
					left = 0
				}
			}

			found := left == 0
			fmt.Printf("Case #%v: ", testCase)
			if !found {
				fmt.Printf("IMPOSSIBLE\n")
			} else {
				fmt.Printf("%v %v %v %v\n", result[0], result[1], result[2], result[3])
			}

			testCase++
			if testCase > testCases {
				break
			}
		}
	}
}
