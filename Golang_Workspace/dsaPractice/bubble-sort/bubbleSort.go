package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	nums := len(arr)

	for i := 0; i < nums-1; i++ {
		for j := 0; j < nums-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}