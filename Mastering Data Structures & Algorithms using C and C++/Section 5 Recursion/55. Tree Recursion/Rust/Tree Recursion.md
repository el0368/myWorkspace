Here's a study guide on Tree Recursion, adapted for Rust, based on the provided transcript. It focuses on the concepts, memory implications, and complexity.

---

## Tree Recursion in Rust: Concepts, Tracing, and Complexity

### 1. Introduction: Linear vs. Tree Recursion

Recursion involves a function calling itself.1 The way it calls itself determines if it's "linear" or "tree" recursion.

- **Linear Recursion:**
    
    - A recursive function is said to be linear if it calls itself **only one time** within its body.2
        
    - **Examples:** Factorial calculation, Fibonacci (if implemented iteratively or with memoization to avoid re-calculation), simple array traversals.
        
    - **Previous lessons:** We've seen examples like head recursion (recursive call is the last operation) and tail recursion (recursive call is the last operation and its result is directly returned). Even if there's processing before/after the call, if there's only one recursive call, it's linear.
        
- **Tree Recursion:**
    
    - A recursive function is said to be tree recursion if it calls itself **more than one time** within its body.3
        
    - **Structure:** This leads to a "tree-like" execution flow, where each function call branches into multiple sub-calls.
        
    - **Examples:** The classic Fibonacci sequence (without memoization), traversing actual tree data structures (binary trees, N-ary trees), certain search algorithms (e.g., brute-force pathfinding).
        
    - **Impact:** Tree recursion generally involves higher time complexity but sometimes lower space complexity compared to linear recursion (for problems where the intermediate results can be discarded).
        

### 2. Example of Tree Recursion in Rust

Let's use a Rust equivalent of the pseudo-code example from the transcript to illustrate.

Rust

```Rust
// This function prints 'n' and then recursively calls itself twice.
// This is a simple example to demonstrate tree recursion's call pattern.
fn fun(n: i32) { // Note: Returns '()' (Unit type) because it only prints, doesn't return value
    if n > 0 {
        println!("{}", n); // Step 1: Print current 'n'
        fun(n - 1);         // Step 2: First recursive call
        fun(n - 1);         // Step 3: Second recursive call (This is what makes it tree recursion)
    }
}

fn main() {
    println!("Starting Tree Recursion (fun(3)):");
    fun(3);
    println!("\nTree Recursion finished.");
}
```

### 3. Tracing Tree Recursion (Execution Flow and Output)

Tracing a tree recursion visually helps understand the order of operations and calls.

**Tracing `fun(3)`:**

```Rust
fun(3)
├── println!(3)  -> Output: 3
├── fun(2)
│   ├── println!(2) -> Output: 2
│   ├── fun(1)
│   │   ├── println!(1) -> Output: 1
│   │   ├── fun(0)      -> Returns (terminates)
│   │   └── fun(0)      -> Returns (terminates)
│   └── fun(1)         // This call happens AFTER the first fun(1) branch fully completes
│       ├── println!(1) -> Output: 1
│       ├── fun(0)      -> Returns (terminates)
│       └── fun(0)      -> Returns (terminates)
└── fun(2)             // This call happens AFTER the first fun(2) branch fully completes
    ├── println!(2) -> Output: 2
    ├── fun(1)
    │   ├── println!(1) -> Output: 1
    │   ├── fun(0)      -> Returns (terminates)
    │   └── fun(0)      -> Returns (terminates)
    └── fun(1)
        ├── println!(1) -> Output: 1
        ├── fun(0)      -> Returns (terminates)
        └── fun(0)      -> Returns (terminates)
```

**Output in order (as observed by `println!`):**

```Rust
3
2
1
1
2
1
1
```

**Key Observation:** The output `3, 2, 1, 1, 2, 1, 1` shows a depth-first traversal of the implicit call tree. The `println!` for a node is executed, then its _first_ child branch completes entirely, then its _second_ child branch completes entirely.

### 4. Stack Usage in Tree Recursion (Rust's Call Stack)

In Rust, as in C, each function call (including recursive calls) creates a new **stack frame** on the program's **call stack**.4

- **Stack Frame (Activation Record):** A block of memory allocated on the stack for each active function call. It contains:
    
    - Local variables for that specific call (e.g., the `n` parameter for `fun(3)` vs `fun(2)`).
        
    - Return address (where to go back to after the function finishes).
        
    - Other function-specific data.
        
- **Creation & Deletion:**
    
    1. When `fun(n)` is called, a new stack frame for `fun(n)` is pushed onto the stack.
        
    2. When `fun(n)` calls `fun(n-1)` (first call), a new stack frame for `fun(n-1)` is pushed.
        
    3. This continues until a base case is reached (e.g., `fun(0)`).
        
    4. When a call reaches its base case or completes all its operations, its stack frame is **popped (deleted)** from the stack, and control returns to the previous frame.
        
    5. Crucially, for tree recursion, when the first branch of `fun(n)` finishes and its stack frames are popped, `fun(n)` then makes its _second_ call to `fun(n-1)`. This _new_ `fun(n-1)` call gets its _own_ set of stack frames, potentially reusing the memory just freed by the first branch.
        

