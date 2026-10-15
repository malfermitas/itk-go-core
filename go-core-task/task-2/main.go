package main

import (
	"fmt"
	"math/rand"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func sliceExample[T Numeric](s []T) []T {
	newSlice := make([]T, 0)

	for _, num := range s {
		if num%2 == 0 {
			newSlice = append(newSlice, num)
		}
	}

	return newSlice
}

func addElements[T Numeric](dest []T, num T) []T {
	newSlice := make([]T, len(dest))
	copy(newSlice, dest)
	newSlice = append(newSlice, num)
	return newSlice
}

func copySlice[T Numeric](input []T) []T {
	newSlice := make([]T, len(input))
	copy(newSlice, input)
	return newSlice
}

func removeElement[T Numeric](input []T, index int) []T {
	if index < 0 || index >= len(input) {
		return input
	}
	newSlice := make([]T, len(input))
	copy(newSlice, input)
	return append(newSlice[:index], newSlice[index+1:]...)
}

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(50)
	}

	fmt.Println("Original Slice:", originalSlice)
	// 2. Тест sliceExample (фильтрация четных)
	evenNumbers := sliceExample(originalSlice)
	fmt.Println("Even Numbers Only:", evenNumbers)
	// 3. Тест addElements (добавление числа)
	addedSlice := addElements(originalSlice, 999)
	fmt.Println("After adding 999:", addedSlice)
	// 4. Тест copySlice (копирование)
	copiedSlice := copySlice(originalSlice)
	// Изменяем копию для проверки независимости
	copiedSlice[0] = -1
	fmt.Println("Original after copying & modifying copy:", originalSlice)
	fmt.Println("Copied slice (modified):", copiedSlice)
	// 5. Тест removeElement (удаление по индексу, например, индекс 2)
	indexToRemove := 2
	removedSlice := removeElement(originalSlice, indexToRemove)
	fmt.Printf("Original Slice: %v\n", originalSlice)
	fmt.Printf("After removing element at index %d: %v\n", indexToRemove, removedSlice)
}
