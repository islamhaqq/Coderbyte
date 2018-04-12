function LongestWord(sen) {
  // Sanitize string of punctuation and other non-alphabetic characters.
  const sanitizedSentence = sen.trim().replace(/[^\w\s]+/g, '');

  // Determine the longest word.
  const words = sanitizedSentence.split(' ');
  return words.reduce((longestWord, word) => word.length > longestWord.length ? word : longestWord, '');
}

module.exports = LongestWord;
