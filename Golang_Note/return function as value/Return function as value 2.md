Yes, that is exactly what "returning a function as a value" means.

---

### Functions as First-Class Values

In Go, functions are considered "first-class values," which means they can be treated just like any other value (like an `int`, `string`, or `struct`). This allows you to:

1. Assign a function to a variable.
    
2. Pass a function as an argument to another function.
    
3. **Return a function as the result (the value) from another function.**
    

The `createGreeter` example perfectly demonstrates this third point. The function being returned _is_ the value.