package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"os"
	"slices"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func setTabs(w fyne.Window) *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Timer", container.NewVBox(timerTab())),
		container.NewTabItem("Solves", container.NewStack(solvesTab(w))),
		// container.NewTabItem("Stats", container.NewCenter(statsTab())),
	)
	return tabs
}

func timerTab() fyne.CanvasObject {
	ao5, ao12, ao50, ao100, aoAll := binding.NewString(), binding.NewString(), binding.NewString(), binding.NewString(), binding.NewString()
	go func() {
		for range time.Tick(time.Millisecond * 500) {
			newSlice := readFile(0)
			slices.Reverse(newSlice)
			ao5.Set(getAverage(5, newSlice))
			ao12.Set(getAverage(12, newSlice))
			ao50.Set(getAverage(50, newSlice))
			ao100.Set(getAverage(100, newSlice))
			aoAll.Set(getAverage(0, newSlice))
		}
	}()
	time.Sleep(time.Millisecond * 500)
	return container.New(layout.NewGridLayoutWithRows(3), container.NewCenter(gen_scramble_cont()), container.NewCenter(timer), container.NewGridWithRows(5, gen_avg_cont(5, ao5), gen_avg_cont(12, ao12), gen_avg_cont(50, ao50), gen_avg_cont(100, ao100), gen_avg_cont(0, aoAll)))
}

func solvesTab(window fyne.Window) fyne.CanvasObject {
	list1 := widget.NewList(
		func() int {
			return len(decodeFile())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Solves")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(decodeFile()[i][0])
		})
	list2 := widget.NewList(
		func() int {
			return len(decodeFile())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Solves")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(decodeFile()[i][1])
		})
	erase := widget.NewButton("Erase all", func() {
		dialog.ShowConfirm("Confirmation", "Are you sure you want to delete all saved times?\nThis cannot be undone!", func(choice bool) {
			if choice {
				os.Truncate(filePath, 0)
				list1.Refresh()
				list2.Refresh()
				fmt.Println("Erased all times")
			}
		}, window)
	})
	list1.OnSelected = func(id widget.ListItemID) {
		dialog.ShowConfirm("Confirmation", "Are you sure you want to delete this time?\nThis cannot be undone!", func(choice bool) {
			if choice {
				var newSlice [][]string
				newSlice = slices.Delete(decodeFile(), id, id+1)
				if len(newSlice) == 0 {
					os.Truncate(filePath, 0)
					list1.Refresh()
					list2.Refresh()
					list1.UnselectAll()
					return
				}
				os.Truncate(filePath, 0)
				f, _ := os.OpenFile(filePath, os.O_WRONLY, 0644)
				writer := csv.NewWriter(f)
				writer.WriteAll(newSlice)
				defer f.Close()
				list1.Refresh()
				list2.Refresh()
				list1.UnselectAll()
			}
		}, window)
		list1.UnselectAll()
	}
	list2.OnSelected = func(id widget.ListItemID) {
		dialog.ShowConfirm("Confirmation", "Are you sure you want to delete this time?\nThis cannot be undone!", func(choice bool) {
			if choice {
				var newSlice [][]string
				newSlice = slices.Delete(decodeFile(), id, id+1)
				if len(newSlice) == 0 {
					os.Truncate(filePath, 0)
					list2.Refresh()
					list1.Refresh()
					list2.UnselectAll()
					return
				}
				os.Truncate(filePath, 0)
				f, _ := os.OpenFile(filePath, os.O_WRONLY, 0644)
				writer := csv.NewWriter(f)
				writer.WriteAll(newSlice)
				defer f.Close()
				list1.Refresh()
				list2.Refresh()
				list2.UnselectAll()
			}
		}, window)
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
	currentScramble = getScramble()
	go func() {
		for range time.Tick(time.Millisecond * 500) {
			if timesSaved {
				scr := getScramble()
				scrambleText.Text = scr
				scrambleText.Refresh()
				currentScramble = scr
				timesSaved = false
			}
		}
	}()
	scrambleText = canvas.NewText(currentScramble, color.White)
	scrambleText.TextSize = 20
	return scrambleText
}
