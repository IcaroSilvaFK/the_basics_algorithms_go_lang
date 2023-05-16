package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d", arr[i])
	}
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func quickSort(
	arr []int,
	low int,
	high int,
) {

	if low < high {
		pivot := partition(arr, low, high)

		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

}

func partition(
	arr []int,
	low int,
	high int,
) int {
	pivot := arr[high]
	sort := 0
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(&arr[i], &arr[j])
		}
		sort++
	}
	fmt.Printf("Sorted: %d\n", sort)
	swap(&arr[i+1], &arr[high])
	return (i + 1)
}
