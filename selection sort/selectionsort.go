package main

import "fmt"

func Selection_Sort(array []int, size int) []int {
	var min_index int
	for i := 0; i < size-1; i++ {
		min_index = i
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		array[i], array[min_index] = array[min_index], array[i]
	}
	return array
}
func main() {
	var num = 7
	array := []int{2, 4, 3, 1, 6, 8, 5}
	fmt.Println(Selection_Sort(array, num))
}
