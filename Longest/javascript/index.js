function LongestWord(sen) {
  const sanitizedSentence = sen.replace(/[^\w\s]+/g, '');
  const words = sanitizedSentence.split(' ');
  return words.reduce((longestWord, word) => word.length > longestWord.length ? word : longestWord, '');
}

module.exports = LongestWord;
