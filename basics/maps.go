package basics

import (
	"fmt"
	"maps"
)

func mapsFunc() {
	// var m map[keyType]valueType

	//mapVariable = make(map[keyType]valueType)

	//using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)

	fmt.Println("Initial map:", myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println("Map after adding key1:", myMap)
	fmt.Println("Value for 'key1':", myMap["key1"])
	fmt.Println("Not existing key:", myMap["key"])

	myMap["code"] = 21
	fmt.Println("Map after updating 'code':", myMap)

	delete(myMap, "key1")
	fmt.Println("Map after deleting 'key1':", myMap)

	myMap["newKey"] = 42
	myMap["anotherKey"] = 100
	myMap["key2"] = 200
	fmt.Println("Map after adding more keys:", myMap)

	// clear(myMap)
	// fmt.Println("Map after clearing:", myMap)

	_, ok := myMap["key2"]
	fmt.Println("is exists:", ok)

	if ok {
		fmt.Println("Key 'key2' exists in the map")
	} else {
		fmt.Println("Key 'key2' does not exist in the map")
	}

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Map 2:", myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	for k, v := range myMap3 {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("myMap4 is nil")
	} else {
		fmt.Println("myMap4 is not nil")
	}

	val := myMap4["nonExistentKey"]
	fmt.Println("Value for 'nonExistentKey':", val)

	// myMap4["key"] = "Value"
	// fmt.Println("Map 4 after adding a key:", myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println("Map 4 after adding a key:", myMap4)

	fmt.Println("Length of myMap4:", len(myMap))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println("Nested map:", myMap5)
}