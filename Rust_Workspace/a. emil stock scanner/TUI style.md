**No, Thinkorswim does NOT use "TUI Style."**

This is the single biggest difference between **Thinkorswim** (Retail) and **Sterling Trader Pro** (Professional/Prop).

Here is the breakdown of why they look different and why we are choosing the "Sterling" path for your scanner.

### 1. The Thinkorswim Look: "Rich GUI"

Thinkorswim is built on Java (Swing). It is designed to look like a high-tech "Cockpit."

- **Visuals:** It uses gradients, rounded tabs, icons, and variable-width fonts (like Arial or Verdana).
    
- **Padding:** It has "medium" density. There is space between rows so it feels comfortable for the average user.
    
- **Philosophy:** "Make complex data look friendly and approachable."
    

### 2. The Sterling Look: "Spreadsheet / TUI Style"

Sterling is built on C++. It is designed to look like a "Database."

- **Visuals:** It is flat. No shadows. No gradients. Just lines and text.
    
- **Fonts:** Often defaults to Monospace or very tight System fonts.
    
- **Philosophy:** "Fit as much data on the screen as physically possible. If it looks ugly, we don't care, as long as it's fast."
    

### Why we are choosing "TUI Style" for _your_ Scanner

You asked to build a **Scanner** first. In the professional world, a Scanner is a tool for **finding targets**, not analyzing them.

If you use the **Thinkorswim Style (Rich GUI)** for a scanner:

- **Pros:** It looks modern and polished.
    
- **Cons:** You can only fit ~25 rows on a screen. Your eye has to travel further to compare numbers because the decimals don't align perfectly.
    

If you use the **Sterling Style (TUI Style)** for a scanner:

- **Pros:** You can fit ~50 rows on the same screen. The decimal points align vertically (thanks to Monospace), so you can spot a price breakout instantly without reading the digits.
    
- **Cons:** It looks "old school" (like the 1990s).
    

### The Verdict for Your Plan

Since you stated your goal is a **"Sterling-like"** platform, we **must** use the TUI/Spreadsheet style for the Scanner.

- **Retail Traders** (Thinkorswim users) want comfort.
    
- **Prop Traders** (Sterling users) want density.
    

**Your Plan:** We will force the Scanner Module to be **Dense, Flat, and Monospace**.

---

### Step 2: How to Add & Arrange Columns (The Logic)

You asked how to manage columns without code. Here is the **Logical Plan** for the Column System.

We will not hardcode columns like "Price" and "Volume" directly into the UI. We will create a **Column Registry**.

**The Concept:**

1. The Registry (The Menu):
    
    Imagine a list of all possible data points we can calculate (e.g., RSI, VWAP, Sector, Beta).
    
2. The Active List (The Config):
    
    We hold a List<ColumnID> that defines what is currently visible.
    
    - _Default:_ `[Symbol, Price, Change, Volume]`
        
3. The Renderer (The Loop):
    
    The UI loop does not say "Draw Price."
    
    It says: For each Column ID in the Active List -> Draw that specific data.
    

**How to Arrange (Reorder):**

- **Logic:** You simply swap the items in the `Active List`.
    
    - _Before:_ `[Symbol, Price]`
        
    - _Action:_ User moves Price first.
        
    - _After:_ `[Price, Symbol]`
        
- **Effect:** The next time the screen draws (16ms later), the columns instantly swap positions because the loop reads the new list order.
    

**How to Add:**

- **Logic:** You pick an ID from the "Registry" and append it to the "Active List".
    

### Summary of the Plan

1. **Style:** We reject the Thinkorswim "Pretty" look. We embrace the Sterling "Dense" look (TUI Style).
    
2. **Columns:** We build a dynamic system where columns are just a list of IDs, allowing instant reordering and adding/removing without rewriting the rendering code.
    

**Are you ready to execute this plan?** (I can proceed to generating the `data.rs` file now).