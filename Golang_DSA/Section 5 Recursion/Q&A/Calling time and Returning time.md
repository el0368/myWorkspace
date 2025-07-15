**Calling time** is the phase of recursion where a function calls itself, moving deeper toward the base case. **Returning time** is the phase after the base case has been hit, when the functions begin to finish and return control back up the chain.

---

## Calling Time (Ascending Phase) 🔽

This is the period when the recursive function is making nested calls. Any code that appears **before** the recursive call is executed during this phase.

Go

```Go
func example(n int) {
    if n > 0 {
        // This work happens during CALLING TIME.
        fmt.Println("Going down:", n)
        
        example(n - 1) // The recursive call

        // ...
    }
}
```

---

### Returning Time (Descending Phase) 🔼

This is the period when the chain of function calls starts to unwind after reaching the base case. Any code that appears **after** the recursive call is executed during this phase.

Go

```Go
func example(n int) {
    if n > 0 {
        // ...

        example(n - 1) // The recursive call

        // This work happens during RETURNING TIME.
        fmt.Println("Coming back up:", n)
    }
}
```