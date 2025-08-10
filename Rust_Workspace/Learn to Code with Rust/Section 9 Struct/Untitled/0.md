Based on our conversation, you have asked about the following aspects of Rust structs:

### Struct Definition and Types

- What a `struct` is in general (a container for related data).
    
- The three types of structs: named-field, tuple-like, and unit-like.
    
- How to declare a `struct` with the `struct` keyword and `PascalCase` naming.
    
- That a `struct` definition is a blueprint, and an instance is a concrete value.
    

### Struct Instantiation (Creating an Instance)

- How to create an instance by providing values for all fields.
    
- The rule that all fields must be provided with the correct types.
    
- That the order of fields doesn't matter.
    
- The **Field Init Shorthand** syntax (`name,` instead of `name: name,`).
    
- The **Struct Update Syntax** (`..`) to copy fields from an existing struct, and the ownership implications of this.
    

### Struct Fields and Ownership

- How to access a field's value using dot notation (`.`).
    
- The ownership hierarchy: a `struct` owns its fields, and a field owns its value.
    
- How ownership semantics (Copy vs. Move) apply when you access a field.
    
    - `Copy` fields (`i32`, `f64`) are copied, and the original field remains valid.
        
    - Non-`Copy` fields (`String`) are moved, which invalidates the field and partially moves the original `struct`.
        

### Struct Methods and Associated Functions

- What a **method** is (a function attached to an instance).
    
- The `impl` block where methods are defined.
    
- The crucial `self` parameter and its four options: `self`, `mut self`, `&self`, and `&mut self`.
    
- The different ownership and mutability implications of each `self` option.
    
- How to invoke a method (`instance.method()`).
    
- How to call a method from within another method (`self.other_method()`).
    
- What an **associated function** is (a function on the type itself, without `self`).
    
- How to invoke an associated function (`StructName::function()`).
    
- That constructors are a common use case for associated functions (e.g., `String::from`).
    

### Traits and Printing Structs

- That by default, `struct`s don't implement `Display` or `Debug`.
    
- The **`#[derive(Debug)]`** attribute as a shortcut to automatically implement the `Debug` trait for easy printing.
    
- How the `Debug` formatters (`{:?}` and `{:#?}`) work.
    
- That the `Self` alias is a best practice for type annotations within `impl` blocks.
    
- That a single `struct` can have multiple `impl` blocks.