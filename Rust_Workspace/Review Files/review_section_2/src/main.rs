

fn main() {
    let s1 = String::from("Hello");
    let len = calculate_length(&s1);
    println!("{}, {}", s1, len);

    
    let s2 = String::from("readers");
    let r1 = &s2;
    let r2 = &s2;
    println!("Multiple immutable references are fine: {} and {}", r1, r2);
    println!("---");


    let mut s3 = String::from("writer");
    let mr1 = &mut s3;
    // let mr2 = &mut s3; // ERROR: Cannot create a second mutable borrow.
    println!("One mutable reference is fine: {}", mr1);
    println!("---");


    let mr2 = mr1;
    println!("After move, only new mutable ref is valid: {}", mr2);
    println!("---");

    let mut s4 = String::from("LifeTime");
    let mut_ref1 = &mut s4;
    mut_ref1.push_str(" check");

    let immut_ref1 = &s4;
    println!("s4 is now: '{}'", immut_ref1);
    println!("---");

    println!("Rust prevents dangling references.");
    println!("---");

    

}
fn calculate_length(s: &String) -> usize {
    s.len()
}