In C, memory management is manual, leading to common errors like **dangling pointers**, **double frees**, and **memory leaks**. Rust, with its ownership model and compile-time checks, prevents these issues in most cases. Here's a comparison:

### **Dangling Pointers**
- **C**: Dangling pointers occur when a pointer references memory that has been freed or gone out of scope. The programmer must manually ensure pointers are valid, which is error-prone.
  ```c
  int* ptr = malloc(sizeof(int));
  free(ptr);
  *ptr = 42; // Undefined behavior: dangling pointer
  ```
- **Rust**: Rust’s ownership and borrowing system ensures that references cannot outlive the data they point to. The compiler catches dangling references at compile time.
  ```rust
  let s = String::from("hello");
  let r = &s; // Reference to s
  drop(s); // s is dropped
  // println!("{}", r); // Compile error: `r` cannot be used after `s` is dropped
  ```
  Rust’s borrow checker enforces that references are valid, eliminating dangling pointers.

### **Double Frees**
- **C**: Double frees happen when `free()` is called multiple times on the same memory block, leading to undefined behavior (e.g., crashes or corruption).
  ```c
  int* ptr = malloc(sizeof(int));
  free(ptr);
  free(ptr); // Undefined behavior: double free
  ```
- **Rust**: Rust’s ownership model ensures that each value has a single owner. When the owner goes out of scope, the memory is automatically freed exactly once via the `Drop` trait. Double frees are impossible in safe Rust.
  ```rust
  let s = String::from("hello"); // s owns the memory
  // Memory is automatically freed when s goes out of scope
  // No way to "free" s twice in safe Rust
  ```
  In unsafe Rust, manual memory management (e.g., via `std::ptr`) could theoretically allow double frees, but this is rare and explicitly marked as unsafe.

### **Memory Leaks**
- **C**: Memory leaks occur when allocated memory is not freed before losing all references to it, causing memory to remain allocated indefinitely.
  ```c
  int* ptr = malloc(sizeof(int));
  ptr = NULL; // Memory is leaked: no way to access or free it
  ```
- **Rust**: Rust’s ownership model ensures that memory is freed when a value’s owner goes out of scope, preventing leaks in most cases. However, leaks can occur in specific scenarios, such as:
  - Reference cycles with `Rc` or `Arc` (e.g., circular references that never get dropped).
  - Forgetting to drop resources explicitly in certain low-level code.
  ```rust
  use std::rc::Rc;
  use std::cell::RefCell;
  struct Node {
      next: Option<Rc<RefCell<Node>>>,
  }
  let a = Rc::new(RefCell::new(Node { next: None }));
  let b = Rc::new(RefCell::new(Node { next: Some(Rc::clone(&a)) }));
  a.borrow_mut().next = Some(Rc::clone(&b)); // Cycle: memory leak
  ```
  Rust’s leaks are rare and usually involve advanced types like `Rc` or manual memory management in unsafe code. Tools like `cargo-valgrind` can help detect leaks.

### **Key Differences**
1. **Safety Guarantees**:
   - C relies on programmer discipline, making it prone to errors.
   - Rust’s compile-time ownership and borrowing rules eliminate most memory errors in safe code, with no runtime overhead.

2. **Performance**:
   - Both languages are high-performance, but C requires manual memory management, which can lead to bugs or inefficiencies if mishandled.
   - Rust automates memory management without a garbage collector, maintaining performance comparable to C while improving safety.

3. **Debugging**:
   - In C, tools like Valgrind or AddressSanitizer are needed to detect memory issues at runtime, which can be slow and may miss issues in untested code paths.
   - Rust catches most issues at compile time, reducing the need for runtime debugging tools. For leaks or unsafe code, tools like `cargo-valgrind` or `miri` can help.

4. **Unsafe Code**:
   - C is inherently unsafe, with no built-in protections.
   - Rust allows unsafe code for low-level operations, but it’s explicitly marked, and safe Rust avoids these issues entirely.

### **Conclusion**
Rust’s ownership model and borrow checker make it vastly safer than C for memory management, preventing dangling pointers, double frees, and most memory leaks at compile time. C gives programmers full control but requires meticulous care to avoid these errors, often relying on external tools for detection. For projects prioritizing safety without sacrificing performance, Rust is a strong alternative to C.