let color = "rgb(132, 173, 254)";

// 1. get element by ID
let header = document.getElementById("header");
header.style.backgroundColor = color;

// 2.get element by class name
let classElements = document.getElementsByClassName("header");
for (let i = 0; i < classElements.length; i++) {
  classElements[i].style.backgroundColor = "rgb(132, 173, 254)";
}

// 2. get elements by tag name
let tagElements = document.getElementsByTagName("p");
for (let i = 0; i < tagElements.length; i++) {
  tagElements[i].style.backgroundColor = "rgb(132, 173, 254)";
}

let queryElement = document.querySelector("#query");
queryElement.style.backgroundColor = "pink";

let queryAllElement = document.querySelectorAll(".query-all");
queryAllElement.forEach((element) => {
  element.style.backgroundColor = "pink";
});
