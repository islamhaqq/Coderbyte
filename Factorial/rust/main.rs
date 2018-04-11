fn main() {
    println!("{}", factorial(5));
}

// Gets the factorial of a number using the !n = n * !(n-1) algorithm.
fn factorial(num: u32) -> u32 {
    return if num > 1  { num * factorial(num - 1) } else {1};
}
