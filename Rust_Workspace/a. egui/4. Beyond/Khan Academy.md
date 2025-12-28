Here is the specialized roadmap for building a **High-Performance Educational App** (like Khan Academy).

This roadmap is strict about **Memory Usage** and focuses on **Mathematical/Scientific** features. We will move from "Drawing a graph" to "Simulating physics on the GPU."

### 🗺️ Phase 1: The "Zero-Copy" Foundation

Goal: Build an app that uses almost no RAM when idle.

Why: Educational apps often have many pages. If you keep every page in memory, the app crashes.

1. **State Architecture (The "Hot Potato" Pattern):**
    
    - **Bad Way:** Storing a `Vec<f64>` of 1,000,000 items in your GUI struct.
        
    - **Good Way:** Store a _Formula_ or a _Generator_. Calculate values only when `egui` asks for them to draw.
        
    - **Lesson:** Learn to use `Iterator`s in Rust. Pass iterators to `egui`, not Vectors.
        
    - **Memory Trick:** Use `Arc<Vec<T>>` for read-only data. If 5 widgets need to see the same data, they all share one pointer. Do not clone data.
        
2. **Asset Management (The "Library" Pattern):**
    
    - Khan Academy has many images/diagrams.
        
    - **Lesson:** Learn `egui::TextureHandle`.
        
    - **Crucial:** Learn `ctx.forget_image(uri)`. When a user leaves a "Lesson," you must manually tell the GPU to forget the images, or your VRAM will fill up.
        

### 🗺️ Phase 2: The Math Renderer (Textbook Quality)

**Goal:** Render beautiful equations ($ \int_0^\pi \sin(x) dx $) without bloating the app.

1. **The Rendering Pipeline:**
    
    - `egui` cannot draw math equations natively.
        
    - **Tool:** Use a library called **Typst** (via `typst` and `typst-render` crates). It is faster and lighter than LaTeX.
        
2. **The Caching Problem:**
    
    - Rendering an equation is slow (CPU heavy). You cannot do it every frame.
        
    - **Strategy:**
        
        1. Check if equation has changed.
            
        2. If yes, render to an Image (in a background thread).
            
        3. Upload Image to GPU as a Texture.
            
        4. Draw the Texture.
            
    - **Memory:** If the user changes the equation, immediately _delete_ the old texture from the GPU.
        

### 🗺️ Phase 3: Interactive Graphing (The Calculator)

**Goal:** A graphing calculator that runs at 60 FPS even with complex math.

1. **Using `egui_plot` Correctly:**
    
    - Do **not** calculate `y = sin(x)` for every pixel.
        
    - **Lesson:** Implement "Plot Callbacks." You give `egui_plot` a function closure `|x| x.sin()`. The plot widget decides how many points it needs based on the zoom level. This keeps memory usage near zero regardless of how far you zoom out.
        
2. **Custom Interactions:**
    
    - **Task:** Make a "Tangent Line" tool. When the user hovers over the graph, calculate the derivative at that point and draw a straight line.
        
    - **Math:** Use a crate like `fend-core` or `meval` to parse strings ("2x + 5") into math logic safely.
        

### 🗺️ Phase 4: Heavy Simulation (The Laboratory)

**Goal:** Run a physics simulation (e.g., Fluid Dynamics or N-Body Gravity) without freezing the UI.

1. **The "Compute Shader" (Expert Level):**
    
    - If you simulate 10,000 particles on the CPU, the UI will lag.
        
    - **Solution:** Use `wgpu` (the graphics engine behind `egui`) to run "Compute Shaders."
        
    - **Mechanism:**
        
        1. Send initial data to GPU.
            
        2. GPU calculates new positions (Simulate).
            
        3. **Zero-Copy:** Tell `egui` to draw the result _directly_ from the GPU memory using a `PaintCallback`.
            
        4. The CPU never sees the data again. This is maximum performance with minimum RAM.
            

### 🗺️ Phase 5: The "Notebook" Interface

**Goal:** Mix text, math, and interactive widgets in a scrolling list (like a Jupyter notebook).

1. **Virtualization:**
    
    - If a lesson has 100 steps, do not render all 100.
        
    - **Tool:** `egui_extras::TableBuilder` or `egui_virtual_list`.
        
    - **Concept:** Only run the math/logic for the items currently visible on the screen.
        

### ✅ Recommended "Khan Academy" Project: **Interactive Physics Lab**

Don't build a video player. Build an **Interactive Explainer**.

**Core Features to Build:**

1. **The Canvas:** A central area where you draw shapes.
    
2. **The Equation Box:** A text box where the user types `y = x^2 / mass`.
    
3. **The Slider:** A slider for `mass`.
    
4. **The Logic:**
    
    - Use `fend` to parse the text.
        
    - Use `egui_plot` to draw the curve.
        
    - Use `wgpu` to drop a "ball" that rolls down that curve.
        

Memory Challenge:

Ensure the app uses < 100MB of RAM even if the simulation runs for an hour. This requires strict cleanup of old textures and simulation buffers.