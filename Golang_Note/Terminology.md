Yes, that's the perfect way to put it.

While other programmers will understand what you mean by "constructor," the Go community and official documentation don't use that term because the concept doesn't exist as a special feature in the language.

The idiomatic and more accurate term is a **factory function**.

---

### Why the Difference in Terminology?

- Constructor (in other languages)
    
    A constructor is a special, built-in language feature with unique syntax and rules (like __init__ in Python or ClassName() in Java). It's a core part of the language's object-oriented system.
    
- Factory Function (in Go)
    
    A factory function is just a regular function that, by convention, we use to create and set up a struct. It's not a special piece of syntax; it's simply a design pattern that the Go community has adopted because it's simple and effective.
    

So, when you call a function named `New...` in Go, you're not using a special "constructor," you're just calling a normal function that happens to build a struct for you. Using the term "factory function" shows a clearer understanding of Go's design philosophy.