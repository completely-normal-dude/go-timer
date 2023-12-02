package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
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
					result = "Ao5  " + strconv.FormatFloat(number/3, 'f', 2, 64)
					return result
				}
			}
		} else {
			result = "Ao5   -"
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
					result = "Ao12  " + strconv.FormatFloat(number/10, 'f', 2, 64)
					return result
				}
			}
		} else {
			result = "Ao12   -"
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
					result = "Ao50  " + strconv.FormatFloat(number/48, 'f', 2, 64)
					return result
				}
			}
		} else {
			result = "Ao50   -"
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
					result = "Ao100  " + strconv.FormatFloat(number/98, 'f', 2, 64)
					return result
				}
			}
		} else {
			result = "Ao100   -"
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
					result = "AoAll  " + strconv.FormatFloat(number/(float64(length)-2), 'f', 2, 64)
					return result
				}
			}
		} else {
			result = "AoAll   -"
		}
	}
	return result
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
					// timer.SetText(t)
					a++
					seconds += 0.01
					if a == 10 {
						t := strconv.FormatFloat(seconds, 'f', 1, 64)
						timer.SetText(t)
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
		timer.SetText(t)
		data, _ := os.ReadFile(timesPath)
		fmtdata := string(data)
		timesSaved = true
		if fmtdata == "" {
			if solve < 10 {
				fmt.Fprintf(timesFile, "0%s", t)
				fmt.Fprintf(scramblesFile, "%s", currentScramble)
				fmt.Printf("Saved 0%s!\n", t)
			} else {
				fmt.Fprintf(timesFile, "%s", t)
				fmt.Fprintf(scramblesFile, "%s", currentScramble)
				fmt.Printf("Saved %s!\n", t)
			}
		} else {
			if solve < 10 {
				fmt.Fprintf(timesFile, "\n0%s", t)
				fmt.Fprintf(scramblesFile, "\n%s", currentScramble)
				fmt.Printf("Saved 0%s!\n", t)
			} else {
				fmt.Fprintf(timesFile, "\n%s", t)
				fmt.Fprintf(scramblesFile, "\n%s", currentScramble)
				fmt.Printf("Saved %s!\n", t)
			}
		}
	}
}
func readTimes() []string {
	data, _ := os.ReadFile(timesPath)
	fmtdata := string(data)
	content := strings.Split(fmtdata, "\n")
	slices.Reverse(content)
	return content
}
func readScrambles() []string {
	data, _ := os.ReadFile(scramblesPath)
	fmtdata := string(data)
	content := strings.Split(fmtdata, "\n")
	slices.Reverse(content)
	return content
}

func getScramble() string {
	vm := otto.New()
	vm.Run(`Array.prototype.choose = function() {
		var index = Math.floor(Math.random() * this.length);
		return this[index];
};

function scramble(length) {
		var planes = {x: ['L', 'R'], y: ['U', 'D'], z: ['F', 'B']};
		var planeMap = {};
		for (var plane in planes) {
			var sides = planes[plane];
			for (var i = 0; i < sides.length; i++) {
				var side = sides[i];
				planeMap[side] = plane;
			}
		}

		var sides = ['F', 'B', 'R', 'L', 'U', 'D'];
		var modifiers = ['2', '\'', ''];

		var weakBuffer = [], moves = [];
		for (var i = 0; i < length; i++) {
			var mod = modifiers.choose(), side;
			if (weakBuffer.length == 0) {
				side = sides.choose();
			}
			else if (weakBuffer.length == 1) {
				var badSide = weakBuffer[0],
				newSides = sides.slice(),
				badIndex = newSides.indexOf(badSide);
				newSides.splice(badIndex, 1);

				side = newSides.choose();

				if (planeMap[side] != planeMap[badSide]) {
					weakBuffer = []; 
				}
			}
			else {
				var newSides = sides.slice();
				for (var a = 0; a < weakBuffer.length; a++) {
					var badSide = weakBuffer[a],
					badIndex = newSides.indexOf(badSide);
					newSides.splice(badIndex, 1);
				}

				side = newSides.choose();

				weakBuffer = [];
			}
			moves.push(side + mod);
			weakBuffer.push(side);
		}
		return moves.join(' ');
	}
var scramble = scramble(20);
		`)
	value, _ := vm.Get("scramble")
	text, _ := value.ToString()
	return text
}
