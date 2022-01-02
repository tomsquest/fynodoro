package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type integerEntry struct {
	widget.Entry
}

func newIntegerEntryWithData(data binding.String) *integerEntry {
	entry := &integerEntry{}
	entry.ExtendBaseWidget(entry)
	entry.Bind(data)
	return entry
}

func (e *integerEntry) TypedRune(r rune) {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		e.Entry.TypedRune(r)
	}
}

func (e *integerEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseInt(content, 10, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *integerEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
