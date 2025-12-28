# Day 1 User Manual: Scalable Tech Stack

This guide covers the 3 critical workflows: **Coding/Testing**, **Database Management**, and **Safe Transport**.

## 1. How to Run & Test (The "Dev Loop")

**Goal:** working on features, changing code, and seeing results instantly.

### 1. Start Working

1. Open Terminal.
2. Run `bun run dev`.
3. Open `http://localhost:5173`.
4. **Database?** It runs automatically in the background (Windows Service). You do not need to start it.

### 2. Stop Working

1. Press `Ctrl+C` in the terminal to stop the website.
2. Close VS Code.
3. **Shutdown PC.**
    - **Do I need to backup?** NO.
    - **Do I need to stop the database?** NO.
    - Your data is automatically saved to `C:\Program Files\PostgreSQL`.

---

## 2. Moving to a New PC (The "Pack & Move" Protocol)

**Goal:** Move project from PC A -> External Drive -> PC B.

### Step 1: PACK (On This PC)

Only do this when you are ready to move.

1. **Backup Data**:
    
    cd apps/api
    
    bun run db:backup
    
    _This creates `school_db_backup.dump` in your project folder._
    
2. **Copy Project**: Drag the entire `i_have_school` folder (including the dump file) to your External Hard Disk.
    

### Step 2: UNPACK (On New PC)

1. **Copy**: Drag `i_have_school` from External Drive -> **Internal SSD**.
2. **Install Dependencies**:
    
    bun install
    
3. **Start Docker Database**:
    
    docker-compose up -d
    
4. **Restore Data** (One time only):
    
    cd apps/api
    
    bun run db:restore
    
5. **Start App**:
    
    bun run dev
    

---

## 3. Database Management (Advanced)

If you ever need to reset everything and start fresh:

**On This PC (Native):** run `bun run db:push` in `apps/api` to reset schema (data stays).

**On New PC (Docker):**

docker-compose down

rm -rf db-data

docker-compose up -d