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
	Unpack("d\n5abc")
}

var ErrInvalidString = errors.New("invalid string")

func RemoveChar(word string) (result string) {
	length := utf8.RuneCountInString(word)
	if length == 0 || length == 1 {
		result = ""
	} else if length == 2 {
		result = word[0:1]
	} else if length > 2 {
		result = word[0 : length-1]
	}
	return
}

func Unpack(str string) (string, error) {
	// Place your code here.
	fmt.Println("hw02unpackstring - function Unpack")

	//tmpstr := RemoveChar("a2")

	//tmpstr := strings.Repeat("a", 0)
	//fmt.Println("tmpstr=", tmpstr)

	arr := strings.Split(str, "")
	length := len(arr)

	//fmt.Println("length=", length)

	//Если 0-й элемент явл-ся цифрой, то выходим с ошибкой
	_, err := strconv.Atoi(arr[0])

	//fmt.Println("Symbol=", number)
	//fmt.Println("Error=", err)

	if err == nil {
		//fmt.Println("Exit by Error")
		return "", err
	}

	var resultStr string = arr[0]
	var isNumber bool // Символ явл-ся цифрой
	for i := 1; i < length; i++ {
		number, err := strconv.Atoi(arr[i])

		//fmt.Println("Cycle Symbol=", number)
		//fmt.Println("Cycle Error=", err)

		// Тут значение isNumber еще от прошлой итерации. Если прошлый символ число и текущий число, то выходим с ошибкой
		if err == nil && isNumber == true {
			return "", err
		}

		if err == nil {
			isNumber = true
		} else {
			isNumber = false
		}

		// Если символ не цифра, то укажем его в результирующей строке
		if err != nil {
			resultStr = resultStr + arr[i]
		} else { // Если же цифра, то повторим прошлый элемент кол-во раз = цифре
			// Если цифра = 0, то надо срезать последний символ
			if number == 0 {
				resultStr = RemoveChar(resultStr)
			} else {
				// Если не 0, то повторить number-1 раз прошлый символ
				resultStr = resultStr + strings.Repeat(arr[i-1], number-1)
			}
		}

	}

	// fmt.Println(arr)
	fmt.Println("resultStr=", resultStr)

	return resultStr, nil
}
