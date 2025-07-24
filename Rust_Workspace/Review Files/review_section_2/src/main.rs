
fn fun1(num: i64) {
    if num == 0 {
        return
    }
    
    println!("{}", num);
    fun1(num - 1);
}

fn main() {

    let num: i64 = 3;
    fun1(num);


}


