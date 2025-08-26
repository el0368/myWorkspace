
fn main() {

    let pizza_diameters = vec![8, 10, 12, 14];
    
    let pepperoni = String::from("Pepperoni");
    let mushroom = String::from("Mushroom");
    let sausage = String::from("Sausage");

    let pizza_topping = vec![pepperoni, mushroom, sausage];

    let topping_reference = &pizza_topping[1];
    println!("{:?}", topping_reference);

    let option = pizza_topping.get(0);
    match option {
        Some(topping) => {
            println!("{:?}", topping);
        },
         None => {
            println!("No Value");
        }



}
}

