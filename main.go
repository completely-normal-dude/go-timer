package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"os"
)

var (
	seconds              = 0.0
	timer                *canvas.Text
	timerRunning         = false
	configdir, errorConf = os.UserConfigDir()
	filePath             = configdir + string(os.PathSeparator) + "gotimer"
	fileOpen, errorOpen  = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	ch                   = make(chan bool)
	timesSaved           = true
	currentScramble      string
)

func main() {
	go handleErrors() // errors.go
	a := app.New()
	timer = canvas.NewText("0.0", color.White)
	timer.TextSize = 40
	w := a.NewWindow("Go Timer")
	tabs := setTabs()   // tabs.go
	go setkeys(w, tabs) // keys.go
	w.Resize(fyne.NewSize(600, 500))
	w.SetContent(tabs)
	w.ShowAndRun()
	defer fileOpen.Close()
}
