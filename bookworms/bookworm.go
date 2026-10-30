package main

import (
	"encoding/json"
	"os"
)

// A bookworm contains the list of books on a bookworm's shelf.
// following format from json
type Bookworm struct {
	// Name is capitalized so that it is exported and can be used by the json package.
	// struct tags are used to specify how the field is encoded/decoded in JSON.
	Name  string `json:"name"`
	// Go convention is to use plural for slices, so we use Books instead of Book.
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
// struct tags used within above
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms, and their beloved books, found therein.
// func functionName(param TypeParam, ...) (TypeReturn1, TypeReturn2, ...) { function content }
func loadBookworms(filePath string) ([]Bookworm, error) {
	// explicit declaration of variables would be
	// var variableName TypeVariable = value	// type is defined
	// var variableName = value					// type is inferred
	// 				or
	// var (
	// 		variableName1 TypeVariable			// nil value
	// 		variableName2 = value				// type is inferred
	// 		variableName2 TypeVariable = value	// type and value declared
	// )
	// NOTE: type or value (or both) need to be specified
	// inferred declaration of variables
	// variableName := value					// type is inferred
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the variable bookworms.
	// json.NewDecoder(f) creates a new decoder that reads from the file f.
	// Decode(&bookworms) decodes the JSON data from the file and stores it in the variable bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil

}

// takes data which is a list of bookworms and returns the books in common as a slice of Book
func findCommonBooks (bookworms []Bookworm) []Book {
	return nil
	
}
/*
Notes:
Built-in types for data collection:
- map: unordered array that contains a key-value pair.
	- slices, arrays, channels

Writing to a map example:
	myMap[k] = v

Reading from a map example: same format as reading from slices
	var mySlice []string
	...
	v := mySlice[3]
	mapped := map[int]string{}
	...
	v := mapped[3]

difference being maps also return a boolean indicating whether the key was found.
if the key is not found the value returned is zero of the valueType.
	v, ok := mapped[3]
	if !ok {
		// key not found
	}
or
	if v, ok := mapped[3]; !ok {
		// key not found
	}


	*/

