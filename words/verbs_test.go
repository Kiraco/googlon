package words

import (
	"testing"

	helpers "github.com/googlon/helpers"

	"github.com/stretchr/testify/assert"
)

func init() {
	helpers.InitMaps()
}

/*
Verbs are words of:
  6 letters or more
  ends in a bar letter.
	* If a verb starts in a bar letter, the verb is inflected in its subjunctive form

	fooLetters := {"u", "d", "x", "s", "m", "p", "f"}
	barLetters := {"o", "c", "q", "n", "w", "y", "h", "e", "l", "j", "r", "g", "i"}

*/
func TestIsVerbAndIsSubjunctive(t *testing.T) {
	isVerb, isSubjunctive := WordIsVerbAndIsSubjunctive("oocdxo")
	assert.Equal(t, true, isVerb)
	assert.Equal(t, true, isSubjunctive)
}

func TestIsVerbAndIsNotSubjunctive(t *testing.T) {
	isVerb, isSubjunctive := WordIsVerbAndIsSubjunctive("uocdxo")
	assert.Equal(t, true, isVerb)
	assert.Equal(t, false, isSubjunctive)
}

func TestIsNotVerbAndIsNotSubjunctive(t *testing.T) {
	isVerb, isSubjunctive := WordIsVerbAndIsSubjunctive("uocdxu")
	assert.Equal(t, false, isVerb)
	assert.Equal(t, false, isSubjunctive)
}

func TestIsNotVerbAndIsNotSubjunctiveLessLetters(t *testing.T) {
	isVerb, isSubjunctive := WordIsVerbAndIsSubjunctive("uocdx")
	assert.Equal(t, false, isVerb)
	assert.Equal(t, false, isSubjunctive)
}
