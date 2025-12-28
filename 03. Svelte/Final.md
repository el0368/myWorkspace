The new stack you’ve landed on is called the **SPELP** stack.

It replaces the complexity of Rust and the "JSON tax" of traditional Go/Svelte apps with a unified, server-driven flow.

### The SPELP Stack Summary

|**Initial**|**Technology**|**Purpose**|
|---|---|---|
|**S**|**Svelte 5**|High-performance "Interactive Islands" for 3D/2D math and exercises.|
|**P**|**Phoenix**|The web framework that manages the server, real-time sync, and layout.|
|**E**|**Elixir**|The functional language that makes the backend "unbreakable" and self-healing.|
|**L**|**LiveView**|The glue. It syncs state between the server and the browser via WebSockets (No APIs needed).|
|**P**|**Python**|The AI Brain. It runs your Agentic AI course logic (LangGraph, CrewAI) as a sidecar service.|

---

### Key Libraries & Tools

#### 1. Backend (Elixir/Phoenix)

- **LiveView:** Handles the Miller Column navigation and page state.
    
- **LiveSvelte:** The specific bridge that lets you drop Svelte components into Elixir pages seamlessly.
    
- **Ecto:** The database wrapper for PostgreSQL. It handles your hierarchical curriculum data safely.
    
- **Oban:** Manages the background jobs for your Python AI agents (so long AI "thinking" doesn't freeze the UI).
    
- **Phoenix LiveDashboard:** Built-in real-time monitoring of your entire system.
    

#### 2. Frontend (Svelte)

- **Svelte 5:** Uses "Runes" for ultra-fast, reactive math visualizations.
    
- **Tailwind CSS v4:** Scoped styling to ensure exercise UI never breaks the main layout.
    
- **Three.js / PixiJS:** (Optional) For the 3D/2D coordinate planes inside your Svelte islands.
    

#### 3. AI Layer (Python)

- **FastAPI:** A lightweight bridge to allow Elixir to "talk" to your AI agents.
    
- **Pydantic / LangGraph:** For structured AI outputs and agentic workflows from your course.
    

---

### Why this is the "Conclusion"

The **SPELP** stack solves your two biggest pain points:

1. **Broken Layouts:** Because **Phoenix/LiveView** owns the state, the layout is mathematically stable. If a Svelte island crashes, the "Supervisor" restarts it without touching the rest of the page.
    
2. **The Flow Gap:** By using **LiveView + LiveSvelte**, you stop writing JSON APIs. The data flows from your Python AI, through Elixir, and into Svelte as if they were all the same language.
    

This is a professional-grade, solo-developer architecture. It’s built to handle complex hierarchies and real-time AI updates with the least amount of "plumbing" code possible.