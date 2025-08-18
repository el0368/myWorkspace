fn play(instrument_option: Option<&String>) {
    match instrument_option {
        Option::Some(instrument) => {
            println!("{}", instrument);
        },
        Option::None => {
            println!("Singing with my voice");
        }
        
    }
}


fn main() {

    let musical_instruments = vec![
        String::from("Guita"),
        String::from("Drums"),
        String::from("Bass"),
    ];

    let bass = musical_instruments.get(0);
    let invalid_instrument = musical_instruments.get(1);

    play(bass);
    play(invalid_instrument);




}