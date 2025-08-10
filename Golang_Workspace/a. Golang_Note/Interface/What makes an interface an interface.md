What makes an interface an interface is that it's an **abstract type** defined exclusively by a set of **method signatures**.

It is not a class or a struct; it is a **pure specification of behavior**.1

---

### In-Depth Explanation

An interface's identity is defined by three core principles working together:

1. It is a Contract of Behavior2
    
    An interface is nothing more than a named collection of method signatures. It has no data fields and no implementation details. Its entire existence is to declare, "Any type that wants to be considered one of me must have these specific methods." It focuses solely on what a type can do.
    
2. It Is Satisfied Implicitly
    
    This is Go's defining feature for interfaces. A concrete type (like a struct) doesn't need to know the interface exists. If the type possesses all the methods listed in the interface contract—with exactly matching signatures—Go automatically considers it to satisfy that interface.3 This "implicit satisfaction" decouples the implementation from the specification.
    
3. It Enables Polymorphism4
    
    The ultimate purpose of an interface is to allow a single piece of code, like a function, to operate on a variety of different types. By having a function accept an interface type as a parameter, you are creating a flexible API that depends only on the contract, not on a specific concrete type. This allows you to substitute any type that fulfills the contract, making your code abstract, reusable, and decoupled.5