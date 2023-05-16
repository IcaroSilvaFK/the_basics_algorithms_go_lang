package main

import "fmt"

func main() {
	bubbleSort()
}

func bubbleSort() {

	arr := []int32{213, 321, 23, 3123, 2323, 3213, 213123, 213, 213, 12, 3, 312, 313, 213, 31}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = swap
			}
		}
	}

	fmt.Println(arr)

}
