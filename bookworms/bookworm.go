package main

import (
	"encoding/json"
	"os"
)

// A bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms, and their beloved books, found therein.
// func functionName(param TypeParam, ...) TypeReturn
func loadBookworms(filePath string) ([]Bookworm, error) {
	// explicit declaraction of variables would be
	// var variableName TypeVariable = value	// type is defined
	// var variableName = value					// type is inferred
	// 				or
	// var (
	// 		variableName1 TypeVariable			// nil value
	// 		variableName2 = value				// type is inferred
	// 		variableName2 TypeVariable = value	// type and value declared
	// )
	// NOTE: type or value (or both) need to be specified
	// inferred declaraction of variables
	// variableName := value					// type is inferred
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the variable bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil

}
