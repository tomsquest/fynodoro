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
			return errors.New("Not a valid number")
		}
		if v <= min {
			return errors.New(fmt.Sprintf("Must be greater that %d", min))
		}
		if v > max {
			return errors.New(fmt.Sprintf("Must be lesser than %d", max))
		}

		return nil // Nothing to validate with, same as having no validator.
	}
}
