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
		container.NewTabItem("Solves", container.NewStack(solvesTab())),
		// container.NewTabItem("Stats", container.NewCenter(statsTab())),
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
			ao5.Set(getAverage(5, readTimes()))
			ao12.Set(getAverage(12, readTimes()))
			ao50.Set(getAverage(50, readTimes()))
			ao100.Set(getAverage(100, readTimes()))
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	return container.New(layout.NewGridLayoutWithRows(3), container.NewCenter(gen_scramble_cont()), container.NewCenter(timer), container.NewGridWithRows(4, gen_avg_cont(5, ao5), gen_avg_cont(12, ao12), gen_avg_cont(50, ao50), gen_avg_cont(100, ao100)))
}

func solvesTab() fyne.CanvasObject {
	list1 := widget.NewList(
		func() int {
			return len(readTimes())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Solves")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(readTimes()[i])
		})
	list2 := widget.NewList(
		func() int {
			return len(readScrambles())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Solves")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(readScrambles()[i])
		})
	cont1 := container.NewHBox(list1)
	cont2 := container.NewStack(list2)
	cont := container.NewHSplit(cont1, cont2)
	cont.SetOffset(0.0)
	return cont
}

// func statsTab() fyne.CanvasObject {
// 	return widget.NewLabel("Your stats should be here")
// }

func gen_avg_cont(num uint8, data binding.String) fyne.CanvasObject {
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

func gen_scramble_cont() fyne.CanvasObject {
	data := binding.NewString()
	go func() {
		for true {
			if timesSaved == true {
				scr := getScramble()
				data.Set(scr)
				currentScramble = scr
				timesSaved = false
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	wid := widget.NewLabelWithData(data)
	return wid
}
