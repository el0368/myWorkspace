**It is 100% GUI.**

You are building a **Graphical User Interface** (using the `egui` graphics engine). It just happens to be "dressed up" to **look** like a TUI.

Here is the breakdown to clear up the confusion:

### 1. The Technology (It is a GUI)

- **What runs underneath:** Your graphics card (GPU) is drawing pixels, lines, and rectangles.
    
- **The "Columns":** In a real TUI (Terminal), columns are just spaces between text characters. In your app, the "Columns" are invisible, interactive **containers**.
    
    - You can **drag the edge** of a column to make it wider (GUI feature).
        
    - You can **click the header** "Price" to sort (GUI feature).
        
    - Real text terminals cannot do this easily. Your app can.
        

### 2. The Visual Style (It looks like a TUI)

- **Why we call it "TUI Style":** We are stripping away the "fat."
    
    - Standard GUI: Buttons have shadows, rounded corners, and padding.
        
    - Your GUI: We set `rounding: 0px` and `shadow: None`.
        
- **The Result:** It _looks_ flat and boring (like a terminal), but it is essentially a high-performance spreadsheet.
    

### 3. "Left to Right" Explanation

When I listed "The Columns (Left to Right)," I was describing the **Layout Order**.

- **Column 1 (Far Left):** Symbol (e.g., "AAPL")
    
- **Column 2:** Price (e.g., "150.00")
    
- **Column 3:** Change (e.g., "+1.5%")
    
- **...and so on.**
    

In your code (using `TableBuilder`), you will tell the GUI: _"Draw the Symbol first. Then draw the Price next to it."_

### Summary

- **Is it a Terminal?** No. It is a window.
    
- **Can I use a mouse?** Yes.
    
- **Why does it look like that?** To fit more data on the screen.
    

**Are you ready to create the `data.rs` file now?**