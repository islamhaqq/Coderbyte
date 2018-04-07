package main

import "testing"

func TestLongestWord(t *testing.T) {
  governatorQuote1 := LongestWord("I'll be back!")
  expectQuote1 := "back"

  if governatorQuote1 != expectQuote1 {
    t.Errorf("LongestWord was incorrect, got: %s, want: %s", governatorQuote1, expectQuote1)
  }
}
