Here is the rigorous way to check your infrastructure status.

Do not guess. Use the terminal to see exactly what is running.

### 1. The Global Check (Standard)

Run this command in your terminal:

Bash

```
docker ps
```

**How to read the output:**

- **UP (✅ Good):** You see a list of containers. Look at the `STATUS` column. It should say `Up X minutes` or `Up X hours`.
    
- **DOWN (❌ Bad):** You only see the header (`CONTAINER ID IMAGE COMMAND...`) and **no rows** underneath. This means nothing is running.
    

### 2. The Project-Specific Check (Better)

Since you are using `docker-compose.yml`, it is better to ask Docker about _this specific project_ rather than every container on your computer.

Run this inside your project folder:

Bash

```
docker compose ps
```

**The "Status" Decoder:**

|**Status Column**|**Meaning**|**Action Required**|
|---|---|---|
|`Up (...ports)`|**System is Healthy.**|None. Proceed to code.|
|`Exited (0)`|**Stopped Gracefully.**|Run `docker compose up -d` to start.|
|`Exited (1)`|**CRASHED.**|The container tried to start but died (error).|
|`Restarting...`|**Boot Looping.**|It is crashing, retrying, and crashing again.|

### 3. How to fix "Down" or "Crashed"

If the status is not "Up", run these commands strictly in order:

**Option A: It is just stopped**

Bash

```
docker compose up -d
```

Option B: It is crashing (Status: Exited 1)

You need to see why it died. Run the logs:

Bash

```
docker compose logs db
```

_(Replace `db` with the service name from your docker-compose file)._

### Summary Checklist

1. Run `docker ps`.
    
2. Look for your database container name (e.g., `khan-clone-db-1`).
    
3. Check the `STATUS` says `Up`.