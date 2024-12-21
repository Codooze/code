package main

import (
	"fmt"
	"slices"
)

// We will be modifying this later
func CreateSlice() []string{
	return []string{"dog","cat","monkey"}
}

func ModifySlice(sl []string) []string {
        if len(sl) < 3 {
                return sl
        }
        // Create a new slice with a copy of the original elements
        newSlice := make([]string, len(sl))
        copy(newSlice, sl) 

        newSlice[1] = "❤️"
        return newSlice
}

func PopSliceValue(sl []string) []string{
	 return slices.Delete(sl, 2,3)
	 
}

// Where we will run our code
func main() {
	originalSlice := CreateSlice() // Create the initial slice
        modifiedSlice := ModifySlice(originalSlice) // Create a modified copy

        fmt.Println("Original Slice:", originalSlice)  // Output: [dog cat monkey]
        fmt.Println("Modified Slice:", modifiedSlice) // Output: [dog ❤️ monkey]
}
