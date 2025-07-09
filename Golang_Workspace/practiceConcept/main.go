package main

import "fmt"

type product struct {
	id       int
	price    float64
	name     string
	category string
}

// bookList prints the details of a product.
func (p product) bookList() {
	fmt.Println("\n--- Book Details ---")
	fmt.Printf("ID: %d\n", p.id)
	fmt.Printf("Price: $%.2f\n", p.price)
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Category: %s\n", p.category)
}

func main() {
	// Declare variables to store user input.
	var inputID int
	var inputPrice float64
	var inputName string
	var inputCategory string

	// Get book details from the user.
	fmt.Println("Enter Book's ID:")
	fmt.Scanln(&inputID)

	fmt.Println("Enter Book's price:")
	fmt.Scanln(&inputPrice)

	fmt.Println("Enter Book's name:")
	fmt.Scanln(&inputName)

	fmt.Println("Enter Book's category:")
	fmt.Scanln(&inputCategory)

	// Create a product instance with the user's input.
	book := product{
		id:       inputID,
		price:    inputPrice,
		name:     inputName,
		category: inputCategory,
	}

	// Display the details of the book that was just created.
	book.bookList()
}