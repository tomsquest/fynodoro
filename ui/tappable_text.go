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
	return &responsiveTextRenderer{
		text: t,
	}
}

// responsiveTextRenderer adjusts text size to fill available space
type responsiveTextRenderer struct {
	text *TappableText
}

func (r *responsiveTextRenderer) Layout(size fyne.Size) {
	// Calculate optimal text size to fill the available space
	padding := float32(20.0)
	availableWidth := size.Width - padding
	availableHeight := size.Height - padding

	if availableWidth <= 0 || availableHeight <= 0 {
		return
	}

	// Binary search for optimal font size
	minSize := float32(10.0)
	maxSize := availableHeight
	bestSize := minSize

	for i := 0; i < 15; i++ { // 15 iterations for good precision
		testSize := (minSize + maxSize) / 2
		r.text.Label.TextSize = testSize

		textMinSize := r.text.Label.MinSize()

		// Check if text fits within bounds
		if textMinSize.Width <= availableWidth && textMinSize.Height <= availableHeight {
			bestSize = testSize
			minSize = testSize // Try larger
		} else {
			maxSize = testSize // Too big, try smaller
		}
	}

	r.text.Label.TextSize = bestSize

	// Center the text in the available space
	textSize := r.text.Label.MinSize()
	x := (size.Width - textSize.Width) / 2
	y := (size.Height - textSize.Height) / 2
	r.text.Label.Move(fyne.NewPos(x, y))
	r.text.Label.Resize(textSize)
}

func (r *responsiveTextRenderer) MinSize() fyne.Size {
	// Return a reasonable minimum size
	return fyne.NewSize(100, 50)
}

func (r *responsiveTextRenderer) Refresh() {
	r.text.Label.Refresh()
}

func (r *responsiveTextRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.text.Label}
}

func (r *responsiveTextRenderer) Destroy() {
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

func (t *TappableText) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}
