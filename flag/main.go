package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"en",
		"The required language, e.g. en, el, ur, fr, he, vi, ")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents a language
type language string

var phrasebook = map[language]string{
	"el": "ΧΧαίρίρετε ΚόΚόσμε",  // Greek
	"en": "Hello world",         // English
	"fr": "Bonjour le monde",    // French
	"he": "םולשםולש םלועםלוע",   // Hebrew
	"ur": "ولیہولیہ ایندایند",   // Urdu
	"vi": "Xin chào Thếế Giớới", // Vietnamese
}

// greet says hello to the world in various languages
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
