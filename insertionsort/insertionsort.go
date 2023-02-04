package main

import "fmt"

func main() {
	a := []int{1, 12, 1, 23, 44, 567, 77, 1, 0, 1}
	insertionsort(a)
	fmt.Println(a)
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}

}
