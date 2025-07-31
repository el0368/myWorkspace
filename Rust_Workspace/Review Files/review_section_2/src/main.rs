fn main() {
    let first_name = {
        let action_hero = "Arnold Schwarzenegger";
        // Return a slice of the first 6 bytes.
        &action_hero[0..6]
    }; // `action_hero` goes out of scope here.

    println!("The first name is still valid: {}", first_name); // Prints "Arnold"
}