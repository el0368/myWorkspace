Here is the **Intermediate Roadmap**.

At this stage, you know how to open a window and place buttons. The "Intermediate" level is about three things: **Performance**, **Concurrency** (doing things without freezing), and **Customization** (making it not look like a debug tool).

### 🗺️ Phase 1: The Data Handler (Complex Layouts)

**Goal:** Display large amounts of data without the app getting slow.

1. **Tables > Grids:**
    
    - Stop using `egui::Grid` for huge lists. It is slow if you have 1,000 rows.
        
    - **Learn:** `egui_extras::TableBuilder`. It supports sticky headers, resizable columns, and "virtual scrolling" (only drawing what is visible on screen).
        
    - **Task:** Create a table with 10,000 rows of dummy data. Make sure it scrolls smoothly.
        
2. **Virtual Scrolling:**
    
    - Understand the `show_rows` pattern.
        
    - **Concept:** If you have 1,000 items but only 10 fit on screen, `egui` should only execute the code for those 10.
        
    - **Task:** Use `ui.push_id` inside a loop to handle row selection in your big table.
        
3. **Plotting:**
    
    - Learn `egui_plot`. It is the standard way to show graphs.
        
    - **Task:** Make a real-time graph that updates every frame (like a CPU usage monitor).
        

### 🗺️ Phase 2: The Multitasker (Async & Threading)

**Goal:** Click a button to "Download" or "Load File" without the window freezing. This is the **#1 intermediate skill**.

1. **The Channel Pattern:**
    
    - You cannot use `async/await` directly in the UI code easily.
        
    - **Learn:** `std::sync::mpsc` (channels).
        
    - **Pattern:**
        
        1. Spawn a standard thread: `std::thread::spawn`.
            
        2. Do the heavy work in the thread.
            
        3. Send the result back via the channel (`tx.send(data)`).
            
        4. In your `update` loop, check the receiver: `rx.try_recv()`.
            
    - **Task:** Create a button "Calculate Primes". It runs for 5 seconds. The UI must remain responsive (clickable) while it runs.1
        
2. **Repaint Requests:**
    
    - **Concept:** `egui` only redraws when the mouse moves. If your background thread finishes while the user is not moving the mouse, the screen won't update.
        
    - **Fix:** Your background thread must call `ctx.request_repaint()` so the UI updates immediately when the data arrives.2
        

### 🗺️ Phase 3: The Artist (Visuals & Painting)

**Goal:** Break free from the standard widget look.

1. **The `Painter` API:**
    
    - Every `Ui` has a `painter()`. You can draw lines, circles, and rectangles anywhere.
        
    - **Task:** Draw a red circle that follows the mouse cursor.
        
    - **Task:** Draw a custom "Close" icon (an X) using `painter.line_segment`.3
        
2. **The `Sense` API (Custom Interactions):**
    
    - How do you make that red circle clickable?
        
    - **Learn:** `ui.allocate_response(size, Sense::click())`.
        
    - This reserves invisible space and gives you a `Response` object (hovered, clicked, dragged) _before_ you draw the shape on top of it.
        
    - **Task:** Create a "Color Swatch" widget: A colored square that opens a color picker when clicked.4
        
3. **Custom Styling:**
    
    - Don't just use `Visuals::dark()`. Modify the `Style` struct directly.
        
    - **Task:** Change the "selection" color of the app to Neon Pink. Change the corner rounding of all buttons to `0.0` (square).5
        

### 🗺️ Phase 4: Rich Interaction

**Goal:** Professional desktop features.

1. **Drag and Drop:**
    
    - This is native in `egui`.
        
    - **Learn:** `dnd_drag_source` and `dnd_drop_zone`.
        
    - **Task:** Create two lists. Drag items from List A to List B.
        
2. **Modals and Popups:**
    
    - **Learn:** `egui::Window` usage with `set_enabled(false)` on the background.
        
    - **Task:** Create a "Confirm Delete" popup. When open, the buttons behind it should not be clickable.
        
3. **Notifications (Toasts):**
    
    - Use a crate like `egui-toast` or `egui-notify`.
        
    - **Task:** When your background thread finishes "Calculating Primes," show a pop-up toast in the corner saying "Done!".
        

### ✅ The "Intermediate" Final Project

**Project:** **A Stock/Crypto Dashboard**

1. **Left Panel:** A list of 100+ stock symbols (use `TableBuilder` + Virtual Scrolling).
    
2. **Main Area:** A chart (`egui_plot`) showing the price history of the selected stock.
    
3. **Async:** When you click a stock, "fake" a network request (sleep for 1s) in a background thread to load the data. Show a spinner (`ui.spinner()`) while loading.
    
4. **Custom:** Draw a green "Up" arrow or red "Down" arrow next to the price using the `Painter`.
    

Once you can build this, you are ready for **Advanced** topics (Raw Integration, WGPU Shaders, Mobile Gestures).