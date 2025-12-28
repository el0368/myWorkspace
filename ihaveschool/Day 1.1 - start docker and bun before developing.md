**Yes.**

You must always run 

```
docker-compose up -d
```

 before you start developing (

```
bun run dev
```

). The application needs the database to be running, or it will crash.

---

### How to Copy to External Hard Disk (The Safe Way)

To copy your project folder 

```
c:\GitHub\i_have_school
```

 to your external drive without it taking hours:

**1. Clean the Project (On your PC)** Run this specific command in your terminal to delete the thousands of tiny dependency files. This reduces the size from ~500MB to ~2MB.

powershell

# Stop your app (Ctrl+C) and Docker first!

docker-compose down

# Delete the heavy folders

rm -r -force node_modules

rm -r -force .bun

rm -r -force dist

# also inside the apps folders if they exist

rm -r -force apps/web/node_modules

rm -r -force apps/api/node_modules

**2. Copy via File Explorer**

1. Open **File Explorer**.
2. Go to 
    
    ```
    c:\GitHub
    ```
    
    .
3. Right-click 
    
    ```
    i_have_school
    ```
    
     -> **Copy**.
4. Go to your External Hard Disk (e.g., 
    
    ```
    D:\
    ```
    
     or 
    
    ```
    E:\
    ```
    
    ).
5. Right-click -> **Paste**.

**3. Move to New PC**

1. Plug the External Drive into the new PC.
2. **Copy** 
    
    ```
    i_have_school
    ```
    
     from the External Drive to the new PC's **Internal SSD** (e.g., 
    
    ```
    Documents
    ```
    
     or 
    
    ```
    Desktop
    ```
    
    ).
3. Open the folder in terminal on the new PC and run 
    
    ```
    bun install
    ```
    
     to download fresh dependencies.