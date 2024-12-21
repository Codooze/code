package main

type Person struct {
	Name string
	Employed bool
}

// We will be modifying this later
func UpdateNameByValue(p Person, newName string) Person {
    p.Name = newName
    return p
}

func UpdateNameByReference(p *Person, newName string) {
	p.Name = newName
}

func MakePersonEmployed(p *Person) {
	p.Employed = true
}

// Where we will run our code
func main() {}
