#[derive(Debug)]
struct Coffee {
    name: String,
    price: f64,
    is_hot: bool,
}

fn main() {
    let name = String::from("Latte");
    println!("1. Created String in main: {}", name);

    let coffee = make_coffee(name, 4.99, true);

    println!("5. Final result in main: {:?}", coffee);


    let espresso = make_coffee_shorthand(
        String::from("Expresso"),
        2.5,
        true,
    );

    println!("Created with shorthand: {:?}", espresso);

    let coffees = vec![
        make_coffee_shorthand(String::from("Cappucino"), 3.75, true),
        make_coffee_shorthand(String::from("Iced Coffee"), 3.25, true),
        make_coffee_shorthand(String::from("Mocha"), 4.5, true),
    ];

    for coffee in &coffees {
        println!("Coffee: {:?}", coffee);
    }

}

fn make_coffee(name: String, price: f64, is_hot: bool) -> Coffee {

    let coffee_instance = Coffee {
        name: name,
        price: price,
        is_hot: is_hot,
    };

    coffee_instance
}

fn make_coffee_shorthand(name: String, price: f64, is_hot: bool) -> Coffee {

    Coffee {
        name,
        price,
        is_hot,
    }
}


