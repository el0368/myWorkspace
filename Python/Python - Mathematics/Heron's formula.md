Sure, Heron's formula allows you to calculate the area of a triangle if you know the lengths of all three sides. Here it is in LaTeX format:

\[ A = \sqrt{s(s - a)(s - b)(s - c)} \]

Where:
- \( A \) is the area of the triangle,
- \( a, b, c \) are the lengths of the sides of the triangle,
- \( s \) is the semi-perimeter of the triangle, given by \( s = \frac{a + b + c}{2} \).

Here's how you could implement Heron's formula in Python:

```python
import math

def herons_formula(a, b, c):
    # Compute semi-perimeter
    s = (a + b + c) / 2

    # Compute area using Heron's formula
    area = math.sqrt(s * (s - a) * (s - b) * (s - c))

    return area
```

You can use this function like so:

```python
print(herons_formula(3, 4, 5))  # Output: 6.0
```

In this example, the function computes the area of a triangle with side lengths 3, 4, and 5 to be 6 square units.