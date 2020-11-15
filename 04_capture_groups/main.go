package main

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

type Task struct {
	str   string
	pairs []Pair
}

type Pair struct {
	key   string
	value string
}

func main() {
	// Написать функцию, которая будет доставать пары ключ=значение из текста
	// Ключ и значение состоят из латинских букв и цифр
	// Ключ не может начинаться с цифры
	// Если он начинается с цифры, он невалиден, не возвращаем его
	re := regexp.MustCompile(``)

	yourFunc := func(str string) (pairs []Pair) {
		matches := re.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			pairs = append(pairs, Pair{key: match[1], value: match[2]})
		}
		return pairs
	}

	// Код ниже не трогать
	tasks := []Task{
		{str: "key1=value1",
			pairs: []Pair{
				{key: "key1", value: "value1"},
			}},
		{str: "key1=value1, key2=value2",
			pairs: []Pair{
				{key: "key1", value: "value1"},
				{key: "key2", value: "value2"},
			}},
		{str: "some text with keys=values inside it, when (point=v1)",
			pairs: []Pair{
				{key: "keys", value: "values"},
				{key: "point", value: "v1"},
			}},
		{str: "checking number rule 1key=2value, but correct is key2=value1",
			pairs: []Pair{
				{key: "key2", value: "value1"},
			}},
	}

	for _, task := range tasks {
		pairs := yourFunc(task.str)
		if assert.ObjectsAreEqual(pairs, task.pairs) {
			color.Green(task.str)
		} else {
			color.Red(task.str)
		}
		fmt.Println("   ", len(pairs), pairs)
	}
}
