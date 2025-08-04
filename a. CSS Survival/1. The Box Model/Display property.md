The `display` property in CSS controls an element's layout behavior. While every HTML element has a default display type, you can change it to control its placement and sizing.

---

### Hiding Elements: `none` vs. `hidden`

There are two common ways to hide an element, and they behave differently:

- **`display: none;`**: This completely removes the element from the document. It takes up no space, and the layout adjusts as if the element never existed.
    
- **`visibility: hidden;`**: This makes the element invisible but leaves an empty space where it used to be. The layout does not change.
    

In most cases, `display: none;` is the more commonly used property for hiding elements.

---

### `block` vs. `inline`

You can switch an element's default display type to change how it interacts with other elements.

- **`display: block;`**: Makes an element start on a new line and take up the full width available. You can apply this to an inline element, like an `<a>` tag, to make it occupy its own line.
    
- **`display: inline;`**: Makes an element sit next to its neighbors on the same line, taking up only as much width as its content needs. You can apply this to block elements, like `<p>` tags, to make them flow together.
    

---

### `inline-block`

This value combines the features of both `inline` and `block`. An element with `display: inline-block;` will sit on the same line as other elements but also allows you to set a specific **`width`**, **`height`**, and **vertical `margin`**, which is not possible on a standard `inline` element.

This is very useful for styling elements like buttons or navigation menu items that need to be side-by-side but also require specific dimensions and spacing.