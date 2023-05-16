package main

import (
	"fmt"
	"time"
)

func main() {
	bubbleSort()
}

func bubbleSort() {

	init := time.Now()

	arr := []int32{9, 8, 7, 6, 5, 4, 3, 2, 1}
	count := 0

	for i := 0; i < len(arr)-1; i++ { //O(1)
		for j := 0; j < len(arr)-1; j++ { //O(1)
			if arr[j] > arr[j+1] {
				swap := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = swap
			}
			count++
		}
		count++
	}

	fmt.Println(time.Since(init))
	fmt.Println(count)
	fmt.Println(arr)
	// O(n^2)
}
