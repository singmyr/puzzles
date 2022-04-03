package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("Failed reading:", err)
	}

	fmt.Printf("%s", data)
}
