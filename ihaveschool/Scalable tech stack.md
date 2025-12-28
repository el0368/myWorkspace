Here is the strict, high-performance **Scalable Tech Stack** manifest.

### **1. Core Application**

|**Category**|**Tool**|
|---|---|
|**Frontend Framework**|**Vue 3** (Script Setup + Composition API)|
|**Language**|**TypeScript** (Strict Mode)|
|**Runtime**|**Bun**|
|**Build Tool**|**Vite**|
|**Package Manager**|**Bun** (Built-in)|

### **2. Frontend Ecosystem**

|**Category**|**Tool**|
|---|---|
|**State Management**|**Pinia**|
|**Routing**|**Vue Router**|
|**Styling**|**Tailwind CSS** (v4)|
|**Utilities**|**VueUse** (Standard library for Vue)|
|**API Client**|**Eden Treaty** (Elysia’s Type-Safe Client) or **Ky**|
|**Forms/Validation**|**VeeValidate** + **Zod**|

### **3. Backend & Data**

|**Category**|**Tool**|
|---|---|
|**Backend Framework**|**ElysiaJS** (Native Bun performance)|
|**Database**|**PostgreSQL** (v16+)|
|**ORM**|**Drizzle ORM**|
|**Caching / PubSub**|**Redis** (or DragonflyDB for high scale)|
|**Schema Validation**|**Zod**|
|**API Documentation**|**Scalar** (Built into Elysia)|

### **4. Domain Specific (Math/Text)**

|**Category**|**Tool**|
|---|---|
|**Math Rendering**|**KaTeX**|
|**Math Input**|**MathLive** (`mathlive`)|
|**Rich Text Editor**|**Tiptap**|
|**TTS/Speech**|**Web Speech API** (via `@vueuse/core`)|

### **5. Infrastructure & DevOps (Scaling)**

|**Category**|**Tool**|
|---|---|
|**Containerization**|**Docker** (distroless/cc-debian12 base)|
|**Orchestration**|**Kubernetes** (K8s) or **Coolify** (Self-hosted PaaS)|
|**Reverse Proxy**|**Nginx** or **Caddy**|
|**Object Storage**|**MinIO** (Self-hosted S3) or **Cloudflare R2**|
|**CI/CD**|**GitHub Actions**|

### **6. Quality Assurance**

|**Category**|**Tool**|
|---|---|
|**Unit Testing**|**Vitest**|
|**E2E Testing**|**Playwright**|
|**Linting**|**Biome** (Replaces ESLint + Prettier, much faster)|
|**Commit Standards**|**Husky** + **Commitlint**|

### Next Step

I can generate the **`package.json`** file with all these dependencies pre-configured so you can install the entire stack in one command. Would you like that?