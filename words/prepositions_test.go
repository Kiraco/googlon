package words

import (
	"testing"

	helpers "github.com/googlon/helpers"

	"github.com/stretchr/testify/assert"
)

func init() {
	helpers.InitMaps()
}

/* Prepositions has:
	exactly 6 letters
   	end in a foo letter
	do not contain the letter u.

	fooLetters := {"u", "d", "x", "s", "m", "p", "f"}
	barLetters := {"o", "c", "q", "n", "w", "y", "h", "e", "l", "j", "r", "g", "i"}

*/
func TestIsPrepositionTrue(t *testing.T) {
	assert.Equal(t, true, WordIsPreposition("pocdxs"))
}

func TestIsPrepositionFalseNot6Letters(t *testing.T) {
	assert.Equal(t, false, WordIsPreposition("pocdxss"))
}

func TestIsPrepositionFalseNotEndsInFooLetter(t *testing.T) {
	assert.Equal(t, false, WordIsPreposition("pocdxo"))
}

func TestIsPrepositionFalseContainsLetterU(t *testing.T) {
	assert.Equal(t, false, WordIsPreposition("pucdxs"))
}
