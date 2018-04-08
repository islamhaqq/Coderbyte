package main

import "testing"

// Unit test Factorial().
func TestFactorial(t *testing.T) {
  if (Factorial(5) != 120) {
    t.Errorf("Expected Factorial(%d) to be %d, but got %d", 5, 120, Factorial(5))
  }

  if (Factorial(1) != 1) {
    t.Errorf("Expected Factorial(%d) to be %d, but got %d", 1, 1, Factorial(1))
  }

  if (Factorial(15) != 1307674368000) {
    t.Errorf("Expected Factorial(%d) to be %d, but got %d", 15, 1307674368000, Factorial(15))
  }
}
