package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("hw02unpackstring - function main")
	str, err := Unpack("aaa0b")
	fmt.Println("str=", str)
	fmt.Println("err=", err)
}

var ErrInvalidString = errors.New("invalid string")

/*
Функция срезает последний символ строки и возвращает оставшуюся строку.
*/
func RemoveChar(word string) (result string) {
	length := utf8.RuneCountInString(word)

	switch {
	case length == 0 || length == 1:
		result = ""
	case length == 2:
		result = word[0:1]
	case length > 2:
		result = word[0 : length-1]
	default:
		result = ""
	}

	return
}

func Unpack(str string) (string, error) {
	// Place your code here.
	fmt.Println("hw02unpackstring - function Unpack")

	arr := strings.Split(str, "")
	length := len(arr)

	if length == 0 {
		return "", nil
	}

	// Если 0-й элемент явл-ся цифрой, то выходим с ошибкой
	_, err := strconv.Atoi(arr[0])

	if err == nil {
		return "", ErrInvalidString
	}

	var sb strings.Builder
	sb.WriteString(arr[0])

	var isNumber bool // Символ явл-ся цифрой
	for i := 1; i < length; i++ {
		number, err := strconv.Atoi(arr[i])

		// Тут значение isNumber еще от прошлой итерации. Если прошлый символ число и текущий число, то выходим с ошибкой
		if err == nil && isNumber {
			return "", ErrInvalidString
		}

		if err == nil {
			isNumber = true
		} else {
			isNumber = false
		}

		// Если символ не цифра, то укажем его в результирующей строке
		if err != nil {
			sb.WriteString(arr[i])
		} else { // Если же цифра, то повторим прошлый элемент кол-во раз = цифре
			// Если цифра = 0, то надо срезать последний символ
			if number == 0 {
				var tmpStr string = sb.String()
				sb.Reset()
				sb.WriteString(RemoveChar(tmpStr))
			} else {
				// Если не 0, то повторить number-1 раз прошлый символ
				sb.WriteString(strings.Repeat(arr[i-1], number-1))
			}
		}
	}

	//fmt.Println("sb.String()=", sb.String())

	return sb.String(), nil
}
