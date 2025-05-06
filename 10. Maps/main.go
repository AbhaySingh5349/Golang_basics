package main

import "fmt"

func main() {
	// var mp map[KeyType]ValueType

	// mp := make(map[KeyType][ValueType])

	// mp := map[KeyType][ValueType]{
	// 	k1: v1,
	// 	k2: v2
	// }

	mp := make(map[string]int)

	mp["k1"] = 1
	mp["k2"] = 2
	mp["k3"] = 3

	fmt.Println("len of map: ", len(mp))

	for k, v := range mp {
		fmt.Println(k, " : ", v)
	}

	fmt.Println(mp["k4"]) // returns default value of "valueType"

	delete(mp, "k1")

	fmt.Println(mp)

	val, ok := mp["k2"]
	fmt.Println("val: ", val, " , keyExists: ", ok)

	clear(mp)

	fmt.Println(mp)

	nested_mp := make(map[string]map[string]int)
	nested_mp["k1"] = mp
	nested_mp["k2"] = mp

	fmt.Println(nested_mp)
}
