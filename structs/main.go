package main

import "fmt"

type Person struct {
  Name string
  Employed bool
}

// Modify name by value (modifies a copy)
func UpdateNameByValue(p Person, newName string) Person {
  p.Name = newName
  return p
}

// Modify name by reference (modifies the original)
func UpdateNameByReference(p *Person, newName string) {
  p.Name = newName
}

// Set employed status to true
func MakePersonEmployed(p *Person) {
  p.Employed = true
}

// Where we will run our code
func main() {
  // Create a Person struct
  person := Person{Name: "John Doe", Employed: false}

  // Example 1: UpdateNameByValue (modifies a copy)
  updatedPerson := UpdateNameByValue(person, "Jane Doe") // Creates a new Person with the updated name
  fmt.Println("After UpdateNameByValue (copy):", person.Name, updatedPerson.Name) // Output: John Doe Jane Doe (original remains unchanged)

  // Example 2: UpdateNameByReference (modifies the original)
  UpdateNameByReference(&person, "Alice") // Pass the address of the person to modify the original
  fmt.Println("After UpdateNameByReference:", person.Name) // Output: Alice (original is modified)

  // Example 3: MakePersonEmployed (modifies the original)
  MakePersonEmployed(&person)
  fmt.Println("After MakePersonEmployed:", person.Employed) // Output: true (original is modified)
}