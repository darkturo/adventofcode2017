package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextElementFuncInvalidIndex(t *testing.T) {
	i, e := getNextElement(0, 0, 1)
	assert.Equal(t, 0, i)
	assert.NotEqual(t, nil, e)
}

func TestGetNextElementFuncInvalidSize(t *testing.T) {
	i, e := getNextElement(0, -1, 1)
	assert.Equal(t, 0, i)
	assert.NotEqual(t, nil, e)
}

func TestGetNextElementFuncOverflow(t *testing.T) {
	i, e := getNextElement(2, 2, 1)
	assert.Equal(t, 0, i)
	assert.NotEqual(t, nil, e)
}

func TestGetNextElementFuncSimple(t *testing.T) {
	i, e := getNextElement(0, 2, 1)
	assert.Equal(t, 1, i)
	assert.Equal(t, nil, e)
}

func TestGetNextElementFuncTwoSteps(t *testing.T) {
	i, e := getNextElement(0, 4, 2)
	assert.Equal(t, 2, i)
	assert.Equal(t, nil, e)
}

func TestGetNextElementFuncCircularListOneStep(t *testing.T) {
	i, e := getNextElement(1, 2, 1)
	assert.Equal(t, 0, i)
	assert.Equal(t, nil, e)
}

func TestGetNextElementFuncCircularListTwoSteps(t *testing.T) {
	i, e := getNextElement(1, 2, 2)
	assert.Equal(t, 1, i)
	assert.Equal(t, nil, e)
}

func TestGetNextElementFuncCircularListThreeSteps(t *testing.T) {
	i, e := getNextElement(1, 2, 3)
	assert.Equal(t, 0, i)
	assert.Equal(t, nil, e)
}

func TestInvalidSequence(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("02k3", 1))
}

func TestEmptySequence(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("", 1))
}

func TestSequenceOfZeros(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("0000", 1))
}

func TestSequenceWithOneElement(t *testing.T) {
	assert.Equal(t, uint64(1), CalculateSum("1", 1))
}

func TestSequenceWithTwoDistinctNumbers(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("01", 1))
}

func TestSequenceWithTwoNonZeroEqualNumbers(t *testing.T) {
	assert.Equal(t, uint64(2), CalculateSum("11", 1))
}

func TestSequenceWithThreeDistinctNumbers(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("012", 1))
}

func TestSequenceWithThreeNonZeroEqualNumbers(t *testing.T) {
	assert.Equal(t, uint64(3), CalculateSum("111", 1))
}

func TestSequenceWithTwoConsecutives(t *testing.T) {
	assert.Equal(t, uint64(1), CalculateSum("011", 1))
}

func TestExample1122(t *testing.T) {
	assert.Equal(t, uint64(3), CalculateSum("1122", 1))
}

func TestExample1111(t *testing.T) {
	assert.Equal(t, uint64(4), CalculateSum("1111", 1))
}

func TestExample1234(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("1234", 1))
}

func TestExample91212129(t *testing.T) {
	assert.Equal(t, uint64(9), CalculateSum("91212129", 1))
}

func TestExample1212(t *testing.T) {
	assert.Equal(t, uint64(6), CalculateSum("1212", calculateStep("1212")))
}

func TestExample1221(t *testing.T) {
	assert.Equal(t, uint64(0), CalculateSum("1221", calculateStep("1221")))
}

func TestExample123425(t *testing.T) {
	assert.Equal(t, uint64(4), CalculateSum("123425", calculateStep("123425")))
}

func TestExample123123(t *testing.T) {
	assert.Equal(t, uint64(12), CalculateSum("123123", calculateStep("123123")))
}

func TestExample12131415(t *testing.T) {
	assert.Equal(t, uint64(4), CalculateSum("12131415", calculateStep("12131415")))
}
