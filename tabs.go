package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
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
	aoAll := binding.NewString()
	go func() {
		for {
			ao5.Set(getAverage(5, readTimes()))
			ao12.Set(getAverage(12, readTimes()))
			ao50.Set(getAverage(50, readTimes()))
			ao100.Set(getAverage(100, readTimes()))
			aoAll.Set(getAverage(0, readTimes()))
			time.Sleep(time.Millisecond * 500)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	return container.New(layout.NewGridLayoutWithRows(3), container.NewCenter(gen_scramble_cont()), container.NewCenter(timer), container.NewGridWithRows(5, gen_avg_cont(5, ao5), gen_avg_cont(12, ao12), gen_avg_cont(50, ao50), gen_avg_cont(100, ao100), gen_avg_cont(0, aoAll)))
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
	erase := widget.NewButton("Erase all", func() {
		os.WriteFile(timesPath, []byte{}, os.FileMode(os.O_WRONLY))
		os.WriteFile(scramblesPath, []byte{}, os.FileMode(os.O_WRONLY))
		list1.Refresh()
		list2.Refresh()
		fmt.Println("Erased all times")
	})
	list1.OnSelected = func(id widget.ListItemID) {
		list1.Refresh()
		list1.UnselectAll()
	}
	list2.OnSelected = func(id widget.ListItemID) {
		list2.Refresh()
		list2.UnselectAll()
	}
	erasecont := container.NewHBox(erase)
	cont1 := container.NewHBox(list1)
	cont2 := container.NewStack(list2)
	cont := container.NewHSplit(cont1, cont2)
	cont.SetOffset(0.0)
	defcont := container.NewVSplit(erasecont, cont)
	defcont.SetOffset(0.0)
	return defcont
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
	case 0:
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
		for {
			if timesSaved {
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
