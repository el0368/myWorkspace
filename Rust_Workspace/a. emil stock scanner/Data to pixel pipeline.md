To add a new column (e.g., "Sector" or "Beta") to your Sterling-style scanner, you must follow a strict **"Data-to-Pixel Pipeline."**

Because you are using a low-level, high-performance architecture, columns do not "magically" appear. You must manually wire them through three specific layers of your micro-modules.

Here is the exact 3-step workflow you will use every time you add a metric:

### Step 1: The Data Layer (`data.rs`)

"Define it."

First, the data needs to exist in memory.

- You open `data.rs` and find the `StockRow` struct.
    
- You add the new field (e.g., `sector` or `vwap`).
    
- _Result:_ The compiler now knows this data exists for every stock, but the screen doesn't show it yet.
    

### Step 2: The Logic Layer (`grid.rs` - Sorting)

"Index it."

If you want to be able to sort by this new column (e.g., click the header to see all "Energy" stocks together):

- You go to the `sort_rows` logic inside `grid.rs`.
    
- You add a new condition to the "match" statement that tells the system how to compare two rows based on this specific new field.
    

### Step 3: The UI Layer (`grid.rs` - Rendering)

"Draw it."

Finally, you tell the TableBuilder to allocate pixel space for it.

- **Header:** You add a new `.column()` definition to the builder. This tells `egui` to reserve space (e.g., "60 pixels" or "Flexible width").
    
- **Body:** You add the drawing logic inside the loop. You tell it specifically: _"In this new cell, read the `sector` field from Step 1 and draw it as white text."_
    

### Summary of the Process

1. **Add Field** (`data.rs`) $\rightarrow$ 2. **Add Sort Rule** (`grid.rs`) $\rightarrow$ 3. **Add Draw Command** (`grid.rs`)
    

This manual process ensures you never accidentally draw slow data. You explicitly control exactly what appears and how much "weight" it adds to your rendering loop.

**Are you ready to create the `data.rs` file now?**