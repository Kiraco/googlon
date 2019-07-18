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
words
also represent numbers given in base 20, where each letter is a digit. The digits are ordered from the least
significant to the most significant, which is the opposite of our system. That is, the leftmost digit is the unit, the
second digit is worth 20, the third one is worth 400, and so on and so forth.
As an example, the Googlon word gxjrc represents the number 605637.

A number is pretty numbers if it satisfies all of the following properties:
	it is greater than or equal to 81827
	it is divisible by 3
*/
func TestIsPrettyNumberTrue(t *testing.T) {
	assert.Equal(t, true, IsPrettyNumber("gxjrc"))
}

func TestIsPrettyNumberFalse(t *testing.T) {
	assert.Equal(t, false, IsPrettyNumber("gxjrcx"))
}
