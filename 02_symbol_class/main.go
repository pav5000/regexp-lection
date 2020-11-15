package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Написать регулярку, которая будет полностью матчить все слова списка
	yourPattern := ``

	// Код ниже не трогать
	re := regexp.MustCompile(`^` + yourPattern + `$`)
	texts := []string{
		"cat",
		"car",
		"bat",
		"rat",
	}

	for _, str := range texts {
		fmt.Println(str, "—", re.MatchString(str))
	}
}
