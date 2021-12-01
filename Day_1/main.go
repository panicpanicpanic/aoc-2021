// https://adventofcode.com/2021/day/1
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	// open input text file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// part 1: return the # of increased measurements
	var increasedMeasurements int32
	var lastMeasurement int32
	var measurements []int32

	for scanner.Scan() {
		measurement, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		// used for part 2
		measurements = append(measurements, int32(measurement))

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

	log.Printf("part 1: # of increased measurements: %d \n", increasedMeasurements)

	// part 2: return the number of three-measurement sums larger than previous sums
	var bucketSum, lastBucketSum, increasedSum int32
	for i := range measurements {
		if lastBucketSum == 0 {
			for _, r := range measurements[i : i+3] {
				lastBucketSum += r
			}
			continue
		}

		for _, r := range measurements[i : i+3] {
			bucketSum += r
		}

		if bucketSum > lastBucketSum {
			increasedSum++
		}

		lastBucketSum = bucketSum
		bucketSum = 0
	}

	log.Printf("part 2: # of sliding window increased measurements: %d \n", increasedSum)
}
