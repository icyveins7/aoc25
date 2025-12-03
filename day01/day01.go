package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Day 01")

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(content), "\n")

	// Starts at 50
	counter := 0
	val := 50
	for _, line := range lines {
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == 'R' {
			move, err := strconv.Atoi(line[1:])
			if err != nil {
				fmt.Println(err)
				break
			}

			// fmt.Printf("+%d\n", move)
			val = (val + move) % 100
		} else if line[0] == 'L' {
			move, err := strconv.Atoi(line[1:])
			if err != nil {
				fmt.Println(err)
				break
			}

			// fmt.Printf("-%d\n", move)
			val = (val - move) % 100
		} else {
			fmt.Println("Unknown")
			break
		}
		if val == 0 {
			counter += 1
		}
	}

	fmt.Println(counter)
}
