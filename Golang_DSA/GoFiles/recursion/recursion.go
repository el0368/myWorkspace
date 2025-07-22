package main

import "fmt"

func fun(n int) int {

	if n > 0 {
		return fun(n-1) + n
	}

	return 0

}

func main() {
	n := 5
	fmt.Printf("%d", fun(n))
}
