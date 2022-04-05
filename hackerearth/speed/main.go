package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := reader.ReadString('\n')
	if b[len(b)-1] == ' ' || b[len(b)-1] == '\n' || b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}
	if b[len(b)-1] == ' ' || b[len(b)-1] == '\n' || b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}
	testCases, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}

	for testCase := 0; testCase < testCases; testCase++ {
		b, _ := reader.ReadString('\n')
		if b[len(b)-1] == ' ' || b[len(b)-1] == '\n' || b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}
		if b[len(b)-1] == ' ' || b[len(b)-1] == '\n' || b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}
		cars, _ := strconv.Atoi(b)
		b, _ = reader.ReadString('\n')
		if b[len(b)-1] == ' ' || b[len(b)-1] == '\n' || b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}
		if b[len(b)-1] == ' ' || b[len(b)-1] == '\n' || b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}
		speedData := strings.Split(b, " ")
		minSpeed, _ := strconv.Atoi(speedData[0])
		count := 1
		for i := 1; i < cars; i++ {
			speed, _ := strconv.Atoi(speedData[i])
			if speed < minSpeed {
				count++
				minSpeed = speed
			}
		}
		fmt.Println(count)
	}
}
