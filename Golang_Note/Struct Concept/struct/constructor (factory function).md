You build a separate constructor (factory function) for each outer struct (`User` and `Company`). The constructor for the outer struct is responsible for initializing all of its fields, including the embedded `ContactInfo` struct.

---

### The Code with Constructors

Here is your example, refactored to use factory functions.

Go

```Go
package main

import "fmt"

// This struct contains common data fields.
type ContactInfo struct {
	Email string
	Phone string
}

type User struct {
	ContactInfo // Embeds the common data
	Username    string
}

type Company struct {
	ContactInfo // Embeds the same common data
	Name        string
}

// Constructor for the User struct
func NewUser(username, email, phone string) *User {
	return &User{
		Username: username,
		ContactInfo: ContactInfo{
			Email: email,
			Phone: phone,
		},
	}
}

// Constructor for the Company struct
func NewCompany(name, email, phone string) *Company {
	return &Company{
		Name: name,
		ContactInfo: ContactInfo{
			Email: email,
			Phone: phone,
		},
	}
}

func main() {
	// Use the constructors to create the instances.
	user := NewUser("alex", "alex@example.com", "111-222-3333")
	company := NewCompany("Big Corp", "contact@bigcorp.com", "444-555-6666")

	// The fields from ContactInfo are "promoted" and can be accessed directly.
	fmt.Printf("User Email: %s\n", user.Email)
	fmt.Printf("Company Phone: %s\n", company.Phone)
}
```

---

### How It Works

- The `NewUser` function takes all the data needed to create a complete `User`, including the fields required for the embedded `ContactInfo`.
    
- Similarly, the `NewCompany` function gathers all the necessary data for a `Company`.
    
- This approach encapsulates the creation logic for each type, making your code cleaner and ensuring that `User` and `Company` instances are always created in a consistent way.