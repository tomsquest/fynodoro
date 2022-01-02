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
	workDurationFormItem := widget.NewFormItem("Work duration in minutes", newIntegerEntryWithData(binding.IntToString(workDurationBinding)))
	form.AppendItem(workDurationFormItem)

	shortBreakDurationBinding := binding.NewInt()
	_ = shortBreakDurationBinding.Set(5)
	shortBreakDurationFormItem := widget.NewFormItem("Short break duration in minutes", newIntegerEntryWithData(binding.IntToString(shortBreakDurationBinding)))
	form.AppendItem(shortBreakDurationFormItem)

	longBreakDurationBinding := binding.NewInt()
	_ = longBreakDurationBinding.Set(15)
	longBreakDurationFormItem := widget.NewFormItem("Long break duration in minutes", newIntegerEntryWithData(binding.IntToString(longBreakDurationBinding)))
	form.AppendItem(longBreakDurationFormItem)

	workRoundsBinding := binding.NewInt()
	_ = workRoundsBinding.Set(4)
	workRoundsDurationFormItem := widget.NewFormItem("Work rounds", newIntegerEntryWithData(binding.IntToString(workRoundsBinding)))
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
