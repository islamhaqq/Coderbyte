fn main() {
    println!("{}", factorial(5));
}

// Gets the factorial of a number using the !n = n * !(n-1) algorithm.
fn factorial(num: u32) -> u32 {
    return if num > 1  { num * factorial(num - 1) } else {1};
}

// Define unit tests.
#[cfg(test)]
mod tests {
    use factorial;

    #[test]
    fn it_gets_the_factorial() {
        assert_eq!(factorial(0), 1);
        assert_eq!(factorial(1), 1);
        assert_eq!(factorial(2), 2);
        assert_eq!(factorial(3), 6);
        assert_eq!(factorial(4), 24);
        assert_eq!(factorial(5), 120);
        assert_eq!(factorial(6), 720);
    }
}
