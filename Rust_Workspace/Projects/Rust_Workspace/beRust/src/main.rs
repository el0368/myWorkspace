
#[derive(Debug)]

struct Deck {
    cards: Vec<String>
}

fn main() {
    let suits = vec!["Hears", "Spades", "Diamonds"];
    let values = vec!["Ace", "Two", "Three"];

let mut cards: Vec<String> = vec![];

for suit in suits.iter() {
    for value in values.iter() {
        let card = format!("{} of {}", value, suit);
        cards.push(card);
    }
}

    let deck = Deck { cards };

    println!("Heres your deck: {:?}", deck);
}
 