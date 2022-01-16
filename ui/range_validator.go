package ui

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"strconv"
)

func NewRangeValidator(min int64, max int64) fyne.StringValidator {
	return func(text string) error {
		v, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return errors.New("not a valid number")
		}
		if v < min {
			return fmt.Errorf("must be greater than %d", min)
		}
		if v > max {
			return fmt.Errorf("must be lesser than %d", max)
		}

		return nil // Nothing to validate with, same as having no validator.
	}
}
