package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func processRange(start, end int) int {
  sum := 0
  for i := start; i <= end; i++ {
    isInvalid := checkNumber(i)
    // add
    if isInvalid {
      sum += i
    }
  }

  return sum
}

func checkNumber(number int) bool {
  // convert to string

  // exit if string has odd number of chars

  // get first and second half strings

  // compare characters
}

func main() {
  fmt.Println("Day 02")

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	ranges := strings.Split(string(content), ",")

  // loop over ranges
  totalsum := 0
  for _, s := range ranges {
    // split the string
    splitstrings := strings.Split(s, "-")

		start, err = strconv.Atoi(splitstrings[0])
		end, err = strconv.Atoi(splitstrings[1])

    fmt.Printf("processing %d-%d\n", start, end)

    rangesum := processRange(start, end)
    fmt.Printf("sum %d\n", rangesum)

    totalsum += rangesum
  }

  fmt.Printf("=== totalsum %d \n", totalsum)

}
