'use strict';

  

// Selecting elements

const score0El = document.querySelector('#score--0');

const score1El = document.getElementById('score--1');

const diceEl = document.querySelector('.dice');

  

const btnNew = document.querySelector('.btn--new');

const btnRoll = document.querySelector('.btn--roll');

const btnHold = document.querySelector('.btn--hold');

  

// Starting conditions

score0El.textContent = 0;

score1El.textContent = 0;

diceEl.classList.add('hidden');

  

// Rolling dice functionality

btnRoll.addEventListener('click', function () {

  const dice = Math.trunc(Math).random() * 6 + 1;

});

@import url('https://fonts.googleapis.com/css2?family=Nunito&display=swap');

  

* {

  margin: 0;

  padding: 0;

  box-sizing: inherit;

}

  

html {

  font-size: 62.5%;

  box-sizing: border-box;

}

  

body {

  font-family: 'Nunito', sans-serif;

  font-weight: 400;

  height: 100vh;

  color: #333;

  background-image: linear-gradient(to top left, #753682 0%, #bf2e34 100%);

  display: flex;

  align-items: center;

  justify-content: center;

}

  

/* LAYOUT */

main {

  position: relative;

  width: 100rem;

  height: 60rem;

  background-color: rgba(255, 255, 255, 0.35);

  backdrop-filter: blur(200px);

  filter: blur();

  box-shadow: 0 3rem 5rem rgba(0, 0, 0, 0.25);

  border-radius: 9px;

  overflow: hidden;

  display: flex;

}

  

.player {

  flex: 50%;

  padding: 9rem;

  display: flex;

  flex-direction: column;

  align-items: center;

  transition: all 0.75s;

}

  

/* ELEMENTS */

.name {

  position: relative;

  font-size: 4rem;

  text-transform: uppercase;

  letter-spacing: 1px;

  word-spacing: 2px;

  font-weight: 300;

  margin-bottom: 1rem;

}

  

.score {

  font-size: 8rem;

  font-weight: 300;

  color: #c7365f;

  margin-bottom: auto;

}

  

.player--active {

  background-color: rgba(255, 255, 255, 0.4);

}

.player--active .name {

  font-weight: 700;

}

.player--active .score {

  font-weight: 400;

}

  

.player--active .current {

  opacity: 1;

}

  

.current {

  background-color: #c7365f;

  opacity: 0.8;

  border-radius: 9px;

  color: #fff;

  width: 65%;

  padding: 2rem;

  text-align: center;

  transition: all 0.75s;

}

  

.current-label {

  text-transform: uppercase;

  margin-bottom: 1rem;

  font-size: 1.7rem;

  color: #ddd;

}

  

.current-score {

  font-size: 3.5rem;

}

  

/* ABSOLUTE POSITIONED ELEMENTS */

.btn {

  position: absolute;

  left: 50%;

  transform: translateX(-50%);

  color: #444;

  background: none;

  border: none;

  font-family: inherit;

  font-size: 1.8rem;

  text-transform: uppercase;

  cursor: pointer;

  font-weight: 400;

  transition: all 0.2s;

  

  background-color: white;

  background-color: rgba(255, 255, 255, 0.6);

  backdrop-filter: blur(10px);

  

  padding: 0.7rem 2.5rem;

  border-radius: 50rem;

  box-shadow: 0 1.75rem 3.5rem rgba(0, 0, 0, 0.1);

}

  

.btn::first-letter {

  font-size: 2.4rem;

  display: inline-block;

  margin-right: 0.7rem;

}

  

.btn--new {

  top: 4rem;

}

.btn--roll {

  top: 39.3rem;

}

.btn--hold {

  top: 46.1rem;

}

  

.btn:active {

  transform: translate(-50%, 3px);

  box-shadow: 0 1rem 2rem rgba(0, 0, 0, 0.15);

}

  

.btn:focus {

  outline: none;

}

  

.dice {

  position: absolute;

  left: 50%;

  top: 16.5rem;

  transform: translateX(-50%);

  height: 10rem;

  box-shadow: 0 2rem 5rem rgba(0, 0, 0, 0.2);

}

  

.player--winner {

  background-color: #2f2f2f;

}

  

.player--winner .name {

  font-weight: 700;

  color: #c7365f;

}

<!DOCTYPE html>

<html lang="en">

  <head>

    <meta charset="UTF-8" />

    <meta name="viewport" content="width=device-width, initial-scale=1.0" />

    <meta http-equiv="X-UA-Compatible" content="ie=edge" />

    <link rel="stylesheet" href="style.css" />

    <title>Pig Game</title>

  </head>

  <body>

    <main>

      <section class="player player--0 player--active">

        <h2 class="name" id="name--0">Player 1</h2>

        <p class="score" id="score--0">43</p>

        <div class="current">

          <p class="current-label">Current</p>

          <p class="current-score" id="current--0">0</p>

        </div>

      </section>

      <section class="player player--1">

        <h2 class="name" id="name--1">Player 2</h2>

        <p class="score" id="score--1">24</p>

        <div class="current">

          <p class="current-label">Current</p>

          <p class="current-score" id="current--1">0</p>

        </div>

      </section>

  

      <img src="dice-5.png" alt="Playing dice" class="dice" />

      <button class="btn btn--new">🔄 New game</button>

      <button class="btn btn--roll">🎲 Roll dice</button>

      <button class="btn btn--hold">📥 Hold</button>

    </main>

    <script src="script.js"></script>

  </body>

</html>


Something is wrong. My dice didn't roll. help me find the problems


There are a few issues in the code that need to be addressed to ensure the dice rolls correctly:

1. **Incorrect `Math.random()` Usage:**
   - In the `btnRoll` event listener, the way `Math.random()` is used is incorrect. It should be `Math.random()`, not `Math).random()`.
   - The correct line should be:
     ```javascript
     const dice = Math.trunc(Math.random() * 6) + 1;
     ```

2. **Dice Image Update:**
   - After rolling the dice, you need to update the dice image based on the result.
   - You can do this by setting the `src` attribute of `diceEl` to display the correct dice image.
   - Add these lines inside your event listener function after calculating `dice`:
     ```javascript
     diceEl.src = `dice-${dice}.png`;
     diceEl.classList.remove('hidden');
     ```

3. **Ensure Dice Images Exist:**
   - Make sure that you have images named `dice-1.png`, `dice-2.png`, ..., up to `dice-6.png` in your project directory where your HTML file can access them.

Here's how your updated event listener should look:

```javascript
btnRoll.addEventListener('click', function () {
  // Generate a random dice roll
  const dice = Math.trunc(Math.random() * 6) + 1;

  // Display dice
  diceEl.src = `dice-${dice}.png`;
  diceEl.classList.remove('hidden');

  // Further logic for updating scores and switching players can be added here
});
```

With these changes, clicking the Roll Dice button should successfully roll the die and update its image based on the outcome. Ensure that all necessary resources like images are correctly set up in your project directory.


When I restart the game. It doesn't hide my dice before clicking roll dice.


To ensure that the dice is hidden when you start a new game, you should add functionality to the "New Game" button to reset the game state, including hiding the dice. You can achieve this by adding an event listener to the `btnNew` button and resetting all necessary game elements.

Here's how you can modify your code:

1. **Add an Event Listener for `btnNew`:**
   - Create a function that resets the game state, including hiding the dice.

2. **Hide Dice on New Game:**
   - Ensure that the dice image is hidden at the start of a new game by adding `diceEl.classList.add('hidden')` in your reset function.

Here's how your updated code should look:

```javascript
// Function to initialize/reset game
function initGame() {
  // Reset scores
  score0El.textContent = 0;
  score1El.textContent = 0;

  // Hide dice
  diceEl.classList.add('hidden');

  // Reset other necessary elements like current scores and active player if needed
}

// Add event listener for New Game button
btnNew.addEventListener('click', initGame);

// Initialize game on page load
initGame();

// Rolling dice functionality
btnRoll.addEventListener('click', function () {
  // Generate a random dice roll
  const dice = Math.trunc(Math.random() * 6) + 1;

  // Display dice
  diceEl.src = `dice-${dice}.png`;
  diceEl.classList.remove('hidden');

  // Further logic for updating scores and switching players can be added here
});
```

With these changes, clicking the "New Game" button will hide the dice and reset any other necessary elements for starting a new game. Make sure to call `initGame()` initially when your script runs to set up the starting conditions.