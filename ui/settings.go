package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pref"
)

func MakeSettings(win fyne.Window, onSubmit func(newPref pref.Pref)) fyne.CanvasObject {
	myPref := pref.Load()
	form := widget.NewForm()

	workDurationBinding := binding.NewInt()
	_ = workDurationBinding.Set(myPref.WorkDuration)
	addWorkDurationField(form, workDurationBinding)

	shortBreakDurationBinding := binding.NewInt()
	_ = shortBreakDurationBinding.Set(myPref.ShortBreakDuration)
	addShortBreakField(form, shortBreakDurationBinding)

	longBreakDurationBinding := binding.NewInt()
	_ = longBreakDurationBinding.Set(myPref.LongBreakDuration)
	addLongBreakDurationField(form, longBreakDurationBinding)

	workRoundsBinding := binding.NewInt()
	_ = workRoundsBinding.Set(myPref.WorkRounds)
	addWorkRoundsField(form, workRoundsBinding)

	form.OnSubmit = func() {
		workDuration, _ := workDurationBinding.Get()
		shortBreakDuration, _ := shortBreakDurationBinding.Get()
		longBreakDuration, _ := longBreakDurationBinding.Get()
		workRounds, _ := workRoundsBinding.Get()

		newPref := pref.Pref{
			workDuration,
			shortBreakDuration,
			longBreakDuration,
			workRounds,
		}
		pref.Save(newPref)
		onSubmit(newPref)

		win.Close()
	}

	form.OnCancel = func() {
		win.Close()
	}

	form.Refresh()
	return form
}

func addWorkDurationField(form *widget.Form, bind binding.Int) {
	value, _ := bind.Get()
	entry := newIntegerEntryWithData(binding.IntToString(bind))
	entry.Validator = NewRangeValidator(0, 999)
	formItem := widget.NewFormItem("Work duration in minutes", entry)
	formItem.HintText = fmt.Sprintf("Set the duration of the Work period. Default is: %d minutes.", value)
	form.AppendItem(formItem)
}

func addShortBreakField(form *widget.Form, bind binding.Int) {
	value, _ := bind.Get()
	entry := newIntegerEntryWithData(binding.IntToString(bind))
	entry.Validator = NewRangeValidator(1, 999)
	formItem := widget.NewFormItem("Short break duration in minutes", entry)
	formItem.HintText = fmt.Sprintf("Set the duration of the short break. Default is: %d minutes.", value)
	form.AppendItem(formItem)
}

func addLongBreakDurationField(form *widget.Form, bind binding.Int) {
	value, _ := bind.Get()
	entry := newIntegerEntryWithData(binding.IntToString(bind))
	entry.Validator = NewRangeValidator(0, 999)
	formItem := widget.NewFormItem("Long break duration in minutes", entry)
	formItem.HintText = fmt.Sprintf("Set the duration of the long break. Default is: %d minutes. 0 to disable.", value)
	form.AppendItem(formItem)
}

func addWorkRoundsField(form *widget.Form, bind binding.Int) {
	value, _ := bind.Get()
	entry := newIntegerEntryWithData(binding.IntToString(bind))
	entry.Validator = NewRangeValidator(0, 999999)
	formItem := widget.NewFormItem("Work rounds", entry)
	formItem.HintText = fmt.Sprintf("Set how many Work rounds before a long break. Default is: %d. 0 to disable.", value)
	form.AppendItem(formItem)
}
