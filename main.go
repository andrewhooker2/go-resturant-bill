package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hi new Go Dev!"

	// fmt.Println(strings.Contains(greeting, "Hi"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hi", "Yo!"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ev"))
	// fmt.Println(strings.Split(greeting, " "))
	// fmt.Println(greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	names := []string{"Chelsey", "Andy", "Danny", "Alaina", "Honey"}
	fmt.Println("Unsorted Ages: ", ages)
	fmt.Println("Unsorted Names: ", names)

	sort.Ints(ages)
	fmt.Println("Sorted ages slice: ", ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println("Searching sorted ages for 30: ", index)

	sort.Strings(names)
	fmt.Println("Sorted names slice: ", names)
	nameIndex := sort.SearchStrings(names, "Andy")
	fmt.Println("Searching for Andy: ", nameIndex)

}
