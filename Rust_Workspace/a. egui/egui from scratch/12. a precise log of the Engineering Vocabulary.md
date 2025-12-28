As your Project Manager, I keep a precise log of the **Engineering Vocabulary** we use.

Using the correct terminology is not just about sounding smart; it is about **precision**. When you say "Content Bleeding," every senior engineer knows exactly what needs to be fixed (clipping rectangles). When you say "The text looks weird," nobody knows what to do.

Here is the **Official Glossary** of the terms we have used to build your Trading Terminal, categorized by our development phases.

### 🏗️ Architecture & Structure

1. **Monorepo / Workspace:**
    
    - _Definition:_ A single repository that contains multiple distinct projects (Libraries + Apps).
        
    - _Our Use:_ We used this to keep `main_terminal` separate from `stock_scanner` and `sec_engine`.
        
2. **Decomposition:**
    
    - _Definition:_ Breaking a complex system into small, isolated, manageable parts.
        
    - _Our Use:_ We didn't build one giant `main.rs`; we broke it into "The Brain," "The Body," and "The Fuel."
        
3. **Decoupling:**
    
    - _Definition:_ Ensuring that Module A (UI) does not know how Module B (Data) works. They only talk via a "Contract."
        
    - _Our Use:_ The Scanner UI doesn't know if the data comes from the SEC or a Mock Generator. It just asks for a `StockQuote`.
        
4. **The "Walking Skeleton":**
    
    - _Definition:_ The thinnest possible version of an end-to-end system that actually runs.
        
    - _Our Use:_ We built a window that displayed fake data _before_ we wrote any complex math.
        

### ⚡ Concurrency & Data Flow

5. **Async Bridge:**
    
    - _Definition:_ A mechanism to connect a Synchronous system (UI, runs every frame) with an Asynchronous system (Network, waits for data).
        
    - _Our Use:_ We used `tokio::spawn` and `mpsc::channels` to fetch SEC data without freezing the window.
        
6. **Blocking vs. Non-Blocking:**
    
    - _Definition:_ **Blocking** stops the entire program until a task finishes (Freezing). **Non-Blocking** lets the program continue while the task runs in the background.
        
    - _Our Use:_ The "Spinner" animation proves our UI is non-blocking.
        
7. **MPSC Channel (Multi-Producer, Single-Consumer):**
    
    - _Definition:_ A one-way communication pipe. Many workers can throw messages in, but only one receiver (the UI) reads them.
        
    - _Our Use:_ The background threads throw `ScanComplete` messages into the pipe; the UI reads them to update the table.
        

### 🎨 User Interface (GUI)

8. **Content Bleeding:**
    
    - _Definition:_ When UI elements (text, buttons) overflow their container boundaries and cover neighboring elements.
        
    - _Our Use:_ The SEC Feed text was spilling into the AI Panel when resized.
        
9. **Strict Clipping:**
    
    - _Definition:_ A rendering rule that says, "If a pixel is outside this box, do not draw it." effectively cutting off the bleeding content.
        
    - _Our Use:_ We added `ui.set_clip_rect()` to the panels to fix the bleeding.
        
10. **Immediate Mode GUI (egui):**
    
    - _Definition:_ A UI paradigm where the interface is redrawn from scratch every single frame (60 times a second). It has no "memory" of the past layout.
        
    - _Our Use:_ This is why we don't need complex "callbacks" for buttons. We just check `if ui.button().clicked()`.
        
11. **Viewport:**
    
    - _Definition:_ The physical window provided by the Operating System.
        
    - _Our Use:_ We configured the Viewport to start at 1600x900 resolution.
        

### 🧠 Logic & Strategy

12. **Mocking / Mock Data:**
    
    - _Definition:_ Fake data used to simulate real-world inputs for testing purposes.
        
    - _Our Use:_ We created `MockDataFeed` to simulate AAPL and TSLA prices so we could build the UI before paying for a real API.
        
13. **Regex (Regular Expression):**
    
    - _Definition:_ A sequence of characters that defines a search pattern.
        
    - _Our Use:_ We used `r"\(([A-Z]+)\)"` to surgically extract the ticker "AAPL" from the messy string "10-Q - Apple Inc. (AAPL)".
        

Recommendation:

Keep this list. As you grow as a Software Engineer, these terms will be the "Language" you use to design systems with other engineers.

**Shall we proceed to the final step of our session: Packaging the App or adding the last placeholder logic?**