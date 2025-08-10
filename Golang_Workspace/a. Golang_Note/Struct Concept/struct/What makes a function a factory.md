What makes a function a factory is its **purpose**: its primary job is to **create, initialize, and return a new object**.

It's not a special type of function in Go; it's a regular function that we use in a specific way.

---

### The Job of a Factory 🏭

A function acts as a factory when it does more than just create a blank struct. It encapsulates the entire setup process:

- **It hides complexity:** The user doesn't need to know the details of how the object is assembled.
    
- **It enforces rules:** It can validate input and set required default values.
    
- **It returns a ready-to-use instance:** The object it gives back is guaranteed to be in a valid state.
    

---

### It's a Pattern, Not a Language Feature

There is no `factory` keyword in Go. We call it a "factory function" because it follows a well-known programming **pattern**. It's a function that we've assigned the specific role of being an object builder.