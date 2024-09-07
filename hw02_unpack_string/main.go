package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hw02unpackstring - function main")
	Unpack("s1")
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(_ string) (string, error) {
	// Place your code here.
	fmt.Println("hw02unpackstring - function Unpack")
	return "", nil
}
