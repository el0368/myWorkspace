Vue.js primarily utilizes modern **ECMAScript 6 (ES6) and newer JavaScript syntax**, but the way this syntax is applied depends heavily on which of Vue's two main architectural styles you use: the **Options API** or the **Composition API**.

The modern and recommended approach for new projects is the Composition API, especially with the `<script setup>` syntax, which feels more like writing standard JavaScript.

---

### ## 1. The Options API: An Object-Based Approach 🗂️

The Options API, which was the original and only style in Vue 2, organizes component logic using a single large **JavaScript object**. Each property in this object, known as an "option," serves a specific purpose.

The core JavaScript syntax here is the **object literal (`{...}`)** and its properties, where the values are often functions or other objects.

#### Key JavaScript Syntaxes in the Options API:

- data() function:
    
    The data option must be a function that returns an object. This is a crucial JavaScript pattern to ensure each component instance gets its own unique, isolated state.
    
    JavaScript
    
    ```JavaScript
    export default {
      data() {
        // A function that returns a state object
        return {
          message: 'Hello, Vue!',
          counter: 0
        };
      }
    }
    ```
    
- methods object:
    
    This is an object where each property is a function. These functions automatically have their this context bound to the component instance, allowing you to access data and other methods with this.counter.
    
    JavaScript
    
    ```JavaScript
    export default {
      //...data
      methods: {
        // An object containing functions (methods)
        increment() {
          this.counter++;
        },
        greet() {
          alert(this.message);
        }
      }
    }
    ```
    
- computed and watch objects:
    
    Similar to methods, these are objects containing functions that handle reactive computations and side effects.
    
    JavaScript
    
    ```JavaScript
    export default {
      //...data, methods
      computed: {
        doubledCounter() {
          return this.counter * 2;
        }
      }
    }
    ```
    

---

### ## 2. The Composition API: A Function-Based Approach 🧩

Introduced in Vue 3, the Composition API allows you to organize logic by feature rather than by option type. This approach relies heavily on importing functions directly from Vue and using them to define reactive state and logic.

The core JavaScript syntax shifts to **ES6 Modules (`import`/`export`)** and a more **functional style**.

#### Key JavaScript Syntaxes in the Composition API (with `<script setup>`):

The `<script setup>` syntax is a compiler macro that dramatically simplifies Composition API code, making it the preferred modern style.

- ES6 Modules (import):
    
    You import only the functions you need from Vue. This leverages JavaScript's native module system.
    
    JavaScript
    
    ```JavaScript
    import { ref, computed, onMounted } from 'vue';
    ```
    
- Variable Declarations (const, let):
    
    You declare reactive state directly as variables using functions like ref and reactive. This feels much more like standard JavaScript.
    
    JavaScript
    
    ```JavaScript
    // 'ref' creates a reactive reference. Its value is accessed with .value
    const counter = ref(0);
    const message = ref('Hello, Composition API');
    ```
    
- Functions (function or Arrow Functions):
    
    Methods are just regular JavaScript functions declared in the same scope. There's no special methods object.
    
    JavaScript
    
    ```JavaScript
    function increment() {
      // .value is needed to access/mutate the value inside a ref
      counter.value++;
    }
    
    const greet = () => {
      alert(message.value);
    };
    ```
    
- computed with Arrow Functions:
    
    Computed properties are created by passing a getter function (often an arrow function for brevity) to the computed import.
    
    JavaScript
    
    ```JavaScript
    const doubledCounter = computed(() => counter.value * 2);
    ```
    

---

### ## 3. General ES6+ Syntax Used Across Both APIs

Regardless of the API style you choose, a modern Vue developer will constantly use these JavaScript features:

- **Arrow Functions (`=>`):** For concise inline functions, especially in array methods (`.map()`, `.filter()`) and as callbacks.
    
- **Destructuring (`{ ... }`, `[ ... ]`):** Heavily used for importing from Vue and for extracting properties from objects.
    
- **Async/Await:** The standard way to handle asynchronous operations like fetching data from an API, making the code clean and readable.
    
- **Template Literals (`` `...` ``):** Useful for constructing complex strings, though Vue's template syntax often handles this more elegantly.
    
- **Spread Syntax (`...`):** For merging objects or arrays, commonly used when managing state or passing props.
    

In summary, while Vue is built on JavaScript objects and functions, the **Composition API with `<script setup>`** most closely resembles writing modern, idiomatic JavaScript, making it the current standard and the syntax Vue.js uses most heavily in new development.