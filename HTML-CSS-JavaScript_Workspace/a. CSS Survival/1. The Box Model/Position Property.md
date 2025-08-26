The `position` property in CSS specifies how an element is placed in a document. It works in conjunction with the offset properties `top`, `right`, `bottom`, and `left`.

---

### `position: static`

This is the **default** value. The element is positioned according to the normal flow of the page. The offset properties (`top`, `right`, etc.) have no effect on it.

---

### `position: relative`

The element is positioned according to the normal flow, but you can then use `top`, `right`, `bottom`, and `left` to move it **relative to its original position**. The space it would have occupied remains preserved in the layout. This is also commonly used to create a positioning context for child elements.

---

### `position: absolute`

The element is **removed from the normal document flow**, and other elements are positioned as if it isn't there. It's then positioned relative to its nearest **positioned ancestor** (an element whose `position` is not `static`). If no such ancestor exists, it's positioned relative to the page itself.

---

### `position: fixed`

The element is removed from the normal flow and positioned relative to the **browser window (viewport)**. It stays in the same place even when the page is scrolled, making it useful for things like persistent navigation bars or "back to top" buttons.

---

### `position: sticky`

This is a hybrid of `relative` and `fixed`. The element scrolls with the page until it reaches a specified offset (e.g., `top: 0`), at which point it **"sticks"** in place like a `fixed` element. This is often used for navigation bars that should scroll into view and then stay at the top.

---

### `z-index`

The `z-index` property controls the vertical stacking order of positioned elements (any element whose position is not `static`). An element with a higher `z-index` will appear on top of an element with a lower `z-index`. You can also use negative values.