// https://adventofcode.com/2021/day/1
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var increasedMeasurements int32
	var lastMeasurement int32

	for scanner.Scan() {
		measurement, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		if lastMeasurement == 0 {
			lastMeasurement = int32(measurement)
			continue
		}

		if measurement > int64(lastMeasurement) {
			increasedMeasurements++
		}

		lastMeasurement = int32(measurement)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print(increasedMeasurements)
}
