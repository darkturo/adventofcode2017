package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextElementFuncInvalidIndex(t *testing.T) {
	i, e := getNextElement(0, 0)
	assert.Equal(t, 0, i)
	assert.NotEqual(t, nil, e)
}

func TestGetNextElementFuncInvalidSize(t *testing.T) {
	i, e := getNextElement(0, -1)
	assert.Equal(t, 0, i)
	assert.NotEqual(t, nil, e)
}

func TestGetNextElementFuncOverflow(t *testing.T) {
	i, e := getNextElement(2, 2)
	assert.Equal(t, 0, i)
	assert.NotEqual(t, nil, e)
}

func TestGetNextElementFuncSimple(t *testing.T) {
	i, e := getNextElement(0, 2)
	assert.Equal(t, 1, i)
	assert.Equal(t, nil, e)
}

func TestGetNextElementFuncCircular(t *testing.T) {
	i, e := getNextElement(1, 2)
	assert.Equal(t, 0, i)
	assert.Equal(t, nil, e)
}

func TestInvalidSequence(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("02k3"))
}

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
