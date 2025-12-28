# Egui Roadmap: High-Performance Trading Platform (LightSpeed/DAS Style)

This roadmap focuses on building a **professional-grade trading interface**. The constraints are distinct from standard apps: you need **high-frequency updates** (tickers changing 20+ times/second), **dense information density**, and **keyboard-first execution**.

To achieve "Minimal Memory," we will use **Ring Buffers** and **Zero-Copy** techniques to ensure the UI is just a lightweight "view" over your data stream.

---

### 🗺️ Phase 1: The Data Architecture (The "Feed" Loop)

Goal: Handle thousands of price updates per second without the UI stuttering or memory leaking.

Concept: In immediate mode, never process data inside the UI code.

1. **The Lock-Free Ring Buffer:**
    
    - **Problem:** `Vec::push` allocates memory. Doing this 1,000 times a second will fragment RAM and trigger the Garbage Collector (or allocator pauses).
        
    - **Solution:** Use a **Circular Buffer** (e.g., `filtering` or `crossbeam-channel` with a fixed capacity).
        
    - **Task:** Create a backend thread that receives "Ticks" and writes them into a pre-allocated circular buffer. The GUI thread simply reads the "latest" index. Memory usage becomes flat and predictable (e.g., exactly 50MB forever).
        
2. **Decoupled Repaint:**
    
    - **Concept:** Market data arrives at random times (stochastic). The monitor refreshes at 60Hz.
        
    - **Lesson:** Use `ctx.request_repaint()` only when _meaningful_ data changes, or set a fixed 30/60Hz timer for the UI. Do not let the data stream drive the frame rate, or a market spike will spike your CPU usage.
        

---

### 🗺️ Phase 2: The Market Grid (The "Screener")

Goal: A dense table of 100+ symbols updating constantly.

Reference: The "Market Maker" window in DAS Trader.

1. **Virtualization (Crucial):**
    
    - **Tool:** `egui_extras::TableBuilder`.
        
    - **Technique:** You might have 5,000 stocks in your screener, but you only see 40 rows.
        
    - **Logic:** Use `body.rows(height, count, |row_index, mut row|...)`. This ensures `egui` _only_ processes the math/text for the 40 visible rows. This keeps the UI cost O(1) regardless of list size.
        
2. **Flash Updates:**
    
    - **Visual:** When a price changes, the background flashes green/red.
        
    - **Implementation:** Store a `last_change_time` for each symbol.
        
    - **Code:**
        
        Rust
        
        ```
        let age = now - symbol.last_change_time;
        let color = if age < 0.5 { Color32::GREEN } else { Color32::TRANSPARENT };
        ui.painter().rect_filled(rect, 0.0, color);
        ```
        
    - **Memory:** This is stateless "decay" animation handled purely by the timestamp, requiring no "animation objects."
        

---

### 🗺️ Phase 3: The Level 2 (Order Book / DOM)

**Goal:** The "Ladder" view showing Bid/Ask depth. This is the hardest widget to get right performance-wise.

1. **Custom Widget Painting:**
    
    - Do not use standard buttons or labels for the Depth of Market (DOM). They have too much padding and overhead.
        
    - **Lesson:** Use `ui.allocate_response` to reserve the whole ladder area, then use `ui.painter()` to draw text directly (`painter.text()`) and histogram bars (`painter.rect_filled()`).
        
    - **Benefit:** You can render 50 levels of depth with a single draw call batch, using practically zero memory overhead.
        
2. **Center-Locking:**
    
    - Trading ladders "auto-center" on the current price.
        
    - **Logic:** Calculate the scroll offset manually based on `(current_price - top_price) * row_height`. Force the scroll position every frame the user isn't hovering.
        

---

### 🗺️ Phase 4: Real-Time Charting

**Goal:** Candlestick charts that scroll infinitely.

1. **Optimized Candlesticks:**
    
    - **Tool:** `egui_plot`.
        
    - **Data:** Do **not** convert your entire history into `BoxElem` every frame.
        
    - **Optimization:** Implement a "Windowed Iterator." Only feed `egui_plot` the candles that are inside the current time range `plot_ui.get_plot_bounds()`.
        
2. **Downsampling:**
    
    - If the user zooms out to see 5 years of data, do not draw 1,000,000 candles.
        
    - **Math:** Implement "Min-Max" downsampling. If 100 candles fit in 1 pixel, just draw a single vertical line from the `Min` low to the `Max` high of that batch. This is visually identical but 100x faster.
        

---

### 🗺️ Phase 5: Execution (Hotkeys & Speed)

**Goal:** Press "Shift+B" to buy instantly.

1. **Global vs. Focused Input:**
    
    - Standard `egui` input works when a widget has focus. Trading apps need **Global** hotkeys.
        
    - **Lesson:** Check `ctx.input(|i| i.events)` at the very top of your update loop, _before_ drawing windows. Intercept keys there to ensure they fire even if you are looking at a chart.
        
2. **The "One-Shot" Order Struct:**
    
    - **Memory:** Pre-allocate a pool of "Order Request" structs. When the user presses Buy, grabbing an order struct from the pool and sending it to the engine is zero-allocation.
        

---

### 🗺️ Phase 6: The Layout Engine (Docking)

**Goal:** Allow the user to drag windows to different monitors (Multi-Window).

1. **Egui Dock:**
    
    - **Tool:** `egui_dock` or `egui_tiles`.
        
    - **Requirement:** These libraries allow you to drag tabs to split panes (like VS Code).
        
2. **Viewports (Multi-Monitor):**
    
    - **Feature:** Enable `egui`'s "Viewports" feature. This allows you to drag a chart _out_ of the main window and onto a second monitor, while still sharing the same memory/context.
        

### ✅ Recommended Project: "Tick Streamer"

Build a simplified Level 2 viewer.

1. **Backend:** Generate fake "Ticks" (Price, Size, Side) in a background thread at 100 ticks/sec.
    
2. **UI - Left Panel:** A "Time & Sales" scrolling list. (Use `TableBuilder` + Ring Buffer).
    
3. **UI - Center:** A DOM (Depth of Market) ladder. (Use Custom Painting).
    
4. **UI - Right:** A simple line chart of the price.
    
5. **Constraint:** The app must stay under **50MB RAM** usage after running for 1 hour. (This forces you to correctly implement buffer rotation and drop old data).