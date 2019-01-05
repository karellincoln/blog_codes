package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceSort()
}

type aStructure struct {
	person string
	height int
	weight int
}

// sort.Slice: unstable sort, quicksort implement.
// sort.stableSlice: stable sort, insert sort or merge sort
func sliceSort() {
	mySlice := make([]aStructure, 0)
	a := aStructure{"Mihalis", 180, 90}
	mySlice = append(mySlice, a)
	a = aStructure{"Dimitris", 180, 95}
	mySlice = append(mySlice, a)
	a = aStructure{"Marietta", 155, 45}
	mySlice = append(mySlice, a)
	a = aStructure{"Bill", 134, 40}
	mySlice = append(mySlice, a)

	fmt.Println("0:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight < mySlice[j].weight
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(">:", mySlice)
}