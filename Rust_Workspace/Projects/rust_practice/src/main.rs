use std::collections::HashMap;

fn main() {
    // Create an empty HashMap using new()
    let mut map: HashMap<String, i32> = HashMap::new();
    
    // Add some key-value pairs
    map.insert("apple".to_string(), 5);
    map.insert("banana".to_string(), 3);
    map.insert("orange".to_string(), 8);
    
    // Print the HashMap
    println!("HashMap: {:?}", map);
    
    // Access values
    if let Some(value) = map.get("apple") {
        println!("Apple count: {}", value);
    }
    
    // Another example with different types
    let mut scores: HashMap<&str, u32> = HashMap::new();
    scores.insert("Alice", 100);
    scores.insert("Bob", 85);
    scores.insert("Charlie", 92);
    
    println!("Scores: {:?}", scores);
    
    // Example with integer keys
    let mut numbers: HashMap<i32, String> = HashMap::new();
    numbers.insert(1, "One".to_string());
    numbers.insert(2, "Two".to_string());
    numbers.insert(3, "Three".to_string());
    
    println!("Numbers: {:?}", numbers);
}