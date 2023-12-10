package main

import (
	"fmt"
	"os"
	"runtime"
)

func handleErrors() {
	if errorConf != nil {
		switch runtime.GOOS {
		case "windows":
			fmt.Println("No configuration folder found!")
			fmt.Println("Export the %AppData% variable to save your times")
		case "linux":
			fmt.Println("No configuration folder found!")
			fmt.Println("Export the $XDG_CONFIG_HOME variable to save your times")
		case "darwin":
			fmt.Println("No configuration folder found!")
			fmt.Println("Export the $HOME variable to save your times")
		}
	}
	if errorOpen != nil {
		fmt.Println(filePath, "not found!\nCreating file...")
		os.Create(filePath)
		fileOpen, _ = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		fmt.Println("Found", filePath)
	}
}
