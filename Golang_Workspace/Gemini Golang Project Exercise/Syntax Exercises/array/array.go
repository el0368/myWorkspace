package main

import "fmt"

func main() {
	var score = [5]int{88, 92, 77, 68, 95}
	fmt.Println("entire array: ", score)
	fmt.Println("index 2: ", score[2])
	score[2] = 81
	
	fmt.Println("update index 2:",score[2])
	
	var sum int
	for i := 0; i < len(score); i++ {
		sum += score[i]
	}
	fmt.Println("sum of score",sum)

	for index, value := range score {
		fmt.Printf("index %d, value %d\n", index, value)
	}
	fmt.Println("original array: ", score)
	modifyArray(score)
	fmt.Println("original array: ", score)
	// you only modified the copied of the array
}

func modifyArray(pastScore [5]int) {
	pastScore[0] = 100
	fmt.Println("modified array:", pastScore)
	
}