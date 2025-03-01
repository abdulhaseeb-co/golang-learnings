package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Create a slice with length 2 and capacity 4
	s := make([]int, 2, 4)
	s[0] = 10
	s[1] = 20

	fmt.Println("Initial Slice:")
	printSliceInfo(s)

	// Store the initial address to compare later
	initialAddress := getSliceDataAddress(s)

	// Append an element (within capacity)
	s = append(s, 30)

	fmt.Println("\nAfter appending within capacity:")
	printSliceInfo(s)
	currentAddress := getSliceDataAddress(s)

	if initialAddress == currentAddress {
		fmt.Println("Underlying array address: Remains the same (no reallocation)")
	} else {
		fmt.Println("Underlying array address: Changed (reallocation occurred - which should not happen)")
	}
}

func printSliceInfo(s []int) {
	fmt.Printf("Slice: %v\n", s)
	fmt.Printf("Length: %d\n", len(s))
	fmt.Printf("Capacity: %d\n", cap(s))
	fmt.Printf("Data Address (approximate): %p\n", getSliceDataAddress(s))
}

func getSliceDataAddress(s []int) unsafe.Pointer {
	if len(s) == 0 {
		return nil // Avoid panic for empty slices
	}
	return unsafe.Pointer(&s[0])
}