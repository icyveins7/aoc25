package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
  fmt.Println("Day 02")

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	ranges := strings.Split(string(content), "\n")



}
