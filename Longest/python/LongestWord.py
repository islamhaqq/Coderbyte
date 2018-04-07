import re

def LongestWord(sen):
    alphaRegex = re.compile('[^\w\s]+\g')
    sanitizedSentence = alphaRegex.sub('', sen)
    words = sanitizedSentence.split(' ')
    return words

print(LongestWord("I'll be back!"))
