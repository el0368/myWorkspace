You are absolutely right to think of this now. "Day 1" is the only time you can set this up easily. If you wait until Day 50, it becomes a painful migration.

Here is your **Day 1 Automation & Quality Assurance Manifest**.

### 1. The "Quality Gate" (Local Git Hooks)

You must prevent bad code from ever leaving your computer. We will use **Husky** (to trigger hooks) and **Commitlint** (to force strict commit messages).

**Run these commands immediately:**

Bash

```
# 1. Install dev tools
bun add -d husky @commitlint/{cli,config-conventional} lint-staged

# 2. Init Husky (creates .husky folder)
bunx husky init

# 3. Create the 'Commit Message' hook (Forces "feat: add login" format)
echo 'bunx --no -- commitlint --edit $1' > .husky/commit-msg

# 4. Create the 'Pre-Commit' hook (Runs linting & types before you can commit)
echo 'bunx lint-staged' > .husky/pre-commit
```

Configure package.json for lint-staged:

This ensures you only lint the files you changed, not the whole project (which is slow).

JSON

```
"lint-staged": {
  "*.{ts,tsx,vue}": [
    "biome check --write",  // Format & Lint
    "bunx tsc --noEmit"     // Type Check
  ]
}
```

Configure commitlint.config.ts:

Create this file in your root. It forces you to write feat:, fix:, docs: or the commit fails.

TypeScript

```
export default { extends: ['@commitlint/config-conventional'] }
```

---

### 2. The "CI/CD Pipeline" (GitHub Actions)

You need a robot that tests your code every time you push. Create `.github/workflows/ci.yml`.

**Strict Rules Implemented Here:**

1. **Fast Fail:** If the linter fails, don't bother running tests.
    
2. **Caching:** Caches `bun install` so it runs in 2 seconds, not 30.
    
3. **Matrix:** (Optional) You can test on Linux and Windows, but for now, Linux is enough.
    

YAML

```
name: CI Quality Gate

on:
  push:
    branches: [ main, dev ]
  pull_request:
    branches: [ main, dev ]

jobs:
  quality:
    name: 🛡️ Lint, Type & Test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Bun
        uses: oven-sh/setup-bun@v1
        with:
          bun-version: latest

      - name: Install Dependencies
        run: bun install --frozen-lockfile

      - name: 🎨 Lint & Format (Biome)
        run: bunx biome ci .

      - name: 📐 Type Check (TSC)
        run: bunx tsc --noEmit

      - name: 🧪 Unit Tests (Vitest)
        run: bun run test:unit

      - name: 🏗️ Build Check (Vite)
        run: bun run build
```

---

### 3. GitHub Repository Settings (The "Safety Net")

You must configure your GitHub repository to **enforce** the tools you just built.

- Go to **Settings** -> **Branches** -> **Add branch protection rule**.
    
- **Branch name pattern:** `main`
    
- Check **Require a pull request before merging**.
    
- Check **Require status checks to pass before merging**.
    
    - Search for and select: `🛡️ Lint, Type & Test` (The job name from your YAML above).
        
- Check **Require linear history** (Optional, but keeps Git history clean).
    

---

### 4. The `.gitignore` (Bun Specific)

Don't use a standard Node `.gitignore`. Use this strict Bun version to avoid committing unnecessary binary locks or caches.

Plaintext

```
# Dependencies
node_modules
.bun/

# Build
dist
.output

# Environment
.env
.env.local

# Testing
coverage

# OS
.DS_Store
Thumbs.db

# Logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
```

---

### 5. Automatic Versioning (Semantic Release)

Since you want "Automatic," setting up **Semantic Release** allows you to never manually type a version number again.

- If you commit `fix: bug`, it bumps `1.0.0` -> `1.0.1`
    
- If you commit `feat: new thing`, it bumps `1.0.0` -> `1.1.0`
    

Add to your GitHub Action (after the build step):

Note: This requires a GH_TOKEN in your repo secrets.

YAML

```
      - name: 🚀 Release
        if: github.ref == 'refs/heads/main'
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: bunx semantic-release
```

### Next Step

I have the **Docker** setup ready (which effectively acts as your "Production Environment" test). Would you like the `Dockerfile` and `docker-compose.yml` now to complete your "Day 1" infrastructure?