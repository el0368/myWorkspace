Here is the **Advanced Roadmap**.

This level is for people building **Game Engines**, **CAD Tools**, or **Production Desktop Apps** that feel native. We move away from "using the library" to "hacking the library."

### 🗺️ Phase 1: The Integrator (No More Frameworks)

**Goal:** Run `egui` without `eframe`. This is required if you want to add UI to an existing game engine or a custom renderer.

1. **The Raw Loop:**
    
    - Stop using `eframe::App`.
        
    - **Task:** Create a window using `winit` (the raw window library).
        
    - **Task:** Manually create the `egui::Context`.
        
    - **Task:** Feed events: Translate `winit` events into `egui::RawInput` manually.
        
    - **Task:** Handle output: Take `egui::FullOutput`, look at the `textures_delta`, and upload textures to the GPU yourself.1
        
2. **The Renderer (WGPU/OpenGL):**
    
    - `eframe` does the painting for you. Now _you_ must do it.
        
    - **Task:** Set up a `wgpu` render pipeline.
        
    - **Task:** Take the mesh (triangles) from `egui` and draw them in your custom pipeline.
        

### 🗺️ Phase 2: The 3D Embedder (Mixed Reality)

**Goal:** Render a 3D scene _inside_ a specific `egui` widget. This is how tools like Blender or Unity are made.

1. **The `PaintCallback`:**
    
    - `egui` usually draws 2D triangles. You can tell it: "When you get to this rectangle, stop drawing UI and run this custom GPU command."
        
    - **Task:** Create a "Viewport Widget".
        
    - **Task:** Inside that widget, render a spinning 3D cube using `glow` (OpenGL) or `wgpu`.
        
    - **Task:** Ensure the 3D cube responds to `egui` input (dragging the mouse over the cube rotates it).
        

### 🗺️ Phase 3: The Window Manager (Viewports)

**Goal:** Build multi-window applications (pop-out tabs).

1. **Immediate Viewports:**
    
    - Learn to use `ctx.show_viewport_immediate`.
        
    - **Task:** Create a button "Pop Out". When clicked, the content moves from the main window to a _new_ native OS window.
        
    - **Challenge:** Syncing state between two separate OS windows (since they share one `Context` memory).
        
2. **Docking (The Hard Part):**
    
    - `egui` does not have a built-in docker (like Visual Studio).
        
    - **Task:** Integrate `egui_dock`. Learn to manage a "Tree" of tabs that can be split, resized, and dragged between windows.
        

### 🗺️ Phase 4: Advanced Interaction Patterns

**Goal:** Features users expect from professional software.

1. **Undo/Redo Stack:**
    
    - `egui` is immediate mode, so it forgets everything next frame. You must build memory.
        
    - **Task:** Implement the **Command Pattern**. Store a stack of "Actions" (e.g., `Action::ColorChange { old: Red, new: Blue }`).
        
    - **Task:** Press Ctrl+Z to pop the stack and apply the `old` value.
        
2. **Complex Drag and Drop:**
    
    - Not just reordering a list.
        
    - **Task:** Drag a file from your **Desktop** (Windows/Linux) into your `egui` app.
        
    - **Task:** Drag an item from _Window A_ and drop it into _Window B_ (using the payload system).
        

### 🗺️ Phase 5: Deep Optimization

**Goal:** 1,000,000 items at 60 FPS.

1. **Memory Inspection:**
    
    - Use `ctx.memory()` to see what `egui` is remembering.
        
    - **Task:** Identify "stale" IDs (widgets that no longer exist but are taking up RAM) and clean them.
        
2. **Custom Virtualization:**
    
    - `TableBuilder` is good, but what if items have different heights?
        
    - **Task:** Write a custom layout algorithm that only calculates the size of items visible in the `scroll_delta`.
        

### ✅ The "Advanced" Final Project

**Project: A 3D Level Editor (Mini-Unity)**

1. **Center:** A 3D Viewport (Phase 2) rendering a scene.
    
2. **Right Panel:** A "Inspector" that shows properties of the selected 3D object.
    
3. **Bottom Panel:** A "File Browser" that supports dragging files into the 3D view (Phase 4).
    
4. **Feature:** You can right-click the Inspector tab and "Undock" it to a second monitor (Phase 3).
    
5. **Architecture:** Running on your own `winit` event loop, not the default `eframe` one (Phase 1).