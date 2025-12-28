# Day 1 User Manual: Scalable Tech Stack

This guide covers the 3 critical workflows: **Coding/Testing**, **Database Management**, and **Safe Transport**.

## 1. How to Run & Test (The "Dev Loop")

**Goal:** working on features, changing code, and seeing results instantly.

1. **Start the Database (Once per session):** Open a terminal and run:
    
    docker-compose up -d
    
    _This starts PostgreSQL and pgAdmin in the background._
    
2. **Start the App (The "Magic" Command):** Run this in your root folder:
    
    bun run dev
    
    - **Frontend**: Open [http://localhost:5173](http://localhost:5173/) (or 5174/5175 if port taken)
    - **Backend**: API running at [http://localhost:3000](http://localhost:3000/)

**Q: Do I need to restart Docker when I change code?**

- **No.** Vue and Elysia (Backend) have "Hot Reloading". Change a file -> Save -> It updates instantly.
- **Yes, ONLY IF**:
    - You change `.env` variables.
    - You change 
        
        docker-compose.yml.
    - You delete the database to start fresh.

---

## 2. Docker & Database Management

**Tool:** pgAdmin (GUI) is clearer than command line for beginners.

- **URL**: [http://localhost:5050](http://localhost:5050/)
- **Login**: `admin@local.test` / `admin`
- **Add Server**:
    - Hostname: 
        
        db (This is the service name in docker-compose)
    - Username/Password: `postgres` / `postgres`

**Useful Commands:**

- **Stop Everything**: `docker-compose down` (Safe shutdown)
- **Reset Database (Delete Data)**:
    
    docker-compose down
    
    rm -rf db-data  # WARNING: Deletes all saved data!
    
    docker-compose up -d
    

---

## 3. The "Pack & Move" Protocol (Transport)

**Goal:** Move project from PC A -> External Drive -> PC B.

**Step 1: PACK (On PC A)**

1. **Stop Docker**: `docker-compose down`
2. **Clean "Trash"** (Crucial!):
    - Delete `node_modules` (in root, apps/web, apps/api)
    - Delete `.bun`
    - Delete `dist`
    - _Note: Do NOT delete `db-data` if you want to keep your database!_
3. **Copy**: Drag the folder `i_have_school` to your External Hard Disk.

**Step 2: UNPACK (On PC B)**

1. **Copy**: Drag `i_have_school` from External Drive -> **Internal SSD**.
2. **Rehydrate**:
    
    bun install
    
3. **Start**:
    
    docker-compose up -d
    
    bun run dev