Day 1 was about Defense (Git hooks, CI/CD, preventing bad code).

Day 2 is about Connectivity (The "Walking Skeleton").

Your goal for Day 2 is to create a **"Steel Thread"**: a single vertical slice of functionality that cuts through the Database, Backend, API Layer, and Frontend.

If you skip this and start building "features" immediately, you will hit architectural walls (CORS errors, Type import failures) next week.

Here is your rigorous **Day 2 Roadmap**.

---

### 1. The Monorepo Workspace (Mandatory)

To make **Eden Treaty** work (where Frontend sees Backend types), you cannot keep them in separate unconnected folders. You must use **Bun Workspaces**.

Structure Change:

Refactor your root folder immediately. Move your existing code into packages/.

Plaintext

```
my-khan-clone/
├── package.json        # <--- Root config
├── bun.lockb
├── packages/
│   ├── api/            # Backend (Elysia + Drizzle)
│   └── web/            # Frontend (Vue + Vite)
└── .github/            # CI/CD stays at root
```

**The Root `package.json`:**

JSON

```
{
  "name": "khan-clone-monorepo",
  "workspaces": [
    "packages/*"
  ]
}
```

### 2. The Database Layer (Postgres + Drizzle)

Day 2 is when you prove you can talk to the database.

**In `packages/api`:**

1. **Install Drizzle:**
    
    Bash
    
    ```
    bun add drizzle-orm postgres
    bun add -d drizzle-kit
    ```
    
2. Strict Schema (src/db/schema.ts):
    
    Define a simple table to test the connection.
    
    TypeScript
    
    ```
    import { pgTable, serial, text, timestamp } from 'drizzle-orm/pg-core';
    
    export const users = pgTable('users', {
      id: serial('id').primaryKey(),
      email: text('email').notNull().unique(),
      createdAt: timestamp('created_at').defaultNow(),
    });
    ```
    
3. Migration Script:
    
    Run bunx drizzle-kit generate to create the SQL file.
    

### 3. The Backend "Contract" (Elysia)

You need to expose the API and the **Type Definition**.

**In `packages/api/src/index.ts`:**

TypeScript

```
import { Elysia } from 'elysia'
import { cors } from '@elysiajs/cors'
import { db } from './db' // Your drizzle connection
import { users } from './db/schema'

const app = new Elysia()
  .use(cors()) // <--- Critical: Allow Frontend to talk to Backend
  .get('/health', () => 'Server is alive')
  .get('/users', async () => {
    return await db.select().from(users)
  })
  .listen(3000)

// THE GLUE: Export the type for the frontend
export type App = typeof app 

console.log(`🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`)
```

### 4. The Frontend Integration (The "Magic" Step)

This is where you verify if your setup is "Ready to Scale." The Frontend must import the Backend's type without running the Backend code.

**In `packages/web`:**

1. **Add API as dependency:**
    
    Bash
    
    ```
    bun add @khan-clone/api@workspace:* ```
    *(Note: Ensure your api package.json has `"name": "@khan-clone/api"`)*
    
    ```
    
2. **Create the Client (`src/lib/api.ts`):**
    
    TypeScript
    
    ```
    import { treaty } from '@elysiajs/eden'
    // IMPORT TYPE ONLY - This adds 0kb to your bundle
    import type { App } from '@khan-clone/api/src/index' 
    
    export const client = treaty<App>('localhost:3000')
    ```
    
3. **The Test Component (`src/App.vue`):**
    
    Code snippet
    
    ```
    <script setup lang="ts">
    import { onMounted, ref } from 'vue'
    import { client } from '@/lib/api'
    
    const status = ref('Checking...')
    
    onMounted(async () => {
      // IDE Autocomplete should work here instantly
      const { data, error } = await client.health.get()
    
      if (data) status.value = data
      if (error) status.value = 'Connection Failed'
    })
    </script>
    
    <template>
      <h1>System Status: {{ status }}</h1>
    </template>
    ```
    

### 5. The "Smoke Test"

Run two terminals:

1. `cd packages/api && bun run dev`
    
2. `cd packages/web && bun run dev`
    

**Success Criteria:**

1. Frontend shows "System Status: Server is alive".
    
2. **IntelliSense works:** When you type `client.`, you see `.users`, `.health`, `.get`.
    
3. **Type Safety:** If you try to access `data.wrongProperty`, the build fails.
    

### Mentorship Warning

**Common Day 2 Failures:**

1. **CORS:** If you forget `.use(cors())` in Elysia, the browser will block the request.
    
2. **Ports:** Ensure API is on 3000 and Vite is on 5173 (default). Do not let them fight for the same port.
    
3. **Docker:** Ensure your Postgres container is actually running (`docker ps`) before starting the API.
    

### Next Step

Once you confirm the "Steel Thread" works (Frontend can read Backend types), we need to set up the **Database Docker Container** so you can actually save data.

Would you like the `docker-compose.yml` optimized for this Monorepo structure?