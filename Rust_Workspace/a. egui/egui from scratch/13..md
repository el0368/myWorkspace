
### 📋 Engineering Vocabulary Log

1. **Deprecation:** The act of marking a feature (Gap/Volume scanners) as obsolete and removing it from the active codebase.
    
2. **Topology:** The structural arrangement of the UI elements (Grid Layout).
    
3. **Refactor:** Improving the internal structure of the code (cleaning up the grid) without changing its external behavior (logic remains the same).
    
4. **Flex Factor:** In UI layout, the ratio of space a component takes up relative to its container (e.g., `Size::relative(0.33)`).
    
5. **Event Loop:** The `update` function that runs every frame, checking for messages and redrawing the screen.
    
6. **Data Ingestion:** The process of fetching raw data (from `MockDataFeed`) and bringing it into the system.
    
7. **Heuristics:** Simple, practical rules used to solve problems (e.g., "Cash Burn < 3mo" is a heuristic for Urgency).
    

### 🚀 Verification Protocol

1. **Replace** `main.rs`.
    
2. **Run:** `cargo run -p main_terminal`.
    
3. **Validate:**
    
    - Do you see **6 Panels** (3 Top, 3 Bottom)?
        
    - Are the "Gap Up" and "High Volume" panels **Gone**?
        
    - Is the "Main Scanner" (Top Left) and "Small Cap Scanner" (Bottom Left) still active?
        

**Confirm deployment status.**