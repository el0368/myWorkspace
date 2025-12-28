This is the **SHAX Stack Library Roadmap**. Since we are transitioning away from the Decoupled Vue SPA toward a high-performance **Svelte + HTMX + Axum** architecture, this plan integrates every level of mathematics and physics logic while maintaining the "unbreakable" stability you require.

### 1. The Core Infrastructure (The Foundation)

These libraries handle the basic look, feel, and data of the project.

- **Styling:** Tailwind CSS v4 + DaisyUI 5 (handled via your existing `input.css`).
    
- **Server Logic:** Axum 0.7 + SQLx (PostgreSQL) for type-safe data and heavy mathematical pre-processing.
    
- **The Glue:** HTMX for swapping the "Miller Column" fragments without a full page reload.
    
- **Interactivity:** Svelte 5 (Islands) to handle the complex UI logic within those columns.
    

---

### 2. Tier 1: Basic Math & Rendering (The Visuals)

This tier ensures that mathematical notation is beautiful and that basic calculations are handled instantly.

- **KaTeX:** The standard for rendering LaTeX formulas into beautiful symbols. This ensures your curriculum nodes look professional.
    
- **Math.js:** The "Swiss Army Knife" of math. It handles complex numbers, units, and matrices. It will be the primary engine for standard exercise validation.
    

---

### 3. Tier 2: Advanced Symbolic Math (The Thinking)

For "End of the World" scale curriculum (Calculus, Algebra, Engineering), you need libraries that can solve problems, not just calculate numbers.

- **Nerdamer / Algebrite:** These are Symbolic Computation engines. They can simplify expressions (e.g., $x + x = 2x$), solve for $x$, and handle derivatives and integrals.
    
- **Numeric.js:** Used specifically for advanced linear algebra and numerical analysis.
    
- **Function-plot:** A dedicated library for creating interactive 2D function graphs (e.g., plotting $y = sin(x)$).
    

---

### 4. Tier 3: Advanced 2D & Physics (The Simulation)

This tier brings your "Exercise Factory" to life with interactive geometry and physical laws.

- **Native Svelte SVG:** For the most minimalist and high-performance 2D geometry exercises.
    
- **LayerCake:** Used for building responsive, data-driven charts that scale across any screen size without breaking the layout.
    
- **Matter.js:** A 2D physics engine. Perfect for "Physical Math" exercises involving gravity, friction, and collisions.
    
- **PixiJS:** Used only if a node requires 10,000+ interactive elements (High-performance 2D rendering).
    
- **Fabric.js:** The engine for your "Exercise Studio" whiteboard, allowing students to drag, drop, and draw.
    

---

### 5. Tier 4: 3D Space (The Multiverse)

For advanced geometry, 3D plotting, and spatial reasoning.

- **Threlte (Three.js for Svelte):** The powerhouse for 3D islands. It allows you to build 3D multivariable calculus graphs or physics simulations that run directly inside your Svelte components.
    

---

### The "SHAX" Implementation Plan

To prevent the "layout breaking" issues of the past, we use the **Island Architecture**:

1. **Axum (The Boss):** Receives a request and fetches the curriculum data from PostgreSQL. It renders a standard HTML fragment containing an HTMX trigger.
    
2. **HTMX (The Delivery):** Swaps that fragment into your "Miller Column." Inside the fragment is a custom Svelte tag (e.g., `<physics-exercise />`).
    
3. **Svelte (The Fortress):** The Svelte Island "wakes up." It uses **Matter.js** for the physics and **Nerdamer** for the algebra. Because its CSS is **scoped**, the physics simulation can never mess up your main header or sidebar.
    

**Expert Opinion:** This library stack is "Enterprise-Ready." By splitting the libraries into specialized Svelte islands, your app stays lightweight because the user only downloads the "3D Engine" or "Physics Engine" when they actually reach a node that needs it.

**Would you like me to update your `GEMINI.md` (The Constitution) to replace the Vue 3 rules with this "SHAX + Math/3D" library plan?** This will lock these choices in so we can start the final implementation.