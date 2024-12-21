package main

import "fmt"

// We will be modifying this later
func VariableDeclaration() string{
	return "🐛"
}

func StringConcatenation(a, b string) string {
	return a + b
}

func IncrementInt(a int) int {
	return a + 1
}

// Where we will run our code
func main() {
	variablesDeclaration := VariableDeclaration()
	strContat := StringConcatenation("⭐", "👀")
	incInt := IncrementInt(2)
	fmt.Println(variablesDeclaration)
	fmt.Println(strContat)
	fmt.Println(incInt)
}
