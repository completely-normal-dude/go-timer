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
	scrambleText         *canvas.Text
	timerRunning         = false
	configdir, errorConf = os.UserConfigDir()
	filePath             = configdir + string(os.PathSeparator) + "gotimer"
	ch                   = make(chan bool)
	timesSaved           = false
	currentScramble      string
)

func main() {
	app := app.New()
	timer = canvas.NewText("0.0", color.White)
	timer.TextSize = 40
	w := app.NewWindow("Go Timer")
	handleErrors(app, w) // errors.go
	tabs := setTabs(w)   // tabs.go
	go setkeys(w, tabs)  // keys.go
	w.Resize(fyne.NewSize(600, 500))
	w.SetContent(tabs)
	w.ShowAndRun()
}
