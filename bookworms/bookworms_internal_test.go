package main

import (
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Bronte", Title: "Jane Eyre"}
)

// Test*(t *testing.T) functions are the entry point for go tests. They are executed by the `go test` command.
func TestLoadBookworms_Success(t *testing.T) {
	// declare testCase structure
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}
	
	// define sub-tests for go 
	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}
	// loop through sub-tests
	// for loops examples:
	// for i := 0; i < 5; i++ {
	// for i := 0; i < arrayLength; i++ {
	// for i := firstIndex; i < limit; i++ {
	// while logic:
	// for iterator.Next() {
	// for line != lastLine {
	// for !gotResponse || response.invalid() {
	// range returns two values: the key and the value. In this case, the key is the name of the test case, and the value is the testCase struct.
	for name, tc := range tests {
		// this line is general go testing pattern
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantErr {
				if err == nil {
					t.Fatal("expected err, got nothing")
				}
				return
			}

			// we aren't expecting errors here, this should be the happy path
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			// NOTE: see comment below equalbooks()
			// if !reflect.DeepEqual(got, tc.want) {
			// 		t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			// }
			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
// t *Testing.T is passed in order to user the t.Helper() function.
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	// t.Helper() indicates to the go compiler to skip this function when reporting the line number of an error. This is useful for helper functions that are called by multiple test cases, as it allows the error to be reported at the level of the test case rather than the helper function.
	t.Helper()

	if len(bookworms) != len(target) {
		// Early exit!
		return false
	}

	for i := range bookworms {
		// Verify the name of the Bookworm.
		if bookworms[i].Name != target[i].Name {
			return false
		}
		// Verify the content of the collections of Books for each Bookworm.
		
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	// Everything is equal!
	return true
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		// Early exit!
		return false
	}
	//Verify the content of the collections of Books for each Bookworm.
	for i := range target {
		if target[i] != books[i] {
			return false
		}
	}
	// Everything is equal!
	return true
}
// NOTE: another of accomplishing above is with the standard library reflect and reflect.DeepEqual. This method is not recommended for any production code as it is less efficient, but is more concise.