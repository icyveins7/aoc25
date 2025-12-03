package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func detectPassZero(start, move int) int {
	passes := move / 100
	if passes < 0 {
		passes = -passes
	}
	remainder := move % 100
	fmt.Printf("move %d, start %d, initial passes %d, remainder %d\n", move, start, passes, remainder)

	if ( start + remainder > 100 ){
		passes = passes + 1
	} else if ( start + remainder < 0 ){
		passes = passes + 1
	}

	return passes
}

// This is correct, obviously, but it's a bit shit
func bruteforceloop(start, move int) (int, int) {
	fmt.Printf("bruteforceloop %d + %d\n", start, move)
	inc := 1
	if move < 0 {
		inc = -1
		move = -move
	}

	counter := 0
	val := start

	for i := 0; i < move; i++ {
		val += inc
		// fmt.Printf("bruteforceloop: val is now %d\n", val)

		if val >= 100 {
			val -= 100
		}
		if val < 0 {
			val += 100
		}

		if val == 0 {
			fmt.Printf("bruteforceloop: val landed at 0, counter incremented to %d\n", counter)
			counter += 1
		}
	}

	return val, counter
}

func main() {
	fmt.Println("Day 01")

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(content), "\n")

	// Starts at 50
	counter := 0
	bcounter := 0
	val := 50
	bval := 50
	for _, line := range lines {
		movepasses := 0
		move := 0
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

		move, err = strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
			break
		}

		if line[0] == 'L' {
			move = -move
		} else if line[0] != 'R' {
			fmt.Println("Unknown")
			break
		}
		fmt.Printf("move %d\n", move)
		movepasses = detectPassZero(val, move)
		val = (val - move) % 100
		fmt.Printf("movepasses %d\n", movepasses)


		if val == 0 {
			counter += 1
			fmt.Printf("--val landed at 0, counter incremented to %d\n", counter)
		}
		if val < 0 {
			val += 100
		}
		counter += movepasses
		if movepasses > 0 {
			fmt.Printf("---added nonzero movepasses, counter is now %d\n", counter)
		}
		fmt.Printf("=== val is now %d\n", val)

		bcounterinc := 0
		bval, bcounterinc = bruteforceloop(bval, move)
		bcounter += bcounterinc


		// debugging
		// if i == 250 {
		// 	break
		// }
	}

	fmt.Println(counter)
	fmt.Println(bcounter)
}
