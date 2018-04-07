import re

def LongestWord(sen):
    # Sanitize string of punctuation and other non-alphabetic characters.
    alphaRegex = re.compile('[^\w\s]+')
    sanitizedSentence = alphaRegex.sub('', sen)

    # Determine the longest word.
    words = sanitizedSentence.split(' ')
    return reduce(lambda word, longestWord: word if len(word) > len(longestWord) else longestWord, words)

print(LongestWord("I'll bee back!"))
