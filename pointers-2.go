package main

import (
	"fmt"
	"sort"
)

type lessFn func(i, j int) bool

func sortStrings(s []string, less lessFn) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}

	sort.Slice(s, less)
}

func main() {
	food := []string{"apple", "banana", "egg", "cookie"}
	sortStrings(food, nil)
	fmt.Println(food)

	var ingredients []string
	withIngredients := append(ingredients, "salt", "pepper", "vegetable", "meat")
	fmt.Println(withIngredients)
}