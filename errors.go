package main

import "fmt"
import "os"

func handleErrors() {
	if error1 != nil {
		fmt.Println("No configuration folder found!")
		fmt.Println("Export the $HOME variable to save your times")
	} else if error2 != nil {
		fmt.Println(timesPath, "not found!\nCreating file...")
		os.Create(timesPath)
		timesFile, _ = os.OpenFile(timesPath, os.O_APPEND|os.O_RDWR, 0644)
	} else {
		fmt.Println("Found", timesPath)
	}
}
