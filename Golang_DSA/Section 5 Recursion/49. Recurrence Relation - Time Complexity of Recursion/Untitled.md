Of course. Based on the provided materials, here is a detailed explanation of the two methods the instructor uses to find the time complexity of the recursive function.

The core principle for this analysis is the assumption that **every simple statement takes one unit of time** for execution.

### Method 1: The Intuitive Approach (Analysis from the Tracing Tree)

This method involves looking at the function's main work and counting how many times it's performed.

1. **Identify the Primary Work:** The main operation inside the `fun1(n)` function is the `printf` statement, which prints the value of `n`.
    
2. **Count the Operations:** By tracing the function with an input like `n=3`, we see that it prints "3", then "2", then "1". It performs the main work 3 times.
    
3. **Generalize for `n`:** The instructor extends this with an analogy: if stacking one book takes one unit of time, stacking 'n' books takes 'n' units of time. Similarly, if the function is called with `n`, it will perform the print operation `n` times.
    
4. **Determine the Complexity:** Since the total time is directly proportional to the input `n`, the time complexity is **`O(n)`**.
    

---

### Method 2: The Formal Approach (Using a Recurrence Relation)

This method uses a more rigorous, mathematical approach to arrive at the same conclusion.

#### 1. Formulate the Recurrence Relation

First, we define **`T(n)`** as the total time taken by the function for an input `n`. We then express this time in terms of the work done inside the function.

- The `if(n>0)` check takes **1 unit** of time.
    
- The `printf` statement takes **1 unit** of time.
    
- The recursive call `fun1(n-1)` is not a simple statement. Its time cost is, by definition, **`T(n-1)`**.
    

The total time $T(n)$ is the sum of these parts: `T(n) = T(n-1) + 1 + 1`, which simplifies to `T(n) = T(n-1) + 2`. For complexity analysis, all constant work can be represented by a single unit `1`, giving the final relation:

$T(n)=T(n−1)+1$

The **base case** is when `n=0`. The function only performs the `if` check and exits, so **`T(0) = 1`**.

#### 2. Solve the Relation using "Successive Substitution"

Next, we solve this relation by repeatedly substituting the formula into itself to find a general pattern.

- **Start:** $T(n)=T(n−1)+1$
    
- Substitute for T(n-1): We know T(n-1) = T(n-2) + 1. Plugging this in gives:
    
    $T(n)=(T(n−2)+1)+1=T(n−2)+2$
    
- Substitute again: We know T(n-2) = T(n-3) + 1. Plugging this in gives:
    
    $T(n)=(T(n−3)+1)+2=T(n−3)+3$
    
- Generalize the Pattern: After k steps of substitution, a clear pattern emerges:
    
    $T(n)=T(n−k)+k$
    

#### 3. Apply the Base Case

We continue this process until `n-k` equals our base case, `0`.

- Assume `n - k = 0`, which means `k = n`.
    
- Substitute k=n into the generalized formula:
    
    $T(n)=T(n−n)+n$
    
    $T(n)=T(0)+n$
    
- Since we know $T(0) = 1$, the final solution is **$T(n) = n + 1$**.
    

#### 4. Final Result

The solved equation `T(n) = n + 1` is a linear polynomial. The highest degree term is `n`, so the time complexity is **`O(n)`**, which confirms the result from the first, more intuitive method.