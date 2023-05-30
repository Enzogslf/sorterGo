package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ngl i copied this off of youtube, still trying to understand the code the guy in the tutorial didnt rly explain anything
func main() {
	//lenght of array to be sorted
	piece := createpiece(15)
	fmt.Println("\n Unsorted: \n", piece)
	quicksort(piece)
	fmt.Println("\n sorted \n", piece)
}

// makes array with numbers from 999 to -999
func createpiece(size int) []int {
	piece := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		piece[i] = rand.Intn(999) - rand.Intn(999)
	}
	return piece
}
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	center := rand.Int() % len(a)
	a[center], a[right] = a[right], a[center]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}
