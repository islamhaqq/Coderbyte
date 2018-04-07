package main
import "fmt"
import "strings"
import "regexp"

// Sanitize the string of punctuation.
func sanitizeString(theString string) string {
  alphaRegex, _ := regexp.Compile(`[^\w\s]+`)
  return alphaRegex.ReplaceAllString(theString, "")
}

// Get the longest word in a sentence.
func LongestWord(sen string) string{
  // Get the words of the sentence.
  words := strings.Split(sanitizeString(sen), " ")

  // Determine the longest word amongst those words.
  longestWord := words[0]
  for _, word :=range words {
    if len(word) > len(longestWord) {
      longestWord = word
    }
  }

  return longestWord
}

func main() {
  fmt.Println(LongestWord("fun&!! time"))
}
