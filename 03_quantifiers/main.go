package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type Task struct {
	Str     string
	Count   int
	Matches []string
}

func main() {
	// Написать регулярку, которая будет матчить html-цвет
	yourPattern := ``

	// Код ниже не трогать
	re := regexp.MustCompile(yourPattern)
	tasks := []Task{
		{Str: "this color is red: #ff3322",
			Matches: []string{"#ff3322"}},
		{Str: "if we mix #12ccbb with #64cb8d, what do we get?",
			Matches: []string{"#12ccbb", "#64cb8d"}},
		{Str: "my favourite colors are: #ff0000 #00ff00 #0000ff",
			Matches: []string{"#ff0000", "#00ff00", "#0000ff"}},
		{Str: "no colors here",
			Matches: []string{}},
		{Str: "this color: #0f2c6 is not valid",
			Matches: []string{}},
		{Str: "#f7 — why is it not right?",
			Matches: []string{}},
		{Str: "oops, you forgot that #FFaabb can be uppercase",
			Matches: []string{"#FFaabb"}},
	}

	for _, task := range tasks {
		matches := re.FindAllString(task.Str, -1)
		if equalSlice(matches, task.Matches) {
			color.Green(task.Str)
		} else {
			color.Red(task.Str)
		}
		fmt.Println("   ", len(matches), strings.Join(matches, " "))
	}
}

func equalSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
