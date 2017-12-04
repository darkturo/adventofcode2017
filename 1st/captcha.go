// http://adventofcode.com/2017/day/1
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sumIfHalfwayAround := flag.Bool("halfway-around", false, "Apply rules for the second part of the exercise")
	for i := 1; i < len(os.Args); i++ {
		step := 1
		if !*sumIfHalfwayAround {
			if len(os.Args[i]) % 2 != 0 {
				fmt.Println("Size of sequence is not even")
				continue
			}

			step = calculateStep(os.Args[i])
		}
		fmt.Println(CalculateSum(os.Args[i], step))
	}
}

func calculateStep(sequence string) int {
	var step int
	if len(sequence) % 2 == 0 {
		step = len(sequence)/2
	}
	return step
}

func CalculateSum(sequence string, step int) uint64 {
	var sum uint64;
	for i := range sequence {
		j, err := getNextElement(i, len(sequence), step)
		if err != nil {
			return 0
		}
		if sequence[i] == sequence[j] {
			digit, err := strconv.ParseUint(string(sequence[i]), 10, 64)
			if err != nil {
				fmt.Printf("Invalid input: '%s'\n", sequence)
				return 0
			}
			sum = sum + digit
		}
	}
	return sum;
}

func getNextElement(currentIndex, sequenceLenght, step int) (int, error) {
	if currentIndex >= sequenceLenght {
		return 0, errors.New("Invalid parameters") 
	}

	return (currentIndex + step) % sequenceLenght, nil
}
