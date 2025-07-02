Could you please generate bubble sort in golang?

 Certainly! Below is an implementation of the Bubble Sort algorithm in Go (Golang):

```go
package main

import (
	"fmt"
)

// BubbleSort sorts an array using the bubble sort algorithm.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Track whether any swapping was done
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no two elements were swapped by inner loop, then break
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)

	BubbleSort(arr)

	fmt.Println("Sorted array:  ", arr)
}
```

### Explanation:

- **Function `BubbleSort`:** This function takes a slice of integers as input and sorts it using the bubble sort algorithm.
  
- **Outer Loop (`for i`):** This loop iterates over the entire list. After each iteration, the largest unsorted element bubbles up to its correct position.

- **Inner Loop (`for j`):** This loop compares adjacent elements and swaps them if they are in the wrong order. The loop runs fewer iterations each time since the largest elements settle at the end of the list.

- **Optimized with Swapping Check:** The `swapped` variable tracks if any swaps were made during an iteration. If no swaps were made, it means the list is already sorted and we can break out of the loop early.

### Running The Code:
To run this code:

1. Ensure you have Go installed on your machine.
2. Save this code in a file with a `.go` extension (e.g., `bubblesort.go`).
3. Open your terminal or command prompt.
4. Navigate to the directory where your file is located.
5. Run the program using `go run bubblesort.go`.

You'll see output showing both the unsorted and sorted arrays.