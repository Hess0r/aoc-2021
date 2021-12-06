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

	var windowedMeasurements []int

	for i, val := range measurements {
		if i+1 <= len(measurements)-1 && i+2 <= len(measurements)-1 {
			sum := val + measurements[i+1] + measurements[i+2]
			windowedMeasurements = append(windowedMeasurements, sum)
		}
	}

	var count int

	for i, val := range windowedMeasurements[1:] {
		if val > windowedMeasurements[i] {
			count++
		}
	}

	fmt.Println(count)
}
