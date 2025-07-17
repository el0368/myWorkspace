// EXAMINE THE DOCUMENT OBJECT //

// console.dir(document);
// console.log(document.domain);
// console.log(document.URL);
// console.log(document.title);
// document.title = 123;
// console.log(document.doctype);
// console.log(document.head);
// console.log(document.body);
// console.log(document.all);
// console.log(document.all[9]);
// // document.all[9].textContent = "hello";
// console.log(document.forms);
// console.log(document.links);
// console.log(document.images);

// GET ELEMENT BY ID //
// console.log(document.getElementById("header-title"));
// var headerTitle = document.getElementById("header-title");
// var header = document.getElementById("main-header");
// console.log(headerTitle);
// headerTitle.textContent = "Hello";
// headerTitle.innerText = "Goodbye";
// console.log(headerTitle.innerText);
// headerTitle.innerHTML = "<h3>Hi</h3>";
// header.style.borderBottom = "solid 10px #000";

// GET ELEMENTS BY CLASS NAME //

// var items = document.getElementsByClassName("list-group-item");
// console.log(items);
// console.log(items[0]);
// items[1].textContent = "hello2";
// items[1].style.fontWeight = "bold";

// // give error, need to iterate
// // items[1].style.backgroundColor = "yellow";

// for (var i = 0; i < items.length; i++) {
//   items[i].style.backgroundColor = "#f4f4f4";
// }

// GET ELEMENT BY TAG NAME //

// var li = document.getElementsByTagName("li");
// console.log(li);
// console.log(li[0]);
// li[1].textContent = "hello2";
// li[1].style.fontWeight = "bold";
// li[1].style.backgroundColor = "yellow";

// // give error, need to iterate
// // items[1].style.backgroundColor = "yellow";

// for (var i = 0; i < li.length; i++) {
//   li[i].style.backgroundColor = "#f4f4f4";
// }

// QuerySelector //
// var header = document.querySelector("#main-header");
// header.style.borderBottom = "solid 4px #ccc";

// var input = document.querySelector("input");
// input.value = "hello world";

// var submit = document.querySelector('input[type="submit"]');
// submit.value = "SEND";

// var item = document.querySelector(".list-group-item");
// item.style.color = "red";

// var lastItem = document.querySelector(".list-group-item:last-child");
// lastItem.style.color = "blue";

// var secondItem = document.querySelector(".list-group-item:nth-child(2)");
// secondItem.style.color = "green";
// secondItem.style.fontWeight = "bold";

// // QuerySelectorAll //
// var titles = document.querySelectorAll(".title");

// console.log(titles);
// titles[0].textContent = "Add Items, Now!";

// var odd = document.querySelectorAll("li:nth-child(odd)");
// var even = document.querySelectorAll("li:nth-child(even)");

// for (var i = 0; i < odd.length; i++) {
//   odd[i].style.backgroundColor = "#f4f4f4";
//   even[i].style.backgroundColor = "#ccc";
// }

// Traversing the DOM

// var itemList = document.querySelector("#items");
// console.log(itemList.parentNode);
// itemList.parentNode.style.backgroundColor = "#f4f4f4";
// console.log(itemList.parentNode.parentNode);

// parentElement
// console.log(itemList.parentElement);
// itemList.parentElement.style.backgroundColor = "#f4f4f4";
// console.log(itemList.parentElement.parentElement);

// childNodes
// console.log(itemList.childNodes);

// console.log(itemList.children);
// console.log(itemList.children[1]);
// itemList.children[1].style.backgroundColor = "yellow";

// // // FirstChild
// console.log(itemList.firstChild);

// // firstElementChild
// console.log(itemList.firstElementChild);
// itemList.firstElementChild.textContent = "hello 1";

// // LastChild
// console.log(itemList.lastChild);

// // LastElementChild
// console.log(itemList.lastElementChild);
// itemList.lastElementChild.textContent = "hello 4";

// // nextSibling
// console.log(itemList.nextSibling);
// // nextElementSibling
// console.log(itemList.nextElementSibling);

// // previousSibling
// console.log(itemList.previousSibling);
// // previousElementSibling
// console.log(itemList.previousElementSibling);
// itemList.previousElementSibling.style.color = "green";

// createElement

// // create a div
// var newDiv = document.createElement("div");

// // Add Class
// newDiv.className = "helloClassName";

// // Add ID
// newDiv.id = "HelloID";

// // Add attr
// newDiv.setAttribute("title", "hello div");

// // create text node
// var newDivText = document.createTextNode("Hello World!");

// // add Text to div
// newDiv.appendChild(newDivText);

// var container = document.querySelector("header .container");
// var h1 = document.querySelector("header h1");

// console.log(newDiv);

// newDiv.style.fontSize = "20px";

// container.insertBefore(newDiv, h1);

function buttonClick() {}
