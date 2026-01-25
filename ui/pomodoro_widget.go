package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/tomsquest/fynodoro/pomodoro"
	"github.com/tomsquest/fynodoro/pref"
)

type PomodoroWidget struct {
	widget.BaseWidget

	pomodoro   *pomodoro.Pomodoro
	timer      *TappableText
	playButton *widget.Button
	stopButton *widget.Button
	nextButton *widget.Button
	buttons    *fyne.Container
	content    *fyne.Container
}

func NewPomodoroWidget(thePomodoro *pomodoro.Pomodoro) *PomodoroWidget {
	l := &PomodoroWidget{
		pomodoro: thePomodoro,
	}
	l.ExtendBaseWidget(l)

	l.timer = NewTappableText(formatDuration(thePomodoro.RemainingTime), nil, nil)
	l.timer.Label.TextStyle.Bold = true
	l.timer.Label.Alignment = fyne.TextAlignCenter

	l.playButton = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	l.stopButton = widget.NewButtonWithIcon("", theme.MediaStopIcon(), nil)
	l.nextButton = widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), nil)
	l.buttons = container.NewCenter(container.NewHBox(l.playButton, l.stopButton, l.nextButton))

	l.timer.OnTapped = func() {
		l.PlayPause()
	}
	l.timer.OnTappedSecondary = func(e *fyne.PointEvent) {
		playPauseItem := fyne.NewMenuItem("Play/Pause", l.PlayPause)
		playPauseItem.Icon = theme.MediaPlayIcon()
		stopItem := fyne.NewMenuItem("Stop", l.Stop)
		stopItem.Icon = theme.MediaStopIcon()
		nextItem := fyne.NewMenuItem("Next", l.Next)
		nextItem.Icon = theme.MediaSkipNextIcon()
		menu := fyne.NewMenu("", playPauseItem, stopItem, nextItem)
		widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(l.timer), e.AbsolutePosition)
	}
	l.playButton.OnTapped = func() {
		l.PlayPause()
	}
	l.stopButton.OnTapped = func() {
		l.Stop()
	}
	l.nextButton.OnTapped = func() {
		l.Next()
	}

	thePomodoro.OnTick = func() {
		fyne.Do(func() {
			l.updateTimerDisplay()
		})
	}
	thePomodoro.OnEnd = func(kind pomodoro.Kind) {
		fyne.Do(func() {
			l.updateTimerDisplay()
			l.playButton.Icon = theme.MediaPlayIcon()
			l.playButton.Refresh()
		})
		notifyPomodoroDone(kind)
	}

	l.content = container.NewCenter(container.NewVBox(l.timer, l.buttons))
	l.ApplyPreferences()
	return l
}

func (l *PomodoroWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(l.content)
}

func (l *PomodoroWidget) ApplyPreferences() {
	prefs := pref.Load()
	l.timer.Label.TextSize = float32(prefs.TimerFontSize)
	l.timer.Label.Color = parseHexColor(prefs.TimerFontColor)
	l.timer.Refresh()
	l.SetButtonsVisible(prefs.ShowButtons)
}

func (l *PomodoroWidget) SetButtonsVisible(visible bool) {
	if visible {
		l.buttons.Show()
	} else {
		l.buttons.Hide()
	}
}

func (l *PomodoroWidget) updateTimerDisplay() {
	l.timer.SetText(formatDuration(l.pomodoro.RemainingTime))
	l.timer.Refresh()
}

func (l *PomodoroWidget) PlayPause() {
	if l.pomodoro.Running {
		l.pomodoro.Pause()
		l.playButton.Icon = theme.MediaPlayIcon()
	} else {
		l.pomodoro.Start()
		l.playButton.Icon = theme.MediaPauseIcon()
	}
	l.playButton.Refresh()
}

func (l *PomodoroWidget) Stop() {
	l.pomodoro.Stop()
	l.playButton.Icon = theme.MediaPlayIcon()
	l.playButton.Refresh()
	l.updateTimerDisplay()
}

func (l *PomodoroWidget) Next() {
	l.pomodoro.Next()
	l.playButton.Icon = theme.MediaPlayIcon()
	l.playButton.Refresh()
	l.updateTimerDisplay()
}
