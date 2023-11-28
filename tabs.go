package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func setTabs() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Timer", container.NewVBox(timerTab())),
		container.NewTabItem("Solves", container.NewPadded(solvesTab())),
		container.NewTabItem("Stats", container.NewCenter(statsTab())),
	)
	return tabs
}

func timerTab() fyne.CanvasObject {
	timer.TextStyle.Monospace = true
	ao5 := binding.NewString()
	ao12 := binding.NewString()
	ao50 := binding.NewString()
	ao100 := binding.NewString()
	go func() {
		for true {
			ao5.Set(getAverage(5, manageFile()))
			ao12.Set(getAverage(12, manageFile()))
			ao50.Set(getAverage(50, manageFile()))
			ao100.Set(getAverage(100, manageFile()))
			time.Sleep(3)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	return container.New(layout.NewGridLayoutWithRows(3), container.NewCenter(gen_scramble_tab()), container.NewCenter(timer), container.NewGridWithRows(4, gen_avg_tab(5, ao5), gen_avg_tab(12, ao12), gen_avg_tab(50, ao50), gen_avg_tab(100, ao100)))
}

func solvesTab() fyne.CanvasObject {
	return widget.NewList(
		func() int {
			return len(manageFile())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Solves")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(manageFile()[i])
		})
}

func statsTab() fyne.CanvasObject {
	// content := widget.NewLabel("Your stats should be here")
	return widget.NewLabel("Your stats should be here")
}

func gen_avg_tab(num uint8, data binding.String) fyne.CanvasObject {
	switch num {
	case 5:
		widget := widget.NewLabelWithData(data)
		widget.TextStyle.Bold = true
		widget.Alignment = fyne.TextAlignTrailing
		return widget
	case 12:
		widget := widget.NewLabelWithData(data)
		widget.TextStyle.Bold = true
		widget.Alignment = fyne.TextAlignTrailing
		return widget
	case 50:
		widget := widget.NewLabelWithData(data)
		widget.TextStyle.Bold = true
		widget.Alignment = fyne.TextAlignTrailing
		return widget
	case 100:
		widget := widget.NewLabelWithData(data)
		widget.TextStyle.Bold = true
		widget.Alignment = fyne.TextAlignTrailing
		return widget
	}
	var widget fyne.CanvasObject
	return widget
}

func gen_scramble_tab() fyne.CanvasObject {
	scramble := getScramble()
	return scramble
}
