package main

import "fmt"

func main() {
	mapAku := make(map[string]int)
	stringAku := "Pandita Amanu Adi"

	for i := 0; i < len(stringAku); i++ {
		huruf := string(stringAku[i])

		fmt.Println(huruf)

		mapAku[huruf] += 1

	}

	fmt.Print(mapAku)
}
