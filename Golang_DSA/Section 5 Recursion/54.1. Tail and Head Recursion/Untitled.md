Of course. Here is the Go code that demonstrates the difference between converting tail and head recursion into loops.

### Tail Recursion: Easy to Convert

A tail-recursive function's logic maps directly to a simple `for` loop. Both functions below produce the exact same output (`3, 2, 1`).

Go

```Go
package main

import "fmt"

// 1. The Tail-Recursive Version
func tailRecursivePrint(n int) {
	// Base case
	if n <= 0 {
		return
	}
	// Work is done first
	fmt.Println(n)
	// Recursive call is the last action
	tailRecursivePrint(n - 1)
}

// 2. The Equivalent Loop Version
func loopEquivalentForTail(n int) {
	// The base case becomes the loop condition
	for n > 0 {
		// Work is the same
		fmt.Println(n)
		// The recursive call becomes a state change
		n--
	}
}
```

---

### Head Recursion: Difficult to Convert

A head-recursive function cannot be directly converted because it relies on the call stack to reverse the order of operations. To get the same output (`1, 2, 3`), the loop must use a completely different logic.

Go

```Go
package main

import "fmt"

// 1. The Head-Recursive Version
func headRecursivePrint(n int) {
	// Base case
	if n <= 0 {
		return
	}
	// Recursive call happens first
	headRecursivePrint(n - 1)
	// Work is done after the call returns
	fmt.Println(n)
}

// 2. The Loop Version (Requires Different Logic)
// Notice we can't just reorder the previous loop.
// We need a new counter 'i' that counts up.
func loopForHeadOutput(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
```

---

### Running the Examples

You can run this `main` function to see the outputs for yourself.

Go

```Go
package main

import "fmt"

func main() {
	fmt.Println("--- Tail Recursion ---")
	tailRecursivePrint(3)

	fmt.Println("\n--- Loop for Tail ---")
	loopEquivalentForTail(3)

	fmt.Println("\n--- Head Recursion ---")
	headRecursivePrint(3)

	fmt.Println("\n--- Loop for Head ---")
	loopForHeadOutput(3)
}
```