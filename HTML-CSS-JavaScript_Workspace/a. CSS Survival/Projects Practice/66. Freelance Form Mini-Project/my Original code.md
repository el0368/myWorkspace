```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>Document</title>
    <link rel="stylesheet" href="style.css" />
</head>

<body>
    <h1>HTML & CSS</h1>
    <div class="container">
        <div class="inner-detail">
            <h2>Need to hire a freelancer?</h2>
            <p>Sign up, type in what you need and receive free quotes in seconds.</p>
            <input type="text" name="username" id="username" />
            <div class="button">
                <p>Post a Project</p>
            </div>
        </div>
    </div>
</body>

</html>
```

```css
/* * {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
} */

body {
  font-family: "Poppins", sans-serif;
}

.container {
  background-color: black;
  color: white;
  width: 300px;
  /* border: 10px solid red; */
  margin: 50px auto;
  padding: 30px;
}

.inner-detail {
  /* border: 2px dashed whitesmoke; */
  padding: 20px 0;

  display: flex;
  flex-direction: column;
  align-items: center;
}

.inner-detail p {
  margin-bottom: 20px;
  text-align: center;
}

.inner-detail input {
  height: 30px;
  width: 100%;
}
.button {
  background-color: orange;
  margin-top: 20px;
  height: 40px;
  border: 1px solid black;

  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
```