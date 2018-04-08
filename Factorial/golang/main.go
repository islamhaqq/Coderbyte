package main

import "fmt"

// Calculate the factorial of a number.
func Factorial(n int) int {
  // We know !1 is 1.
  if n == 1 {
    return 1;
  }
  // Since !5 is 5 * !4 we can say !n = n * !(n-1).
  return n * Factorial(n-1)
}

func main() {
  fmt.Println(Factorial(5));
}
