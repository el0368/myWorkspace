package main

import "fmt"

func main() {
	n := 3
	fun(n)
}

func fun(n int) {

	if n == 0 {
		return
	}

	fun(n - 1)
	fmt.Println(n)

}