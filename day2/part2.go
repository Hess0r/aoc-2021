package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var horizontal int
var depth int
var aim int

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var commands []string

	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, line)
	}

	for _, command := range commands {
		parts := strings.Split(command, " ")
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		switch parts[0] {
		case "forward":
			horizontal += value
      depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

  fmt.Printf("horizontal: %v\n", horizontal)
  fmt.Printf("depth: %v\n", depth)
  fmt.Println(horizontal * depth)
}
