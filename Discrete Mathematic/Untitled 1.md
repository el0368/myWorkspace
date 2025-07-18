---
summary: This document explains fundamental set theory concepts, including universal sets, elements, subsets, cardinality, and operations like intersection, union, and complement, using a toy collection as an illustrative example.
---
Of course, here is the text you provided, reformatted with proper mathematical notation.

---

### ## Defining Sets and Elements

First, we define our **universal set**, U, which contains all the toys in our collection.

Let's say our toy box contains a red car, a blue truck, a red block, and a blue car. We can write this as:

U={red car, blue truck, red block, blue car}

Now, let's define two subsets based on properties:

- Let R be the set of all red toys.
    
- Let C be the set of all cars.
    

Listing the elements of these subsets, we get:

R={red car, red block}

C={red car, blue car}

The symbol ∈ means "is an element of." We can see that the red car is an element of both sets:

- `red car` ∈R
    
- `red car` ∈C
    

The `red block` is an element of R but not C:

- `red block` ∈R
    
- `red block` ∈/C
    

---

### ## Subsets and Cardinality

A set A is a **subset** of a set B if all elements of A are also found in B. This is written as A⊆B. In our example, both R and C are subsets of the universal set U:

- R⊆U
    
- C⊆U
    

The **cardinality** of a set is simply the number of elements it contains, denoted by ∣S∣.

- The cardinality of the set of red toys is ∣R∣=2.
    
- The cardinality of the set of cars is ∣C∣=2.
    

---

### ## Set Operations

This is where we formalize the relationships between our sets.

#### ## Intersection (∩)

The intersection of two sets contains only the elements that exist in both sets. The intersection of R and C is:

R∩C={red car}

The cardinality of the intersection is ∣R∩C∣=1.

å#### ## Union (∪)

The union of two sets contains all elements that are in either set, listing each unique element only once. The union of R and C is:

R∪C={red car, red block, blue car}

The cardinality of the union is ∣R∪C∣=3.

This illustrates the Principle of Inclusion-Exclusion:

∣R∪C∣=∣R∣+∣C∣−∣R∩C∣

3=2+2−1

#### ## Complement (Ac)

The complement of a set contains everything in the universal set U that is not in that set. The complement of R is:

Rc={blue truck, blue car}

convert information above to LaTex pls ?




