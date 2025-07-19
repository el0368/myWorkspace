Got it, my apologies for the misunderstanding! You just want the LaTeX math expressions, enclosed in dollar signs for inline math, without any Python code or `IPython.display` calls.

I will provide the exact LaTeX syntax for the mathematical expressions from the "Sets and Sorting with Numbers" example.

---

## Sets and Sorting with Numbers (The Numbered Toy Box Problem 🧸🔢) - LaTeX Math Expressions Only

Here are the mathematical expressions in LaTeX, using dollar signs for inline math:

---

### **I. Core Concepts of Sets (Numerical Perspective)**

**A. What is a Set with Numbers?**

- Set of Toy IDs: T={101,205,312,103,450,201}
    

**B. Representing Sets of Numbers**

- Roster Method: T={101,205,312,103,450,201}
    
- Set-Builder Notation: S={id∣id∈T and toy with id has size<10 cm}
    
- Example Result: S={101,103,201}
    

**C. Elements and Membership (Numerical)**

- 101∈T
    
- 500∈/T
    

**D. Types of Sets (Numerical)**

- Finite Set Example: T={101,205,312,103,450,201}
    
- Empty Set: ∅={}
    

E. Set Operations (Numerical Examples)

Let's define two subsets:

- Red Toy IDs: R={id∣id∈T and toy with id is RED}
    
- Example R: R={101,312,201}
    
- Car Toy IDs: C={id∣id∈T and toy with id is a CAR}
    
- Example C: C={205,103}
    

1. Union: R∪C={101,312,201}∪{205,103}={101,103,201,205,312}
    
2. Intersection: R∩C={101,312,201}∩{205,103}=∅
    
3. Difference: T−R={101,205,312,103,450,201}∖{101,312,201}={205,103,450}
    

**F. Subsets and Supersets (Numerical Examples)**

- Let P={101,201} (IDs of plastic toys).
    
- Subset example: {101,201}⊆{101,312,201}
    
- Proper Subset example: {205,103}⊂{101,205,312,103,450,201}
    

---

### **II. Introduction to Sorting (Numerical Perspective)**

**B. Basic Sorting Principles (Numerical)**

- Numerical Order (Ascending) Result: [101,103,201,205,312,450]
    
- Numerical Order (Descending) Result: [450,312,205,201,103,101]
    
- Sorted by Weight (Ascending) Result: [(201,0.3),(101,0.5),(103,0.7),(312,0.8),(205,1.2),(450,2.0)]
    

---

### **III. The Numbered Toy Box Problem: Applying Numerical Concepts**

**A. Scenario:** Tdata​={(101,0.5),(205,1.2),(312,0.8),(103,0.7),(450,2.0),(201,0.3)}

**B. Tasks using Numerical Set Concepts:**

1. Set of "Heavy" Toys: H={id∣(id,weight)∈Tdata​ and weight>1.0}
    
    - Result: H={205,450}
        
2. Set of "Even ID" Toys: E={id∣(id,weight)∈Tdata​ and id is even}
    
    - Result: E={205,312,450}
        
3. Intersection: H∩E={205,450}∩{205,312,450}={205,450}
    
4. Union: H∪E={205,450}∪{205,312,450}={205,312,450}
    

**C. Tasks using Numerical Sorting Concepts:**

1. Sorted by ID (Ascending): [(101,0.5),(103,0.7),(201,0.3),(205,1.2),(312,0.8),(450,2.0)]
    
2. Sorted by Weight (Descending): [(450,2.0),(205,1.2),(312,0.8),(103,0.7),(101,0.5),(201,0.3)]