package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hw03_frequency_analysis - function main")
}

func Top10(inputStr string) []string {
	arr := strings.Split(inputStr, " ")
	length := len(arr)

	// Каждый перебираемый элемент цикла - это слово
	for i := 0; i < length; i++ {
		elementValue := arr[i]
		count := 1

		// Посчитаем, сколько раз слово (elementValue) встречается в тексте
		// Элементы массива левее i не анализируем, т.к. мы их уже обработали
		for j := i + 1; j < length; j++ {
			if elementValue == arr[j] {
				count = count + 1
			}
		}
	}

	return nil
}
