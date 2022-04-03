package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(ints ...int) int {
	max := -1
	for _, v := range ints {
		if max == -1 {
			max = v
		} else if v > max {
			max = v
		}
	}

	return max
}

type Node struct {
	fun       int
	chain     *Node
	activated bool
}

func (n *Node) trigger() int {
	if !n.activated {
		fun := 0
		n.activated = true
		if n.chain != nil {
			fun += n.chain.trigger()
		}
		return max(fun, n.fun)
	}
	return 0
}

func inSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	return append(slice[s+1:], slice[:s]...)
}

func getPermutations(a []int) [][]int {
	if len(a) <= 1 {
		return [][]int{a}
	}
	perms := [][]int{}

	for i, c := range a {
		n := remove(a, i)
		for _, p := range getPermutations(n) {
			perms = append(perms, append([]int{c}, p...))
		}
	}

	return perms
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

			numModules, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()

			fun := []int{}
			chains := []int{}

			for _, f := range strings.Split(scanner.Text(), " ") {
				t, _ := strconv.Atoi(f)
				fun = append(fun, t)
			}
			scanner.Scan()
			for _, f := range strings.Split(scanner.Text(), " ") {
				t, _ := strconv.Atoi(f)
				chains = append(chains, t)
			}

			nonInitiators := []int{}
			nodes := make([]*Node, numModules)
			for i, c := range chains {
				if i >= numModules {
					break
				}
				if c == 0 {
					nodes[i] = &Node{
						fun: fun[i],
					}
				} else {
					nonInitiators = append(nonInitiators, c-1)
					nodes[i] = &Node{
						fun:   fun[i],
						chain: nodes[c-1],
					}
				}
			}

			initiatorsIndex := []int{}
			initiators := []*Node{}
			for i, node := range nodes {
				if !inSlice(i, nonInitiators) {
					initiators = append(initiators, node)
					initiatorsIndex = append(initiatorsIndex, len(initiatorsIndex))
				}
			}

			best := 0
			for _, p := range getPermutations(initiatorsIndex) {
				// Reset all nodes.
				for _, n := range nodes {
					n.activated = false
				}
				sum := 0
				for _, n := range p {
					sum += initiators[n].trigger()
				}
				if sum > best {
					best = sum
				}
			}

			fmt.Printf("Case #%v: %v\n", testCase, best)
		}
	}
}
