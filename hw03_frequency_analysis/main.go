package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hw03_frequency_analysis - function main")

	// inputSrc := make([]WordFreq, 6)
	// inputSrc[0] = WordFreq{"boy", 2}
	// inputSrc[1] = WordFreq{"dog", 2}
	// inputSrc[2] = WordFreq{"and", 2}
	// inputSrc[3] = WordFreq{"girl", 3}
	// inputSrc[4] = WordFreq{"girll", 3}
	// inputSrc[5] = WordFreq{"abc", 1}
	//lexycoGraphicSort(inputSrc)

	//inputStr := "boy dog cat and"
	inputStr := "cat and dog, one dog,two cats and one man and dog"
	//inputStr := ""

	/*
		var inputStr = `Как видите, он  спускается  по  лестнице  вслед  за  своим
		другом   Кристофером   Робином,   головой   вниз,  пересчитывая
		ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
		сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
			кажется, что можно бы найти какой-то другой способ, если бы  он
		только   мог   на  минутку  перестать  бумкать  и  как  следует
		сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
			Как бы то ни было, вот он уже спустился  и  готов  с  вами
		познакомиться.
		- Винни-Пух. Очень приятно!
			Вас,  вероятно,  удивляет, почему его так странно зовут, а
		если вы знаете английский, то вы удивитесь еще больше.
			Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
		вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
		лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
		очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
		громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
		можешь  сделать вид, что ты просто понарошку стрелял; а если ты
		звал его тихо, то все подумают, что ты  просто  подул  себе  на
		нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
		Робин решил отдать его своему медвежонку, чтобы оно не  пропало
		зря.
			А  Винни - так звали самую лучшую, самую добрую медведицу
		в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
		Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
		честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
		знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
		забыл.
			Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
			Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
		иногда,  особенно  когда  папа  дома,  он больше любит тихонько
		посидеть у огня и послушать какую-нибудь интересную сказку.
			В этот вечер...`
	*/

	res := Top10(inputStr)
	fmt.Println("Топ 10 слов: ", res)
}

type WordFreq struct {
	word  string // Слово
	count int    // Кол-во вхождений слова
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

/*
Функция проверяет, содержит ли переданный массив кол-во указания count
*/
func isRepeatOfCountInArray(count int, inputArr []WordFreq) bool {
	res := false
	arrLength := len(inputArr)

	for i := 0; i < arrLength; i++ {
		if inputArr[i].count == count {
			return true
		}
	}

	return res
}

/*
Функция возвращает массив, отсортированный лексикографически внутри тех слайсов элементов,
для которых частота указания в тексте совпадает.
*/
func lexycoGraphicSort(inputArr []WordFreq) []WordFreq {
	outputArr := make([]WordFreq, 0)

	for i := 0; i < len(inputArr); i++ {
		currentCount := inputArr[i].count

		if isRepeatOfCountInArray(currentCount, outputArr) {
			continue
		}

		currentIndexStart := i // Самый ранний номер элемента с данным значением currentCount (частоты указания слова)
		currentIndexEnd := i   // Самый поздний номер элемента с данным значением currentCount (частоты указания слова)
		currentWords := make([]string, 0)

		for j := i + 1; j < len(inputArr); j++ {
			if currentCount == inputArr[j].count {
				currentIndexEnd = j
			}
		}

		// Лексикографическая сортировка нужна лишь в том случае, если есть несколько элементов (слов) с одинаковым значением частоты использования
		if currentIndexEnd > currentIndexStart {
			for k := currentIndexStart; k <= currentIndexEnd; k++ {
				currentWords = append(currentWords, inputArr[k].word)
			}
			sort.Strings(currentWords)

			// Заменим элемент массива, чтобы соблюдалась сортировка слов
			for k := currentIndexStart; k <= currentIndexEnd; k++ {
				//outputArr[k] = WordFreq{currentWords[k], currentCount}
				outputArr = append(outputArr, WordFreq{currentWords[k-currentIndexStart], currentCount})
			}
		} else {
			outputArr = append(outputArr, inputArr[i])
		}
	}

	return outputArr
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

		wordFreqArray = append(wordFreqArray, WordFreq{elementValue, count})
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

	// Лексикографическая сортировка
	lexicalSortedTopWordFreqArray := lexycoGraphicSort(topWordFreqArray)

	var res []string
	for k := 0; k < resultLength; k++ {
		res = append(res, lexicalSortedTopWordFreqArray[k].word)
	}

	return res
}
