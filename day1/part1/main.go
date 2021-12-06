package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	measurements := []int{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		int, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		measurements = append(measurements, int)
	}

	var count int

	for i, val := range measurements[1:] {
		if val > measurements[i] {
			count++
		}
	}

	fmt.Println(count)
}
