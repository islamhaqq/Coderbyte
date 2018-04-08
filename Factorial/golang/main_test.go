package main

import "testing"

// Unit test Factorial().
func TestFactorial(t *testing.T) {
  if (Factorial(5) != 120) {
    t.Errorf("Expected Factorial(%d) to be %d, but got %d", 5, 120, Factorial(5))
  }
}
