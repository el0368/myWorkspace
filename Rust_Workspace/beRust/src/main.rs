
#[derive(Debug)]

struct Deck {
    cards: Vec<String>
}

fn main() {
    let suits = vec!["Hears", "Spades", "Diamonds"];
    let values = vec!["Ace"],

    let deck = Deck { cards: vec![]};

    println!("Heres your deck: {:?}", deck);
}
 