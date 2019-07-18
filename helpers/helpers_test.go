package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wordStartsBar = "oodciu"
var wordEndsFoo = "podciu"
var wordEndsBar = "podcio"

func init() {
	InitMaps()
}

func TestContainsLetterTrue(t *testing.T) {
	assert.Equal(t, true, ContainsLetter("p", wordEndsFoo))
}

func TestContainsLetterFalse(t *testing.T) {
	assert.Equal(t, false, ContainsLetter("s", wordEndsFoo))
}

func TestEndsWithFooLetterTrue(t *testing.T) {
	assert.Equal(t, true, EndsWithFooLetter(wordEndsFoo))
}

func TestEndsWithFooLetterFalse(t *testing.T) {
	assert.Equal(t, false, EndsWithFooLetter(wordEndsBar))
}

func TestEndsWithBarLetterTrue(t *testing.T) {
	assert.Equal(t, true, EndsWithBarLetter(wordEndsBar))
}

func TestEndsWithBarLetterFalse(t *testing.T) {
	assert.Equal(t, false, EndsWithBarLetter(wordEndsFoo))
}

func TestHasOnlyXLettersTrue(t *testing.T) {
	assert.Equal(t, true, HasOnlyXLetters(6, wordEndsFoo))
}

func TestHasOnlyXLettersFalse(t *testing.T) {
	assert.Equal(t, false, HasOnlyXLetters(5, wordEndsFoo))
}

func TestHasXOrMoreLettersMoreLettersTrue(t *testing.T) {
	assert.Equal(t, true, HasXOrMoreLetters(5, wordEndsFoo))
}
func TestHasXOrMoreLettersEqualLettersTrue(t *testing.T) {
	assert.Equal(t, true, HasXOrMoreLetters(6, wordEndsFoo))
}
func TestHasXOrMoreLettersLessLettersFalse(t *testing.T) {
	assert.Equal(t, false, HasXOrMoreLetters(7, wordEndsFoo))
}

func TestStartsWithBarLetterTrue(t *testing.T) {
	assert.Equal(t, true, StartsWithBarLetter(wordStartsBar))
}

func TestStartsWithBarLetterFalse(t *testing.T) {
	assert.Equal(t, false, StartsWithBarLetter(wordEndsFoo))
}
