package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func setkeys(win fyne.Window, tabs *container.AppTabs) {
	win.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if tabs.SelectedIndex() == 0 {
			switch key.Name {
			case fyne.KeySpace:
				startTimer(!timerRunning)
			}
		}
	})
}
