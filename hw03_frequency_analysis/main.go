package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hw03_frequency_analysis - function main")
	//inputStr := "cat and dog, one dog,two cats and one man and dog"
	inputStr := ""

	Top10(inputStr)
}

type WordFreq struct {
	number int    // Номер слова в исходном массиве
	word   string // Слово
	count  int    // Кол-во вхождений слова
}

/*
Функция проверяет, содержит ли переданный массив значение переданной строки
*/
func isRepeatOfWordInArray(inputWord string, inputArr []WordFreq) bool {
	res := false
	arrLength := len(inputArr)

	for i := 0; i < arrLength; i++ {
		if inputArr[i].word == inputWord {
			return true
		}
	}

	return res
}

func Top10(inputStr string) []string {
	if inputStr == "" {
		return nil
	}

	const topCnt = 10 // Число выдаваемых самых популярных слов
	arr := strings.Split(inputStr, " ")
	length := len(arr)

	var wordFreqArray []WordFreq

	// Каждый перебираемый элемент цикла - это слово
	for i := 0; i < length; i++ {
		elementValue := arr[i]
		if isRepeatOfWordInArray(elementValue, wordFreqArray) {
			continue
		}
		count := 1

		// Посчитаем, сколько раз слово (elementValue) встречается в тексте
		// Элементы массива левее i не анализируем, т.к. мы их уже обработали
		for j := i + 1; j < length; j++ {
			if elementValue == arr[j] {
				count = count + 1
			}
		}

		wordFreqArray = append(wordFreqArray, WordFreq{i, elementValue, count})
	}

	// Отсортируем данные по кол-ву вхождений каждого слова. Первые элементы - это самые популярные слова
	sort.Slice(wordFreqArray, func(i, j int) bool {
		return wordFreqArray[i].count > wordFreqArray[j].count
	})

	// Кол-во строк в выдаваемом ответе
	resultLength := topCnt
	if len(wordFreqArray) < topCnt {
		resultLength = len(wordFreqArray)
	}
	topWordFreqArray := wordFreqArray[0:resultLength]

	var res []string
	for k := 0; k < resultLength; k++ {
		res = append(res, topWordFreqArray[k].word)
	}

	// fmt.Println("Top10 - finish")
	// fmt.Println(topWordFreqArray)

	return res
}
