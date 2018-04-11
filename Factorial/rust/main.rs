fn main() {
    println!("{}", factorial(5));
}

// Gets the factorial of a number.
fn factorial(num: u32) -> u32 {
    return if num == 1 { 1 } else { num * factorial(num - 1) };
}
