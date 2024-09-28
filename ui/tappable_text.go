package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type TappableText struct {
	widget.BaseWidget
	Label    *canvas.Text
	OnTapped func()
}

func NewTappableText(title string, color color.Color, tapped func()) *TappableText {
	item := &TappableText{
		Label:    canvas.NewText(title, color),
		OnTapped: tapped,
	}
	item.ExtendBaseWidget(item)
	return item
}

func (t *TappableText) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.Label)
}

func (t *TappableText) Tapped(*fyne.PointEvent) {
	if onTapped := t.OnTapped; onTapped != nil {
		onTapped()
	}
}

func (t *TappableText) SetText(text string) {
	t.Label.Text = text
	t.Refresh()
}

func (t *TappableText) SetTextSize(textSize float32) {
	t.Label.TextSize = textSize
	t.Refresh()
}

func (t *TappableText) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (t *TappableText) SetTextStyle(style fyne.TextStyle) {
	t.Label.TextStyle = style
	t.Refresh()
}
