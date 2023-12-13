package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func getAverage(ao uint8, solves []string) (result string) {
	var number float64 = 0
	switch ao {
	case 5:
		if len(solves) >= 5 {
			slice1 := solves[0:5]
			slices.Sort(slice1)
			solvesSlice := slice1[1:4]
			for i := 0; i < 3; i++ {
				a, _ := strconv.ParseFloat(solvesSlice[i], 64)
				number += a
				if i == 2 {
					result = "Ao5   " + strconv.FormatFloat(number/3, 'f', 2, 64)
					return
				}
			}
		} else {
			result = "Ao5          - "
		}

	case 12:
		if len(solves) >= 12 {
			slice1 := solves[0:12]
			slices.Sort(slice1)
			solvesSlice := slice1[1:11]
			for i := 0; i < 10; i++ {
				a, _ := strconv.ParseFloat(solvesSlice[i], 64)
				number += a
				if i == 9 {
					result = "Ao12    " + strconv.FormatFloat(number/10, 'f', 2, 64)
					return
				}
			}
		} else {
			result = "Ao12          - "
		}

	case 50:
		if len(solves) >= 50 {
			slice1 := solves[0:50]
			slices.Sort(slice1)
			solvesSlice := slice1[1:49]
			for i := 0; i < 48; i++ {
				a, _ := strconv.ParseFloat(solvesSlice[i], 64)
				number += a
				if i == 47 {
					result = "Ao50    " + strconv.FormatFloat(number/48, 'f', 2, 64)
					return
				}
			}
		} else {
			result = "Ao50          - "
		}
	case 100:
		if len(solves) >= 100 {
			slice1 := solves[0:100]
			slices.Sort(slice1)
			solvesSlice := slice1[1:99]
			for i := 0; i < 98; i++ {
				a, _ := strconv.ParseFloat(solvesSlice[i], 64)
				number += a
				if i == 97 {
					result = "Ao100    " + strconv.FormatFloat(number/98, 'f', 2, 64)
					return
				}
			}
		} else {
			result = "Ao100          - "
		}
	case 0:
		length := len(solves)
		if length > 1 {
			slice1 := solves[0:length]
			slices.Sort(slice1)
			solvesSlice := slice1[1 : length-1]
			for i := 0; i < length-2; i++ {
				a, _ := strconv.ParseFloat(solvesSlice[i], 64)
				number += a
				if i == length-3 {
					if length < 10 {
						a := fmt.Sprintf("Ao%d  ", length)
						result = a + strconv.FormatFloat(number/(float64(length)-2), 'f', 2, 64)
						return
					}
					a := fmt.Sprintf("Ao%d   ", length)
					result = a + strconv.FormatFloat(number/(float64(length)-2), 'f', 2, 64)
					return
				}
			}
		} else {
			result = ""
		}
	}
	return
}

func startTimer(f bool) {
	var a uint8 = 0
	switch f {
	case true:
		go func() {
			timerRunning = true
			fmt.Printf("Timer started... ")
			for range time.Tick(10 * time.Millisecond) {
				// for range time.Tick(100 * time.Millisecond) {
				select {
				case <-ch:
					return
				default:
					// seconds += 0.1
					// t := strconv.FormatFloat(seconds, 'f', 1, 64)
					// timer.Text=t
					// timer.Refresh()
					a++
					seconds += 0.01
					if a == 10 {
						t := strconv.FormatFloat(seconds, 'f', 1, 64)
						timer.Text = t
						timer.Refresh()
						a = 0
					}
				}
			}
		}()

	default:
		ch <- true
		timerRunning = false
		solve := seconds
		seconds = 0
		t := strconv.FormatFloat(solve, 'f', 2, 64)
		timer.Text = t
		timer.Refresh()
		timesSaved = true
		f, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
		writer := csv.NewWriter(f)
		if solve < 10 {
			a := fmt.Sprintf("0%s", t)
			save := [][]string{
				{a, currentScramble},
			}
			writer.WriteAll(save)
			fmt.Printf("Saved 0%s!\n", t)
		} else {
			save := [][]string{
				{t, currentScramble},
			}
			writer.WriteAll(save)
			fmt.Printf("Saved %s!\n", t)
		}
	}
}

func readFile(Index uint8) (newSlice []string) {
	f, _ := os.Open(filePath)
	reader := csv.NewReader(f)
	data, _ := reader.ReadAll()
	for _, row := range data {
		newSlice = append(newSlice, row[Index])
	}
	return
}
func decodeFile() (data [][]string) {
	f, _ := os.Open(filePath)
	reader := csv.NewReader(f)
	data, _ = reader.ReadAll()
	slices.Reverse(data)
	return
}

func Choose(slice []string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rand.Intn(len(slice))
	return slice[index]
}

func Scramble(length int) string {
	// Generate map of planes associated with each side
	planes := map[string][]string{"x": {"L", "R"}, "y": {"U", "D"}, "z": {"F", "B"}}
	planeMap := make(map[string]string)

	for plane, sides := range planes {
		for _, side := range sides {
			planeMap[side] = plane
		}
	}

	sides := []string{"F", "B", "R", "L", "U", "D"}
	modifiers := []string{"2", "'", ""}

	// Create buffer of moved sides
	// Once a plane is crossed, a move on those sides is then permissible again
	var weakBuffer []string
	var moves []string

	for i := 0; i < length; i++ {
		mod := Choose(modifiers)
		var side string

		if len(weakBuffer) == 0 {
			side = Choose(sides)
		} else if len(weakBuffer) == 1 {
			badSide := weakBuffer[0]
			newSides := make([]string, len(sides))
			copy(newSides, sides)
			badIndex := -1

			for j, s := range newSides {
				if s == badSide {
					badIndex = j
					break
				}
			}

			newSides = append(newSides[:badIndex], newSides[badIndex+1:]...)

			side = Choose(newSides)

			if planeMap[side] != planeMap[badSide] {
				weakBuffer = nil // planes have been crossed
			}
		} else {
			// Double plane weakness
			// Neither side in the weak buffer can be chosen

			newSides := make([]string, len(sides))
			copy(newSides, sides)

			for _, badSide := range weakBuffer {
				badIndex := -1

				for j, s := range newSides {
					if s == badSide {
						badIndex = j
						break
					}
				}

				newSides = append(newSides[:badIndex], newSides[badIndex+1:]...)
			}

			side = Choose(newSides)
			weakBuffer = nil
		}

		moves = append(moves, side+mod)
		weakBuffer = append(weakBuffer, side)
	}

	return strings.Join(moves, " ")
}
