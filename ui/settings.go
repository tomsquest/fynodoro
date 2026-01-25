package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pref"
)

type Settings interface {
	Show()
	SetOnSubmit(callback func())
	SetOnClosed(callback func())
}

// type validation
var _ Settings = (*settings)(nil)

type settings struct {
	win  *fyne.Window
	form *widget.Form
}

func NewSettings() Settings {
	w := fyne.CurrentApp().NewWindow("Settings")
	f := makeForm(w)
	f.OnCancel = func() {
		w.Close()
	}
	// Close the window on Submit as well
	originalSubmit := f.OnSubmit
	f.OnSubmit = func() {
		originalSubmit()
		w.Close()
	}
	// Need to "refresh" to make the Submit and Cancel buttons appears
	f.Refresh()

	w.SetContent(f)
	return &settings{win: &w, form: f}
}

func (s *settings) Show() {
	(*s.win).Show()
}

func (s *settings) SetOnSubmit(callback func()) {
	formSubmit := s.form.OnSubmit
	s.form.OnSubmit = func() {
		formSubmit()
		callback()
	}
	(*s.form).Refresh()
}

func (s *settings) SetOnClosed(callback func()) {
	(*s.win).SetOnClosed(callback)
}

func makeForm(win fyne.Window) *widget.Form {
	myPref := pref.Load()
	form := widget.NewForm()

	workDurationBinding := binding.NewInt()
	_ = workDurationBinding.Set(myPref.WorkDuration)
	form.AppendItem(newIntegerFormItem(workDurationBinding, "Work duration in minutes", "Set the duration of the Work period. Default is: 25 minutes.", NewRangeValidator(0, 999)))

	shortBreakDurationBinding := binding.NewInt()
	_ = shortBreakDurationBinding.Set(myPref.ShortBreakDuration)
	form.AppendItem(newIntegerFormItem(shortBreakDurationBinding, "Short break duration in minutes", "Set the duration of the short break. Default is: 5 minutes. 0 to disable.", NewRangeValidator(0, 999)))

	longBreakDurationBinding := binding.NewInt()
	_ = longBreakDurationBinding.Set(myPref.LongBreakDuration)
	form.AppendItem(newIntegerFormItem(longBreakDurationBinding, "Long break duration in minutes", "Set the duration of the long break. Default is: 15 minutes. 0 to disable.", NewRangeValidator(0, 999)))

	workRoundsBinding := binding.NewInt()
	_ = workRoundsBinding.Set(myPref.WorkRounds)
	form.AppendItem(newIntegerFormItem(workRoundsBinding, "Work rounds", "Set how many Work rounds before a long break. Default is: 4. 0 to disable.", NewRangeValidator(0, 999)))

	timerFontSizeBinding := binding.NewInt()
	_ = timerFontSizeBinding.Set(myPref.TimerFontSize)
	form.AppendItem(newIntegerFormItem(timerFontSizeBinding, "Timer font size", "Set the font size of the timer. Default is: 60.", NewRangeValidator(10, 200)))

	timerFontColorBinding := binding.NewString()
	_ = timerFontColorBinding.Set(myPref.TimerFontColor)
	colorEntry := widget.NewEntryWithData(timerFontColorBinding)
	colorEntry.PlaceHolder = "#555555"
	colorButton := widget.NewButton("Pick...", func() {
		picker := dialog.NewColorPicker("Timer Font Color", "Select a color for the timer", func(c color.Color) {
			_ = timerFontColorBinding.Set(colorToHex(c))
		}, win)
		picker.Advanced = true
		currentHex, _ := timerFontColorBinding.Get()
		picker.SetColor(parseHexColor(currentHex))
		picker.Show()
	})
	colorContainer := container.NewBorder(nil, nil, nil, colorButton, colorEntry)
	colorFormItem := widget.NewFormItem("Timer font color", colorContainer)
	colorFormItem.HintText = "Color of the timer text. Default is #555555."
	form.AppendItem(colorFormItem)

	showButtonsBinding := binding.NewBool()
	_ = showButtonsBinding.Set(myPref.ShowButtons)
	showButtonsCheck := widget.NewCheckWithData("Show buttons", showButtonsBinding)
	form.AppendItem(widget.NewFormItem("Buttons", showButtonsCheck))

	startMinimizedBinding := binding.NewBool()
	_ = startMinimizedBinding.Set(myPref.StartMinimized)
	startMinimizedCheck := widget.NewCheckWithData("Start minimized to tray", startMinimizedBinding)
	form.AppendItem(widget.NewFormItem("Startup", startMinimizedCheck))

	enableNotificationPopupBinding := binding.NewBool()
	_ = enableNotificationPopupBinding.Set(myPref.EnableNotificationPopup)
	enableNotificationPopupCheck := widget.NewCheckWithData("Enable notification popup", enableNotificationPopupBinding)
	form.AppendItem(widget.NewFormItem("Notification popup", enableNotificationPopupCheck))

	notificationScriptBinding := binding.NewString()
	_ = notificationScriptBinding.Set(myPref.NotificationScript)
	notificationScriptEntry := widget.NewEntryWithData(notificationScriptBinding)
	notificationScriptEntry.PlaceHolder = "Path to notification script (empty to disable)"
	notificationScriptFormItem := widget.NewFormItem("Notification script", notificationScriptEntry)
	notificationScriptFormItem.HintText = "Path to a script to run when a pomodoro ends. Leave empty to disable. Default: /usr/share/fynodoro/notify.sh"
	form.AppendItem(notificationScriptFormItem)

	form.OnSubmit = func() {
		workDuration, _ := workDurationBinding.Get()
		shortBreakDuration, _ := shortBreakDurationBinding.Get()
		longBreakDuration, _ := longBreakDurationBinding.Get()
		workRounds, _ := workRoundsBinding.Get()
		timerFontSize, _ := timerFontSizeBinding.Get()
		timerFontColor, _ := timerFontColorBinding.Get()
		showButtons, _ := showButtonsBinding.Get()
		startMinimized, _ := startMinimizedBinding.Get()
		enableNotificationPopup, _ := enableNotificationPopupBinding.Get()
		notificationScript, _ := notificationScriptBinding.Get()

		newPref := pref.Pref{
			WorkDuration:            workDuration,
			ShortBreakDuration:      shortBreakDuration,
			LongBreakDuration:       longBreakDuration,
			WorkRounds:              workRounds,
			TimerFontSize:           timerFontSize,
			TimerFontColor:          timerFontColor,
			ShowButtons:             showButtons,
			StartMinimized:          startMinimized,
			EnableNotificationPopup: enableNotificationPopup,
			NotificationScript:      notificationScript,
		}
		pref.Save(newPref)
	}

	return form
}

func newIntegerFormItem(bind binding.Int, entryText string, hintText string, validator fyne.StringValidator) *widget.FormItem {
	value, _ := bind.Get()
	entry := newIntegerEntryWithData(binding.IntToString(bind))
	entry.Validator = validator
	formItem := widget.NewFormItem(entryText, entry)
	formItem.HintText = fmt.Sprintf(hintText, value)
	return formItem
}

func parseHexColor(hex string) color.Color {
	if hex == "" {
		return color.Gray{Y: 128}
	}
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return color.Gray{Y: 128}
	}
	var r, g, b uint8
	_, _ = fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(r>>8), uint8(g>>8), uint8(b>>8))
}
