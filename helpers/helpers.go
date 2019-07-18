package helpers

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// FooLetters - Map with index for Foo
var FooLetters = make(map[string]int)

// BarLetters - Map with index for Foo
var BarLetters = make(map[string]int)

// LexOrder - Map with index of the alphabet for numerical position
var LexOrder = make(map[string]int)

var alphabetOrder = "sxocqnmwpfyheljrdgui"
var fooLetters = "udxsmpf"
var barLetters = "ocqnwyheljrgi"
var weights = map[rune]int{}

// InitMaps - Initialize the maps with the correct index for search and sorting
func InitMaps() {

	for i, s := range fooLetters {
		FooLetters[string(s)] = i
	}
	for i, s := range barLetters {
		BarLetters[string(s)] = i
	}
	for i, s := range alphabetOrder {
		weights[s] = i
		LexOrder[string(s)] = i
	}
}

func ContainsLetter(letter string, word string) bool {
	return strings.Contains(word, letter)
}

func EndsWithFooLetter(word string) bool {
	_, exists := FooLetters[word[len(word)-1:]]
	return exists
}

func EndsWithBarLetter(word string) bool {
	_, exists := BarLetters[word[len(word)-1:]]
	return exists
}

func HasOnlyXLetters(numberOfLetters int, word string) bool {
	return len(word) == numberOfLetters
}

func HasXOrMoreLetters(numberOfLetters int, word string) bool {
	return len(word) >= numberOfLetters
}

func StartsWithBarLetter(word string) bool {
	_, exists := BarLetters[word[:1]]
	return exists
}

// CustomSort - custom sort algorithm for alphabet given
func CustomSort(words []string) []string {
	sort.Slice(words, func(i int, j int) bool {
		return less(words[i], words[j])
	})
	return words
}

func less(s1, s2 string) bool {
	for {
		switch e1, e2 := len(s1) == 0, len(s2) == 0; {
		case e1 && e2:
			return false // Both empty, they are equal (not less)
		case !e1 && e2:
			return false // s1 not empty but s2 is: s1 is greater (not less)
		case e1 && !e2:
			return true // s1 empty but s2 is not: s1 is less
		}

		r1, size1 := utf8.DecodeRuneInString(s1)
		r2, size2 := utf8.DecodeRuneInString(s2)

		// Check if both are custom, in which case we use custom order:
		custom := false
		if w1, ok1 := weights[r1]; ok1 {
			if w2, ok2 := weights[r2]; ok2 {
				custom = true
				if w1 != w2 {
					return w1 < w2
				}
			}
		}
		if !custom {
			// Fallback to numeric rune comparison:
			if r1 != r2 {
				return r1 < r2
			}
		}
		s1, s2 = s1[size1:], s2[size2:]
	}
}

func RemoveDuplicatesFromSlice(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; ok {
			// duplicate item
		} else {
			m[item] = true
		}
	}
	var result []string
	for item := range m {
		result = append(result, item)
	}
	return result
}
