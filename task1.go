package test_task

import (
	"log"
	"regexp"
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
