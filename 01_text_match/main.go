package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`rober`)
	texts := []string{
		"tom ford",
		"bob green",
		"mark rober",
		"lily valley",
	}

	for _, str := range texts {
		fmt.Println(str, "—", re.MatchString(str))
	}
}
