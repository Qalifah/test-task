package test_task

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// generate takes a boolean flag and generates random correct strings if the parameter is true and random incorrect strings if the flag is false.
// function is fairly hard to implement, should take me at most 20 minutes to implement.
// took me around 20 minutes to implement.
func generate(correct bool) string {
	if correct {
		rand.Seed(time.Now().Unix())
		strconv.Itoa(rand.Int())
		return strings.Join([]string{strconv.Itoa(rand.Int()), randSeq(3), strconv.Itoa(rand.Int()), randSeq(3)}, "-")
	}
	return randSeq(10)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSeq is a utility function that return a random string
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
