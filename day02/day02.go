package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func processRange(start, end int) (int, int) {
  sum := 0
  sumany := 0
  for i := start; i <= end; i++ {
    isInvalid := checkNumber(i)
    // add
    if isInvalid {
      sum += i
    }

    isInvalidAny := checkNumberAnyRepeat(i)
    // add
    if isInvalidAny {
      sumany += i
    }
  }

  return sum, sumany
}

func checkNumberAnyRepeat(number int) bool {
  // convert to string
  s := strconv.Itoa(number)

  // maximum length is half, +1 for safety
  maxRepeatLen := len(s) / 2 + 1

  for i := 1; i < maxRepeatLen; i++ {
    // get the string
    repeatString := s[:i]


    // no point if not a multiple
    if len(s) % i != 0 {
      continue
    }

    // check if every string after is the same
    found := true
    for j := i; j < len(s); j += i {
      if s[j : j+i] != repeatString {
        found = false
        break
      }
    }
    if found {
      return true
    }
  }

  return false
}

func checkNumber(number int) bool {
  // convert to string
  s := strconv.Itoa(number)

  // exit if string has odd number of chars
  if len(s) % 2 != 0 {
    return false
  }

  // get first and second half strings
  i1 := 0
  i2 := len(s) / 2

  // compare characters
  for i := 0; i < len(s) / 2; i++ {
    if s[i1] != s[i2] {
      return false
    }
    i1 += 1
    i2 += 1
  }

  return true
}

func main() {
  fmt.Println("Day 02")

  content, err := os.ReadFile("input.txt")
  if err != nil {
    fmt.Println(err)
  }

  ranges := strings.Split(strings.TrimSpace(string(content)), ",")

  // loop over ranges
  totalsum := 0
  totalsumany := 0
  for _, s := range ranges {
    // split the string
    splitstrings := strings.Split(s, "-")

    start, err := strconv.Atoi(splitstrings[0])
    if err != nil {
      fmt.Println(err)
      break
    }
    end, err := strconv.Atoi(splitstrings[1])
    if err != nil {
      fmt.Println(err)
      break
    }

    fmt.Printf("processing %d-%d\n", start, end)

    rangesum, rangesumany := processRange(start, end)
    fmt.Printf("sum %d\n", rangesum)

    totalsum += rangesum
    totalsumany += rangesumany
  }

  fmt.Printf("=== totalsum %d \n", totalsum)
  fmt.Printf("=== totalsumany %d \n", totalsumany)

}