**Stack State for `fun(3)` (Simplified Walkthrough):**

1. `fun(3)` pushed
    
2. `fun(2)` pushed (by `fun(3)`)
    
3. `fun(1)` pushed (by `fun(2)`)
    
4. `fun(0)` pushed (by `fun(1)`)
    
    - **Max Stack Height: 4 frames** (`main` -> `fun(3)` -> `fun(2)` -> `fun(1)` -> `fun(0)`)
        
5. `fun(0)` pops
    
6. `fun(1)` makes second `fun(0)` call, `fun(0)` pushed
    
7. `fun(0)` pops
    
8. `fun(1)` pops
    
9. `fun(2)` makes second `fun(1)` call, `fun(1)` pushed
    
10. `fun(0)` pushed (by `fun(1)`)
    
11. `fun(0)` pops
    
12. `fun(1)` makes second `fun(0)` call, `fun(0)` pushed
    
13. `fun(0)` pops
    
14. `fun(1)` pops
    
15. `fun(2)` pops
    
16. `fun(3)` makes second `fun(2)` call, `fun(2)` pushed
    
    - ...and so on...
        

**Key takeaway on Stack:** The **maximum height** of the stack is generally what determines space complexity, not the total number of calls, because memory is reused.

### 5. Time Complexity Analysis (O(2^n) for this example)

Time complexity is a measure of how the running time of an algorithm grows with the input size (`n`).5 For recursive functions, it often relates to the number of function calls.

- **Counting Calls (Level by Level):**
    
    - Level 0 (initial call `fun(n)`): 1 call (20)
        
    - Level 1 (`fun(n-1)` from initial call): 2 calls (21)
        
    - Level 2 (`fun(n-2)` calls): 4 calls (22)
        
    - Level `k` (`fun(n-k)` calls): 2k calls
        
    - This pattern continues until the base case `fun(0)` is reached. For input `n`, there are `n+1` levels (from `n` down to `0`).
        
- Total Calls for fun(n): The total number of calls is the sum of calls at each level:
    
    20+21+22+dots+2n
    
    This is a geometric series sum, which has a formula: 2(n+1)−1.
    
- Big O Notation: In Big O notation, we simplify to the dominant term, ignoring constants and lower-order terms.6
    
    O(2(n+1)−1)=O(2ncdot21−1)=O(2n)
    
- **Conclusion for `fun(n)`:** The time complexity for this specific `fun(n)` example is **O(2n)**. This is exponential complexity and grows very rapidly, making it inefficient for large `n`.
    
- **Important Note:** Not _every_ tree recursion will have O(2n) time complexity. It depends on how many times the function calls itself and how much work is done per call. The Fibonacci sequence (naive recursive implementation) is a classic O(2n) example of tree recursion.
    

### 6. Space Complexity Analysis (O(n) for this example)

Space complexity measures how much memory an algorithm uses, often focusing on the auxiliary space (beyond input storage).7 For recursion, this is primarily determined by the **maximum depth of the call stack**.

- **Maximum Stack Height:**
    
    - For `fun(3)`, the deepest call sequence goes: `fun(3)` -> `fun(2)` -> `fun(1)` -> `fun(0)`.
        
    - This creates 4 stack frames simultaneously before returning.
        
    - For `fun(n)`, the maximum stack height will be `n + 1` (including the initial call).
        
- **Memory Reuse:** Even though many total calls are made (15 for `n=3`), the stack memory for a completed branch is immediately deallocated and can be reused by a subsequent branch. You don't need all 15 activation records on the stack at the _same time_.
    
- Big O Notation:
    
    O(n+1)=O(n)
    
- **Conclusion for `fun(n)`:** The space complexity for this specific `fun(n)` example is **O(n)**. This is linear space complexity, which is generally more manageable than exponential time complexity.
    

### 7. Summary of Tree Recursion in Rust

- **Definition:** A recursive function that makes **more than one recursive call** within its body.
    
- **Execution:** Results in a branching, tree-like sequence of function calls.
    
- **Stack Usage:** Each call creates a stack frame. Max stack height determines space complexity. Memory is reused for different branches.
    
- **Time Complexity:** Can be very high (e.g., O(2n) for this example), depending on the number of branches per call.
    
- **Space Complexity:** Often linear, O(n), corresponding to the maximum depth of the recursion.
    

While tree recursion can sometimes elegantly model problems, its high time complexity (especially exponential) often necessitates optimization techniques like memoization (caching results) or dynamic programming to be practical for larger inputs.