// Challenge:
// Review a sequence of digits and find the sum of all digits that match the
// next digit in the list. The list is circular, so the digit after the last
// digit is the first digit in the list.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := range os.Args[1:] {
		fmt.Println(CalculateSum(os.Args[i]))
	}
}

func CalculateSum(sequence string) uint64 {
	var sum uint64;
	for i := range sequence {
		j := getNextElement(i, len(sequence))
		if sequence[i] == sequence[j] {
			digit, err := strconv.ParseUint(string(sequence[i]), 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Invalid input: '%s'", sequence))
			}
			sum = sum + digit
		}
	}
	return sum;
}

func getNextElement(currentIndex, lenghtSequence int) int {
	if currentIndex + 1 == lenghtSequence {
		return 0
	} else {
		return currentIndex + 1
	}
}
