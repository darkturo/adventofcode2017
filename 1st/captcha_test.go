package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptySequence(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum(""))
}

func TestSequenceOfZeros(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("0000"))
}

func TestSequenceWithOneElement(t *testing.T) {
	assert.Equal(t, uint64(1), CalculateSum("1"))
}

func TestSequenceWithTwoDistinctNumbers(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("01"))
}

func TestSequenceWithTwoNonZeroEqualNumbers(t *testing.T) {
	assert.Equal(t, uint64(2), CalculateSum("11"))
}

func TestSequenceWithThreeDistinctNumbers(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("012"))
}

func TestSequenceWithThreeNonZeroEqualNumbers(t *testing.T) {
	assert.Equal(t, uint64(3), CalculateSum("111"))
}

func TestSequenceWithTwoConsecutives(t *testing.T) {
	assert.Equal(t, uint64(1), CalculateSum("011"))
}

func TestExample1122(t *testing.T) {
	assert.Equal(t, uint64(3), CalculateSum("1122"))
}

func TestExample1111(t *testing.T) {
	assert.Equal(t, uint64(4), CalculateSum("1111"))
}

func TestExample1234(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("1234"))
}

func TestExample91212129(t *testing.T) {
	assert.Equal(t, uint64(9), CalculateSum("91212129"))
}
