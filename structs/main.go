package main

type Person struct {
	Name string
	Employed bool
}

// We will be modifying this later
func UpdateNameByValue(user Person, newName string) Person{
	user.Name = newName
	return user
}

func UpdateNameByReference(user *Person, newName string) {
	user.Name = newName
}

func MakePersonEmployed(user *Person) {
	user.Employed = true
}

// Where we will run our code
func main() {}
