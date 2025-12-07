package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func processLine(string line) int {

}

func main {
  fmt.Println("Day 03")

  content, err := os.ReadFile("input.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  // process each line
  lines := strings.Split(string(content), "\n")
  total := 0
  for _, line := range lines {
    maxFromLine := processLine(line)
    total += maxFromLine
  }

  fmt.Printf("--- total %d\n")
}
