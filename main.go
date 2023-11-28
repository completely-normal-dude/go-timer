package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"os"
)

var (
	seconds           = 0.000
	timer             = widget.NewLabel("0.00")
	timerRunning      = false
	configdir, error1 = os.UserConfigDir()
	timesPath         = configdir + string(os.PathSeparator) + "times"
	timesFile, error2 = os.OpenFile(timesPath, os.O_APPEND|os.O_RDWR, 0644)
	ch                = make(chan bool)
	timesSaved        = true
)

func main() {
	go handleErrors() // errors.go
	a := app.New()
	w := a.NewWindow("GoTime")
	go setkeys(w)     // keys.go
	tabs := setTabs() // tabs.go
	w.Resize(fyne.NewSize(400, 500))
	w.SetContent(tabs)
	w.ShowAndRun()
	defer timesFile.Close()
}
