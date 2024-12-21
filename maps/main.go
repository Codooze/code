package main

// We will be modifying this later
func CreateMap(key1, key2, val1, val2 string) map[string]string{
	return map[string]string{
		key1: val1,
		key2: val2,
	}
}

func AddToMap(myMap map[string]string, key, val string) map[string]string{
	myMap[key] = val
	return myMap
}

func DeleteFromMap(myMap map[string]string, key string) map[string]string{
	delete(myMap, key)
	return myMap
}

// Where we will run our code
func main() {}
