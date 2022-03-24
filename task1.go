package test_task

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// testValidity checks if a string is a sequence of numbers followed by dash followed by text
// function is easy to implement, should take me at most three minutes to implement.
// took me around 2 minutes to implement.
func testValidity(input string) bool {
	pattern := `([0-9]+-[a-zA-Z]+-?)+`
	matched, err := regexp.MatchString(pattern, input)
	if err != nil {
		log.Fatal(err)
	}
	return matched
}

// averageNumber takes a string and returns the average number from all the numbers
// function is easy to implement, should take me at most five minutes to implement.
// took me around 5 minutes to implement
func averageNumber(input string) int {
	r, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	matchedArray := r.FindAllString(input, -1)
	var intVal int
	for _, i := range matchedArray {
		val, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		intVal += val
	}
	return intVal / len(matchedArray)
}

// wholeStory takes the string and returns a text that is composed of all the text words separated by spaces
// function is easy to implement, should take me at most five minutes to implement.
// took me around 4 minutes to implement
func wholeStory(input string) string {
	r, err := regexp.Compile(`[a-zA-Z]+`)
	if err != nil {
		log.Fatal(err)
	}
	matchedArray := r.FindAllString(input, -1)
	return strings.Join(matchedArray, " ")
}

// storyStats returns the shortest word, the longest word, the average word length and the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down
// function is fairly hard to implement, should take me at most 12 minutes to implement.
func storyStats(input string) (string, string, int, []string) {
	return "", "", 0, nil
}
