package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"os"
)

func handleErrors(app fyne.App, window fyne.Window) {
	if errorConf != nil {
		fmt.Println("No configuration folder found!")
		dialog.ShowInformation("No config folder", "No configuration folder found!", window)
	}
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Found", filePath)
	}
}
