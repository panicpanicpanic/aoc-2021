// https://adventofcode.com/2021/day/2
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open input text file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// x/y positioning
	var x, y int32

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), " ")

		direction := instructions[0]
		magnitude, err := strconv.ParseInt(instructions[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		switch {
		case direction == "forward":
			x += int32(magnitude)
		case direction == "up":
			y -= int32(magnitude)
		case direction == "down":
			y += int32(magnitude)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("part 1: x-coordinate: %d y-coordinate: %d product: %d \n", x, y, x*y)
}
