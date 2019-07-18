package words

import (
	helpers "github.com/googlon/helpers"
)

/* Prepositions has:
	exactly 6 letters
   	end in a foo letter
   	do not contain the letter u.
*/

//WordIsPreposition - verifies if  a word is a preposition
func WordIsPreposition(word string) (isPreposition bool) {
	numberOfLetters := 6
	letterNotToContain := "u"
	if helpers.ContainsLetter(letterNotToContain, word) {
		return false
	}
	if !helpers.EndsWithFooLetter(word) {
		return false
	}
	if !helpers.HasOnlyXLetters(numberOfLetters, word) {
		return false
	}
	return true
}
