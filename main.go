package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	helpers "github.com/googlon/helpers"
	wordsProcess "github.com/googlon/words"
)

func main() {
	helpers.InitMaps()

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/process", ProcessText)
	http.ListenAndServe(":8080", nil)
}

// Ping - Verify if app is running
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App is running... %s!", r.URL.Path[1:])
}

//ProcessText - process the test posted to it and returns the processed answer
func ProcessText(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(text))
	var prepositions []string
	var verbs []string
	var subjunctiveVerbs []string
	var prettyNumbers []string
	var orderedWords []string
	for _, word := range words {
		isPreposition := wordsProcess.WordIsPreposition(word)
		if isPreposition {
			prepositions = append(prepositions, word)
		}
		isVerb, isSubjunctive := wordsProcess.WordIsVerbAndIsSubjunctive(word)
		if isVerb {
			verbs = append(verbs, word)
			if isSubjunctive {
				subjunctiveVerbs = append(subjunctiveVerbs, word)
			}
		}
	}
	distinctWords := helpers.RemoveDuplicatesFromSlice(words)
	for _, word := range distinctWords {
		isPrettyNumber := wordsProcess.IsPrettyNumber(word)
		if isPrettyNumber {
			prettyNumbers = append(prettyNumbers, word)
		}
	}
	orderedWords = helpers.CustomSort(distinctWords)

	response := fmt.Sprintf("1) There are %d prepositions in the text\n", len(prepositions))
	response = response + fmt.Sprintf("2) There are %d verbs in the text\n", len(verbs))
	response = response + fmt.Sprintf("3) There are %d subjunctive verbs in the text\n", len(subjunctiveVerbs))
	response = response + fmt.Sprintf("4) Vocabulary list: %s\n", orderedWords)
	response = response + fmt.Sprintf("5) There are %d distinct pretty numbers in the text", len(prettyNumbers))

	w.Write([]byte(response))
}
