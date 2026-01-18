package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
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
	f := makeForm()
	f.OnCancel = func() {
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
		(*s.win).Close()
		callback()
	}
	(*s.form).Refresh()
}

func (s *settings) SetOnClosed(callback func()) {
	(*s.win).SetOnClosed(callback)
}

func makeForm() *widget.Form {
	myPref := pref.Load()
	form := widget.NewForm()

	workDurationBinding := binding.NewInt()
	_ = workDurationBinding.Set(myPref.WorkDuration)
	form.AppendItem(newIntegerFormItem(workDurationBinding, "Work duration in minutes", "Set the duration of the Work period. Default is: %d minutes.", NewRangeValidator(0, 999)))

	shortBreakDurationBinding := binding.NewInt()
	_ = shortBreakDurationBinding.Set(myPref.ShortBreakDuration)
	form.AppendItem(newIntegerFormItem(shortBreakDurationBinding, "Short break duration in minutes", "Set the duration of the short break. Default is: %d minutes. 0 to disable.", NewRangeValidator(0, 999)))

	longBreakDurationBinding := binding.NewInt()
	_ = longBreakDurationBinding.Set(myPref.LongBreakDuration)
	form.AppendItem(newIntegerFormItem(longBreakDurationBinding, "Long break duration in minutes", "Set the duration of the long break. Default is: %d minutes. 0 to disable.", NewRangeValidator(0, 999)))

	workRoundsBinding := binding.NewInt()
	_ = workRoundsBinding.Set(myPref.WorkRounds)
	form.AppendItem(newIntegerFormItem(workRoundsBinding, "Work rounds", "Set how many Work rounds before a long break. Default is: %d. 0 to disable.", NewRangeValidator(0, 999)))

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

	fontSizeBinding := binding.NewInt()
	_ = fontSizeBinding.Set(myPref.FontSize)
	form.AppendItem(newIntegerFormItem(fontSizeBinding, "Font size", "Set the font size for the timer display. Default is: %d pixels.", NewRangeValidator(10, 200)))

	showButtonsBinding := binding.NewBool()
	_ = showButtonsBinding.Set(myPref.ShowButtons)
	showButtonsCheck := widget.NewCheckWithData("Show control buttons", showButtonsBinding)
	form.AppendItem(widget.NewFormItem("Display", showButtonsCheck))

	fontFamilyBinding := binding.NewString()
	_ = fontFamilyBinding.Set(myPref.FontFamily)
	fontFamilySelect := widget.NewSelectWithData(
		[]string{"", "Monospace", "Sans", "Serif"},
		fontFamilyBinding,
	)
	fontFamilySelect.PlaceHolder = "System Default"
	form.AppendItem(widget.NewFormItem("Font family", fontFamilySelect))

	form.OnSubmit = func() {
		workDuration, _ := workDurationBinding.Get()
		shortBreakDuration, _ := shortBreakDurationBinding.Get()
		longBreakDuration, _ := longBreakDurationBinding.Get()
		workRounds, _ := workRoundsBinding.Get()
		startMinimized, _ := startMinimizedBinding.Get()
		enableNotificationPopup, _ := enableNotificationPopupBinding.Get()
		notificationScript, _ := notificationScriptBinding.Get()
		fontSize, _ := fontSizeBinding.Get()
		showButtons, _ := showButtonsBinding.Get()
		fontFamily, _ := fontFamilyBinding.Get()

		newPref := pref.Pref{
			WorkDuration:            workDuration,
			ShortBreakDuration:      shortBreakDuration,
			LongBreakDuration:       longBreakDuration,
			WorkRounds:              workRounds,
			StartMinimized:          startMinimized,
			EnableNotificationPopup: enableNotificationPopup,
			NotificationScript:      notificationScript,
			FontSize:                fontSize,
			ShowButtons:             showButtons,
			FontFamily:              fontFamily,
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
