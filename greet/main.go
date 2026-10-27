package main

import (
	"fmt"
)

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
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
