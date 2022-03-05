package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"Andy", "Danny", "Chelsey", "Alaina"}

	fmt.Println(names, len(names))
	fmt.Println(ages, len(ages))

	// Slices ( Bascially resizable arrays )
	var scores = []int{100, 200, 300}
	scores = append(scores, 85)
	fmt.Println(scores)

	// Slice Ranges ( Grabbs a range of values in the slice )
	rangeOne := names[1:3]  // from one up to 3 but not 3
	rangeTwo := names[:3]   // from start up to 3 but not 3
	rangeThree := names[0:] // from 0 to end
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
