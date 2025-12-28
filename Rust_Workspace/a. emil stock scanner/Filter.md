**Yes, you are exactly right.**

Before the scanner can show "Top Gainers," it needs to know what "Top Gainers" means. The **Filter** is the input; the **Grid** is the output.

In a professional "Sterling-style" platform, the filter is not a popup window that blocks your view. It is a permanent **"Control Strip"** fixed at the very top of the scanner window.

Here is the plan for how the Filter looks and behaves in TUI Style.

### 1. The Visual Layout ("The Control Strip")

Imagine a horizontal bar, exactly **28 pixels high**, running across the top of your scanner tab. It is packed with controls. There is no whitespace.

**Left-to-Right Layout:**

1. **The Template Switcher (The "Channel Changer")**
    
    - **Visual:** `[ ▲ | ▼ ]` `[ 1. Top Gainers ▼ ]`
        
    - **Behavior:**
        
        - **Arrows:** Click Up/Down to instantly cycle to the next profile (like changing TV channels). This is the fastest way to scan the market.
            
        - **Dropdown:** Click the name to see the full list of 20+ templates.
            
2. **The Quick Filters (The "Tuner")**
    
    - _These override the current template instantly._
        
    - **Symbol:** `[ Sym.. ]` (Type "A" to see AAPL, AMD).
        
    - **Price:** `[ $ Min ]` `[ $ Max ]`.
        
    - **Volume:** `[ Vol > ]` (e.g., type "500k").
        
3. **The Toggles (The "Switches")**
    
    - **Visual:** Small text boxes that turn Green/Red when clicked.
        
    - **`[H]` (Halted):** Show/Hide halted stocks.
        
    - **`[S]` (Shortable):** Show only stocks you can short.
        
    - **`[O]` (Options):** Show only stocks with options.
        
4. **The Action Buttons (Right Side)**
    
    - **`[ ⚙ ]` (Config):** Opens the deep settings (where you define complex rules like "RSI > 70").
        
    - **`[ ⏸ ]` (Pause):** Freezes the scanner so rows stop jumping while you try to click one.
        

---

### 2. The Logic Flow ("The Pipeline")

You asked: _"Before go to stock scanner I need filter for input right?"_

**Yes.** The data flows through the system in this exact order every 16 milliseconds (60 FPS):

1. **The Source (Raw Data):**
    
    - Input: 8,000 Stocks.
        
    - _System holds everything in memory._
        
2. **The Filter (The Gatekeeper):**
    
    - Input: 8,000 Stocks + Current Template Rules.
        
    - _Question:_ "Is Price > $5.00?" "Is Volume > 1M?"
        
    - _Result:_ 50 Stocks pass the test.
        
3. **The Sorter (The Ranker):**
    
    - Input: 50 Stocks.
        
    - _Action:_ Sort by "% Change" (Highest to Lowest).
        
4. **The Renderer (The TUI):**
    
    - Input: Top 20 sorted stocks.
        
    - _Action:_ Draw them on the screen.
        

---

### 3. How to "Plan" the Filter Component

We will build this in `src/ui/panes/scanner/filters.rs`.

The Data Structure (State):

You need a struct to hold the values of those inputs.

- `min_price: Option<Decimal>`
    
- `max_price: Option<Decimal>`
    
- `min_volume: Option<u64>`
    
- `symbol_query: String`
    

**The "Dirty" Logic:**

- When you type in the `[ $ Min ]` box, the scanner detects a change.
    
- It immediately re-runs the **Pipeline** (Step 2 above).
    
- The Grid updates instantly. You do not need a "Search" button. It should filter _as you type_.
    

### Summary

- **Look:** A dense, single-line horizontal bar at the top.
    
- **Inputs:** Dropdown for templates, Text boxes for quick overrides.
    
- **Behavior:** Instant filtering. No "Submit" button.
    

**Are you ready to create the `data.rs` file now to define the Stock Structure that these filters will act upon?**