The one thing that makes a function a **method** is the **receiver**.

---

### What is a Receiver?

The receiver is a special parameter that comes before the function name.1 It binds the function to a specific type, making it a method of that type.

Because it's bound to a type, you call the method _on an instance_ of that type (e.g., `myObject.doSomething()`).2

### Function vs. Method

- Function: A standalone piece of code.3
    
```Go
func doSomething()
```
    
- Method: A function that belongs to a type, linked by a receiver.
```Go
func (o myObject) doSomething()
```
    
    

In this example, `(o myObject)` is the receiver. It attaches `doSomething` to the `myObject` type, turning it from a regular function into a method.