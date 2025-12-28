Here is your complete architectural blueprint for the **Stock Scanner Phase**.

This plan focuses on building a **High-Frequency, Institutional-Grade Scanner** (Sterling/LightSpeed style) that runs inside your existing Rust application.

---

### 1. The Dependencies (`Cargo.toml`)

To build a professional scanner, you need specific crates to handle table virtualization, financial math, and time formatting. You cannot rely on standard `f64` floats for money.

**Add these to your `[dependencies]` section:**

1. **`egui_extras`**
    
    - **Purpose:** Provides the **`TableBuilder`**. This is the engine of your scanner. It handles "Virtualization" (rendering only the 20 rows visible on screen while scrolling through 10,000 stocks) and "Sticky Headers" (keeping column names visible).
        
2. **`rust_decimal`** (with `serde` feature)
    
    - **Purpose:** **Precision Math.** Never use `f64` for stock prices (it creates errors like `$150.00000004`). `Decimal` ensures exact penny accuracy for sorting and filtering.
        
3. **`chrono`** (with `serde` feature)
    
    - **Purpose:** **Time.** Used to handle the "Last Updated" timestamp for every tick. Essential for calculating if a quote is "stale" or fresh.
        
4. **`thousands`**
    
    - **Purpose:** **Readability.** Automatically formats volume `1540291` as `"1,540,291"`. Traders cannot read unformatted numbers quickly.
        
5. **`rand`** (Dev Dependency only)
    
    - **Purpose:** **Simulation.** Used in your `scanner_debug` harness to generate fake "market noise" so you can test if the UI flashes correctly when prices change.
        

---

### 2. The UX/UI Blueprint ("The Command Station")

We are skipping "Modern/Soft" design. We are building a **Data Terminal**.

#### **A. The Visual Language (TUI Style)**

- **Font:** Strictly **Monospace** (e.g., Hack, JetBrains Mono, or system default monospace). This ensures `100.00` and `999.99` align perfectly vertically.
    
- **Padding:** **Zero.** Standard UI has 5-10px padding. We want 0px. Data should touch the cell borders.
    
- **Colors:**
    
    - **Background:** Pure Black (`#000000`).
        
    - **Text:** White or Light Grey (`#E0E0E0`).
        
    - **Uptick:** Neon Green (`#00FF00`).
        
    - **Downtick:** Pure Red (`#FF0000`).
        
    - **Selection:** Deep Grey highlight (No round corners).
        

#### **B. The Layout (Vertical Stripes)**

Using `StripBuilder`, we divide the tab into three rigid zones:

- **Zone 1: The Control Deck (Top - Fixed Height)**
    
    - **Symbol Input:** A focused text box to filter the list instantly (e.g., typing "T" shows TSLA, T, TM).
        
    - **Price Filter:** "Min" and "Max" inputs to hide penny stocks.
        
    - **Universe Toggle:** A dropdown to switch lists (e.g., "All Stocks", "Gap Up", "Halted").
        
- **Zone 2: The Data Grid (Middle - Flexible Height)**
    
    - This takes up 90% of the screen.
        
    - **Scrollbar:** Always visible on the right.
        
    - **Headers:** Clickable headers to sort by that column (Ascending/Descending).
        
- **Zone 3: The Status Bar (Bottom - Fixed Height)**
    
    - **Left:** "Total Scanned: 8,412"
        
    - **Right:** "Latency: 12ms" (Heartbeat indicator).
        

#### **C. The Columns (Left to Right)**

This is the standard Institutional layout:

1. **Symbol:** Bold text. Left-aligned.
    
2. **Last Price:** The most active column. Flashes Green/Red on update.
    
3. **% Change:** The primary sorting metric for day traders.
    
4. **Net Change:** The dollar amount move.
    
5. **Volume:** Total volume today (formatted with commas).
    
6. **Imbalance:** (Advanced) Shows Buy vs. Sell MOC imbalance.
    
7. **Time:** The HH:MM:SS of the last trade.
    

---

### 3. The Interaction Model

How the user "feels" the software.

- **The "Flash" Effect:**
    
    - When a price updates, the _background_ of that specific cell flashes Green or Red for 300ms, then fades back to Black. This draws the eye to movement.
        
- **Keyboard Navigation:**
    
    - **Up/Down Arrows:** Move the selection highlight without touching the mouse.
        
    - **Enter:** Opens the selected symbol in the Linked Chart.
        
- **Sorting:**
    
    - Clicking "% Change" instantly reorders the 10,000 rows. The "Virtual Window" must handle this calculation in under 1ms.
        

---

### 4. Implementation Plan

This is the sequence of coding tasks for your `scanner_debug.rs` harness.

**Step 1: The Data Skeleton**

- Create the `StockRow` struct with `Decimal` fields.
    
- Generate a `Vec<StockRow>` with 5,000 dummy items.
    

**Step 2: The Grid Engine**

- Implement `TableBuilder` inside the UI loop.
    
- Map the `StockRow` fields to table columns.
    
- Prove that scrolling is smooth at 60 FPS.
    

**Step 3: The "Flash" Logic**

- Add a `last_update_time` and `prev_price` to the state.
    
- Write logic: _If `now - last_update < 0.5s`, draw background color._
    

**Step 4: The Controls**

- Connect the "Search" input to a `.filter()` function.
    
- Only pass the _filtered_ list to the `TableBuilder`.
    

**This plan is ready for execution.** You have the file structure, the dependencies, and the design spec. We can begin coding the Data Skeleton whenever you are ready.