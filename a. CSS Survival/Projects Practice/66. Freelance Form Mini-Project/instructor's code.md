Sure! Below is your updated code separated into **HTML** and **CSS**, each inside its own copy-friendly code box.

---

### ✅ HTML (`index.html`)

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <link rel="stylesheet" href="style.css" />
</head>

<body>
    <main class="container">
        <div class="card">
            <h1>Need to hire a freelancer?</h1>
            <p>Sign up, type in what you need and receive free quotes in seconds.</p>

            <form>
                <div class="form-group">
                    <input type="text" name="username" id="username" placeholder="What do you need done?" />
                </div>
                <div class="form-group">
                    <button type="submit">Post a Project</button>
                </div>
            </form>
        </div>
    </main>
</body>

</html>
```

---

### ✅ CSS (`style.css`)

```css
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: "Poppins", sans-serif;
  background: #f1f1f1;
}

.container {
  width: 400px;
  margin: 150px auto;
}

.card {
  background: #373f4d;
  color: #fff;
  border-radius: 5px;
  text-align: center;
  padding: 60px 50px;
  box-shadow: 4px 4px 6px rgba(0, 0, 0, 0.4);
}

.card h1 {
  font-size: 22px;
  margin-bottom: 15px;
  font-weight: 300;
}

.card p {
  font-size: 15px;
  margin-bottom: 30px;
}

.card .form-group {
  margin-bottom: 20px;
}

.card input[type="text"] {
  width: 100%;
  padding: 15px 10px;
  font-size: large;
  border: none;
  background-color: #f1f1f1;
}

.card button {
  width: 100%;
  padding: 15px 10px;
  font-size: large;
  border: none;
  background: linear-gradient(
    90deg,
    rgba(242, 132, 62, 1) 21%,
    rgba(239, 111, 69, 1) 61%
  );
  color: #fff;
  cursor: pointer;
}

.card button:hover {
  background: linear-gradient(
    90deg,
    rgba(232, 132, 62, 1) 21%,
    rgba(229, 111, 69, 1) 61%
  );
}
```

---

Let me know if you’d like a version with media queries or dark mode toggle!