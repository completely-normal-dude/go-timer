package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
)

var (
	seconds               = 0.000
	timer                 = widget.NewLabel("0.0")
	timerRunning          = false
	configdir, error1     = os.UserConfigDir()
	timesPath             = configdir + string(os.PathSeparator) + "times"
	scramblesPath         = configdir + string(os.PathSeparator) + "scrambles"
	timesFile, error2     = os.OpenFile(timesPath, os.O_APPEND|os.O_RDWR, 0644)
	scramblesFile, error0 = os.OpenFile(scramblesPath, os.O_APPEND|os.O_RDWR, 0644)
	ch                    = make(chan bool)
	timesSaved            = true
	currentScramble       string
)

func main() {
	go handleErrors() // errors.go
	a := app.New()
	w := a.NewWindow("Go Timer")
	tabs := setTabs()   // tabs.go
	go setkeys(w, tabs) // keys.go
	w.Resize(fyne.NewSize(600, 500))
	w.SetContent(tabs)
	w.ShowAndRun()
	defer timesFile.Close()
	defer scramblesFile.Close()
}
