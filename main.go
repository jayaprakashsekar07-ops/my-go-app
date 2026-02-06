package main

import "fmt"

func main() {
	fmt.Println("Hello! My Git journey starts now 🚀")

	// fmt.Println("Hellooooooooo")
	store := make(map[int]int)
	var numbers = []int{10, 20, 30, 40, 50, 60, 70}
	for index, val := range numbers {
		fmt.Println(store[index] == val)
	}
}
