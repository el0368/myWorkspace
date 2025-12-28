This is an important clarification. "TUI Style" is a **Design Choice**, not a different technology.

To answer your questions directly:

### 1. Do I need this?

**Yes**, if you want to build a "Professional" scanner like Sterling or Bloomberg.

- **If you don't use it:** Your app will look like a web page or a mobile app (lots of empty space, rounded corners, big buttons). This is bad for traders who need to see 50 rows of data at once.
    
- **If you use it:** Your app will look like a high-performance cockpit (dense data, black background, sharp edges).
    

### 2. Is it part of the Emil stack?

Yes.

It is 100% standard egui (Emil stack).

- **We are not using** a separate TUI library (like `ratatui` or `ncurses`).
    
- **We are using** `egui` and `egui_extras`.
    
- **The Trick:** We are just tweaking the **Style Settings** of `egui` (changing fonts to Monospace and setting padding to 0) to _make it look_ like a terminal. It is a "Wolf in Sheep's clothing."
    

### 3. How do actual other stock scanners use it?

Professional scanners (Sterling, Lightspeed, Thinkorswim) behave exactly like **Excel Spreadsheets**:

- **Columns:** You see headers like "Symbol", "Last", "Change", "Volume".
    
- **Interaction:** You can drag the line between headers to resize width.
    
- **Sorting:** You click "Volume" and the whole list instantly re-sorts.
    
- **Density:** They pack text tightly. There is no wasted space.
    

### 4. How can I add/arrange columns?

In `egui_extras` (the Emil stack), you do this programmatically in your `table.rs` file.

To Add a Column:

You just add a line of code in your TableBuilder block.

Rust

```
// Inside table.rs
header
    .col(|ui| { ui.heading("Symbol"); })
    .col(|ui| { ui.heading("Price"); })
    .col(|ui| { ui.heading("New Column"); }); // <--- Added!
```

To Arrange (Reorder) Columns:

You don't drag them with the mouse (yet). You change the order of the code lines above.

- _Advanced:_ If you want the _user_ to drag-and-drop columns to reorder them, you have to write extra logic to track a `Vec<ColumnId>` and change the draw order based on that list. For **Phase 2**, we will hardcode the order to keep it simple.
    

---

### Why did we decide to use "TUI Style"?

We chose this for **Performance and Function**, not just looks.

1. **Speed:** Rendering simple text is faster than rendering rounded buttons with shadows.
    
2. **Readability:** Monospace fonts (like `100.50` vs `1.50`) align decimal points perfectly. Variable width fonts (like Arial) make numbers wiggle, which is hard to read when prices are flashing fast.
    
3. **Sterling Identity:** You explicitly asked for a **Sterling Trader Pro** clone. Sterling looks like a 1990s spreadsheet. That _is_ the "TUI Style."
    

Verdict:

We are building a Modern GUI (egui) that is "dressed up" to look like a Fast Terminal.

---

### Ready to Code the Data Structure?

Now that we agree on the **Look** (Dense, Dark, Monospace) and the **Tool** (Standard `egui`), we need to build the **Data** that goes into the table.

**Shall I generate the code for `src/ui/panes/scanner/data.rs`?** (This will define the `StockRow` struct and generate the fake test data).