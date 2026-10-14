package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
)

func printTypes(
	numDecimal, numOctal, numHexadecimal int,
	pi float64,
	name string,
	isActive bool,
	complexNum complex64,
) {
	fmt.Printf("numDecimal:     %s = %d\n", reflect.TypeOf(numDecimal), numDecimal)
	fmt.Printf("numOctal:       %s = %#o\n", reflect.TypeOf(numOctal), numOctal)
	fmt.Printf("numHexadecimal: %s = %#x\n", reflect.TypeOf(numHexadecimal), numHexadecimal)
	fmt.Printf("pi:             %s = %g\n", reflect.TypeOf(pi), pi)
	fmt.Printf("name:           %s = %q\n", reflect.TypeOf(name), name)
	fmt.Printf("isActive:       %s = %t\n", reflect.TypeOf(isActive), isActive)
	fmt.Printf("complexNum:     %s = %v\n", reflect.TypeOf(complexNum), complexNum)
}

func toString(
	numDecimal, numOctal, numHexadecimal int,
	pi float64,
	name string,
	isActive bool,
	complexNum complex64,
) string {
	return strconv.Itoa(numDecimal) +
		strconv.FormatInt(int64(numOctal), 8) +
		strconv.FormatInt(int64(numHexadecimal), 16) +
		strconv.FormatFloat(pi, 'f', -1, 64) +
		name +
		strconv.FormatBool(isActive) +
		strconv.FormatComplex(complex128(complexNum), 'f', 0, 64)
}

func toRunes(s string) []rune {
	return []rune(s)
}

func hashWithSalt(runes []rune) string {
	salt := "go-2024"
	middle := len(runes) / 2

	withSalt := make([]rune, 0, len(runes)+len(salt))
	withSalt = append(withSalt, runes[:middle]...)
	withSalt = append(withSalt, []rune(salt)...)
	withSalt = append(withSalt, runes[middle:]...)

	hash := sha256.Sum256([]byte(string(withSalt)))
	return fmt.Sprintf("%x", hash)
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	printTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	combined := toString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nОбъединенная строка:", combined)

	runes := toRunes(combined)
	fmt.Printf("Срез рун: %v\n", runes)

	hash := hashWithSalt(runes)
	fmt.Println("SHA256 с солью \"go-2024\":", hash)
}
