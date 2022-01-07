package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
)

func MakeSettings(win fyne.Window) fyne.CanvasObject {
	form := widget.NewForm()

	workDurationBinding := binding.NewInt()
	_ = workDurationBinding.Set(25)
	workDurationEntry := newIntegerEntryWithData(binding.IntToString(workDurationBinding))
	workDurationEntry.Validator = NewRangeValidator(0, 999)
	workDurationFormItem := widget.NewFormItem("Work duration in minutes", workDurationEntry)
	form.AppendItem(workDurationFormItem)

	shortBreakDurationBinding := binding.NewInt()
	_ = shortBreakDurationBinding.Set(5)
	shortBreakDurationEntry := newIntegerEntryWithData(binding.IntToString(shortBreakDurationBinding))
	shortBreakDurationEntry.Validator = NewRangeValidator(0, 999)
	shortBreakDurationFormItem := widget.NewFormItem("Short break duration in minutes", shortBreakDurationEntry)
	form.AppendItem(shortBreakDurationFormItem)

	longBreakDurationBinding := binding.NewInt()
	_ = longBreakDurationBinding.Set(15)
	longBreakDurationEntry := newIntegerEntryWithData(binding.IntToString(longBreakDurationBinding))
	longBreakDurationEntry.Validator = NewRangeValidator(0, 999)
	longBreakDurationFormItem := widget.NewFormItem("Long break duration in minutes", longBreakDurationEntry)
	form.AppendItem(longBreakDurationFormItem)

	workRoundsBinding := binding.NewInt()
	_ = workRoundsBinding.Set(4)
	workRoundsEntry := newIntegerEntryWithData(binding.IntToString(workRoundsBinding))
	workRoundsEntry.Validator = NewRangeValidator(0, 999999)
	workRoundsDurationFormItem := widget.NewFormItem("Work rounds", workRoundsEntry)
	form.AppendItem(workRoundsDurationFormItem)

	form.OnSubmit = func() {
		fmt.Println("Submit")
		workDuration, _ := workDurationBinding.Get()
		shortBreakDuration, _ := shortBreakDurationBinding.Get()
		longBreakDuration, _ := longBreakDurationBinding.Get()
		workRounds, _ := workRoundsBinding.Get()
		log.Println("workDuration:", workDuration)
		log.Println("shortBreakDuration:", shortBreakDuration)
		log.Println("longBreakDuration:", longBreakDuration)
		log.Println("workRounds:", workRounds)
	}
	form.OnCancel = func() {
		fmt.Println("Cancel")
		win.Close()
	}

	form.Refresh()
	return form
}
