package main

import "strings"

var VOWELS = map[uint8]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

var NUMBERS_TO_SIMILAR_LETTERS = map[int32]string {
	'0': "o",
	'1': "i",
	'3': "e",
	'5': "s",
	'6': "b",
}

type Node struct {
	lettersToNode map[int32]*Node
	endsWord bool
}

func replaceAt(str* string, newChar string, i int) {
	*str = (*str)[:i] + newChar + (*str)[(i + 1):]
}

func replaceNumbersWithSimilarLetters(str *string) {
	for i, c := range *str {
		mappedLetter, containsNumber := NUMBERS_TO_SIMILAR_LETTERS[c]
		if containsNumber {
			replaceAt(str, mappedLetter, i)
		}
	}
}

func deleteRepeatedLetters(str *string, successiveVowelsLimit int, combineRepeatedLetters bool) int {
	successiveVowelCount := 0
	maxSuccessiveVowels := 0
	for i := 0; i < len(*str); i++ {
		c := (*str)[i]
		_, isVowel := VOWELS[c]
		if isVowel {
			successiveVowelCount++
			if successiveVowelCount > successiveVowelsLimit {
				replaceAt(str, "", i)
				i--
			} else {
				if combineRepeatedLetters && i > 0 && c == (*str)[i - 1] {
					replaceAt(str, "", i)
					i--
				}
				if successiveVowelCount > maxSuccessiveVowels {
					maxSuccessiveVowels = successiveVowelCount
				}
				successiveVowelCount = 0
			}
		}
	}
	return maxSuccessiveVowels
}

func getDetectionFormat(str string, successiveVowelLimit int, combineRepeatedLetters bool) (string, int) {
	detectionStr := strings.ToLower(str)
	replaceNumbersWithSimilarLetters(&detectionStr)
	maxSuccessiveVowels := deleteRepeatedLetters(&detectionStr, successiveVowelLimit, combineRepeatedLetters)
	return detectionStr, maxSuccessiveVowels
}

func makeWordTree(words []string) *Node {
	root := Node{endsWord: false}
	for _, word := range words {
		current := &root
		for i, c := range word {
			endOfWord := len(word) - 1 == i
			next, nodeExists := current.lettersToNode[c]
			if !nodeExists {
				next = &Node{endsWord: endOfWord}
				current.lettersToNode[c] = next
			} else if endOfWord {
				next.endsWord = true
			}
			current = next
		}
	}
	return &root
}

func containsAnyOfWords(str string, wordTree *Node) bool {
	for i, _ := range str {
		current := wordTree
		for j := i; j < len(str); j++ {
			if current.endsWord {
				return true
			}
			c := str[j]
			next, nodeExists := current.lettersToNode[int32(c)]
			if nodeExists {
				current = next
			} else {
				break
			}
		}
	}
	return false
}

func containsBadWords(str string, badWords []string) bool {
	var formattedBadWords []string
	for _, badWord := range badWords {
		formattedBadWord, _ := getDetectionFormat(badWord, 5, true)
		formattedBadWords = append(formattedBadWords, formattedBadWord)
	}
	wordTree := makeWordTree(formattedBadWords)
	return containsAnyOfWords(str, wordTree)
}