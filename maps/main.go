package main

import "fmt"

// We will be modifying this later
func CreateMap(key1, key2, val1, val2 string) map[string]string {
        return map[string]string{
                key1: val1,
                key2: val2,
        }
}

func AddToMap(myMap map[string]string, key, val string) map[string]string {
        myMap[key] = val
        return myMap
}

func DeleteFromMap(myMap map[string]string, key string) map[string]string {
        delete(myMap, key)
        return myMap
}

// Where we will run our code
func main() {
        // Create a map
        myMap := CreateMap("key1", "key2", "value1", "value2")
        fmt.Println("Created Map:", myMap)

        // Add a new key-value pair
        myMap = AddToMap(myMap, "key3", "value3")
        fmt.Println("Map after adding:", myMap)

        // Delete a key-value pair
        myMap = DeleteFromMap(myMap, "key2")
        fmt.Println("Map after deleting:", myMap)
}