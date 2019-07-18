package words

import (
	helpers "github.com/googlon/helpers"
)

/*
Verbs are words of:
  6 letters or more
  ends in a bar letter.
	* If a verb starts in a bar letter, the verb is inflected in its subjunctive form
*/

// WordIsVerbAndIsSubjunctive - returns if a word is a verb and if its also in subjunctive form
func WordIsVerbAndIsSubjunctive(word string) (isVerb bool, isSubjunctiveForm bool) {
	numberOfLetters := 6
	if !helpers.HasXOrMoreLetters(numberOfLetters, word) {
		return false, false
	}
	if !helpers.EndsWithBarLetter(word) {
		return false, false
	}
	// Here we already have a verb, we will check if its in subjunctive form or not
	if !helpers.StartsWithBarLetter(word) {
		return true, false
	}
	return true, true

}
