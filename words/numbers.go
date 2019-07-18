package words

import (
	"math"

	helpers "github.com/googlon/helpers"
)

// IsPrettyNumber - verify if a word is a number and if its a pretty number
func IsPrettyNumber(word string) bool {
	minValueForPretty := 81827
	wordValue := 0
	for index, char := range word {
		value := helpers.LexOrder[string(char)]
		wordValue = wordValue + (int(math.Pow(20, float64(index))) * value)
	}
	if wordValue < minValueForPretty {
		return false
	}
	mod := wordValue % 3
	return mod == 0
}
