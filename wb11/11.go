package main

import "fmt"

func main() {

	// множество

	mp := make(map[string]int)

	arr1 := []string{"da", "net", "boom", "hello"}
	arr2 := []string{"da", "net"}
	arr3 := append(arr1, arr2...)

	var result []string

	for _, key := range arr3 {
		mp[key] += 1
		if mp[key] > 1 {
			result = append(result, key)
		}
	}

	fmt.Println("map: ", mp)
	fmt.Println("result: ", result)
}
