Of course. Here is the previous answer with raw LaTeX formulas.

---

### ## Defining Sets and Elements

First, we define our sets with formal notation. Let the **universal set**, U, be all the toys in the box.

Let's define two subsets based on properties:

Let R be the set of all red toys.

Let C be the set of all cars.

If our toy box contains a red car, a red block, a blue car, and a yellow truck, we can write these sets by listing their elements inside curly braces:

R={red car, red block}

C={red car, blue car}

The symbol ∈ means "is an element of." So, we can state that red car ∈R and red car ∈C.

---

### ## Subsets and Cardinality

A set A is a subset of a set B if all elements of A are also in B. This is written as A⊆B.

In our example, both R and C are subsets of the universal set U, so R⊆U and C⊆U.

The cardinality of a set is the number of elements it contains, denoted by vertical bars ∣S∣.

∣R∣=2

∣C∣=2

---

### ## Set Operations

This is where we formalize the relationships between sets.

#### ## Intersection (∩)

The intersection of two sets contains only the elements that are in both sets.

R∩C={elements in R AND in C}

R∩C={red car}

The cardinality is ∣R∩C∣=1.

#### ## Union (∪)

The union of two sets contains all elements that are in either set, or in both. Duplicates are only listed once.

R∪C={elements in R OR in C}

R∪C={red car, red block, blue car}

The cardinality is ∣R∪C∣=3.

#### ## Principle of Inclusion-Exclusion

This principle provides a formula for the cardinality of a union:

∣R∪C∣=∣R∣+∣C∣−∣R∩C∣

Using our values: 3=2+2−1, which is true. This formula corrects for the double-counting of elements in the intersection.

#### ## Complement (Ac)

The complement of a set contains everything in the universal set U that is not in that set.

Rc={all toys in U that are NOT red}

Rc={blue car, yellow truck}