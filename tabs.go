package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
	scramble := widget.NewLabel(getScramble())
	scramble.Alignment = fyne.TextAlignCenter
	return container.New(layout.NewGridLayoutWithRows(3), container.NewCenter(scramble), container.NewCenter(timer), container.NewGridWithRows(4, gen_avg_tab(5), gen_avg_tab(12), gen_avg_tab(50), gen_avg_tab(100)))
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
	content := widget.NewLabel("Your stats should be here")
	return content
}
