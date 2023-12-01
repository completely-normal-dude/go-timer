package main

import "fyne.io/fyne/v2"

func setkeys(win fyne.Window) {
	win.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case fyne.KeySpace:
			if !timerRunning {
				startTimer(true)
			} else {
				startTimer(false)
			}
		}
	})
}
