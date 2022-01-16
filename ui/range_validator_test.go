package ui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValid(t *testing.T) {
	inputs := []string{"1", "99", "01", "001"}
	validate := NewRangeValidator(0, 100)

	for _, input := range inputs {
		err := validate(input)
		assert.Nil(t, err)
	}
}

func TestNegativeRange(t *testing.T) {
	inputs := []string{"-99", "-11"}
	validate := NewRangeValidator(-100, -10)

	for _, input := range inputs {
		err := validate(input)
		assert.Nil(t, err)
	}
}

func TestNotANumber(t *testing.T) {
	inputs := []string{"", " ", "foo", "0x", "1 ", " 1"}
	validate := NewRangeValidator(10, 100)

	for _, input := range inputs {
		err := validate(input)
		assert.EqualError(t, err, "not a valid number")
	}
}

func TestMinimum(t *testing.T) {
	err := NewRangeValidator(10, 100)("5")

	assert.EqualError(t, err, "must be greater than 10")
}

func TestMinimumIncluded(t *testing.T) {
	assert.Nil(t, NewRangeValidator(10, 100)("10"))
}

func TestMaximum(t *testing.T) {
	err := NewRangeValidator(10, 100)("105")

	assert.EqualError(t, err, "must be lesser than 100")
}

func TestMaximumIncluded(t *testing.T) {
	assert.Nil(t, NewRangeValidator(10, 100)("100"))
}
