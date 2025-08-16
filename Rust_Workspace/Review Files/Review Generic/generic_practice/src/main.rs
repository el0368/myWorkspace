#[derive(Debug)]
struct DaliSanwich {
    ham: String
}

fn main() {
    let result_any = identify_any::<f32>(8.36);
    println!("{:?}", &result_any);

    let my_sandwich = DaliSanwich {
        ham : String::from("Serrano")
    };

    let result_struct = identify_bool(my_sandwich.ham);
    println!("{:?}", result_struct);

}


fn identify_any<T>(value: T) -> T {
    value
}

fn identify_bool<T>(value: T) -> T {
    value
}

