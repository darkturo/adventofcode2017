// Challenge:
// Review a sequence of digits and find the sum of all digits that match the
// next digit in the list. The list is circular, so the digit after the last
// digit is the first digit in the list.
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(CalculateSum(os.Args[i]))
	}
}

func CalculateSum(sequence string) uint64 {
	var sum uint64;
	for i := range sequence {
		j, err := getNextElement(i, len(sequence), 1)
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

func getNextElement(currentIndex, lenghtSequence, step int) (int, error) {
	if currentIndex >= lenghtSequence {
		return 0, errors.New("Invalid parameters") 
	}

	return (currentIndex + step) % lenghtSequence, nil
}
