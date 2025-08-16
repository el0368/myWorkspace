// A generic enum to represent a cheesesteak order.
#[derive(Debug)]
struct User {
    username: String,
    email: String,
    active: bool,
}

fn main() {
    let mut user1 = User {
        username: String::from("user123"),
        email: String::from("user@example.com"),
        active: true,
    };

    println!("{:#?}", user1.username);
    user1.active = false;
    user1.email = String::from("changed_email@example.com");

    println!("{}", user1.email);
}